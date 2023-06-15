package provider

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	frameworkschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

func NewRegionResource() resource.Resource {
	return &regionResource{}
}

type regionResource struct {
	client *api.API
}

func (r *regionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = frameworkschema.Schema{
		MarkdownDescription: "The region resource is used to manage regions for a given cloud provider. See [Activating regions](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/activating_regions/) for more details.",
		Blocks: map[string]frameworkschema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true}),
		},

		Attributes: map[string]frameworkschema.Attribute{
			"cloud_provider": frameworkschema.StringAttribute{
				MarkdownDescription: "Cloud provider. For example, \"aws\" or \"azure\".",
				Required:            true,
			},
			"project_id": frameworkschema.StringAttribute{
				MarkdownDescription: "BigAnimal Project ID.",
				Required:            true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"region_id": frameworkschema.StringAttribute{
				MarkdownDescription: "Region ID of the region. For example, \"germanywestcentral\" in the Azure cloud provider or \"eu-west-1\" in the AWS cloud provider.",
				Required:            true,
			},
			"name": frameworkschema.StringAttribute{
				MarkdownDescription: "Region name of the region. For example, \"Germany West Central\" or \"EU West 1\".",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"status": frameworkschema.StringAttribute{
				MarkdownDescription: "Region status of the region. For example, \"ACTIVE\", \"INACTIVE\", or \"SUSPENDED\".",
				Optional:            true,
				Default:             DefaultString("The default of region desired status", api.REGION_ACTIVE),
			},
			"continent": frameworkschema.StringAttribute{
				MarkdownDescription: "Continent that region belongs to. For example, \"Asia\", \"Australia\", or \"Europe\".",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

type Region struct {
	ProjectID     *string `tfsdk:"project_id"`
	CloudProvider *string `tfsdk:"cloud_provider"`
	RegionID      *string `tfsdk:"region_id"`
	Name          *string `tfsdk:"name"`
	Status        *string `tfsdk:"status"`
	Continent     *string `tfsdk:"continent"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r *regionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

func (r *regionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config Region
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = r.update(ctx, config, resp.State)
	resp.Diagnostics.Append(diags...)
}

func (r *regionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Region
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.read(ctx, state, resp.State)...)
}

func (r *regionResource) read(ctx context.Context, region Region, state tfsdk.State) frameworkdiag.Diagnostics {
	read, err := r.client.RegionClient().Read(ctx, *region.ProjectID, *region.CloudProvider, *region.RegionID)
	if err != nil {
		return fromErr(err, "Error reading region %v", region.RegionID)
	}

	region.Name = &read.Name
	region.Status = &read.Status
	region.Continent = &read.Continent
	return state.Set(ctx, &region)
}

func (r *regionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan Region
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.update(ctx, plan, resp.State)...)
}

func (r *regionResource) update(ctx context.Context, region Region, state tfsdk.State) frameworkdiag.Diagnostics {
	current, err := r.client.RegionClient().Read(ctx, *region.ProjectID, *region.CloudProvider, *region.RegionID)
	if err != nil {
		return fromErr(err, "Error reading region %v", region.RegionID)
	}
	if current.Status == *region.Status { // no change, exit early
		return nil
	}

	tflog.Debug(ctx, fmt.Sprintf("updating region from %s to %s", current.Status, region.Status))

	if err := r.client.RegionClient().Update(ctx, *region.Status, *region.ProjectID, *region.CloudProvider, *region.RegionID); err != nil {
		return fromErr(err, "Error updating region %v", region.RegionID)
	}

	timeout, diagnostics := region.Timeouts.Create(ctx, 60*time.Minute)
	if diagnostics != nil {
		return diagnostics
	}

	err = retry.RetryContext(
		ctx,
		timeout-time.Minute,
		r.retryFunc(ctx, region))
	if err != nil {
		return fromErr(err, "")
	}

	return r.read(ctx, region, state)
}

func (r *regionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Region
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if *state.Status == api.REGION_INACTIVE {
		return
	}

	if err := r.client.RegionClient().Update(ctx, api.REGION_INACTIVE, *state.ProjectID, *state.CloudProvider, *state.RegionID); err != nil {
		resp.Diagnostics.Append(fromErr(err, "Error deleting region %v", state.RegionID)...)
		return
	}

	timeout, diagnostics := state.Timeouts.Create(ctx, 60*time.Minute)
	resp.Diagnostics.Append(diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(
		ctx,
		timeout-time.Minute,
		r.retryFunc(ctx, state))
	if err != nil {
		resp.Diagnostics.Append(fromErr(err, "")...)
	}
}

func (r *regionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.API)
}

func (r *regionResource) retryFunc(ctx context.Context, region Region) retry.RetryFunc {
	return func() *retry.RetryError {
		curr, err := r.client.RegionClient().Read(ctx, *region.ProjectID, *region.CloudProvider, *region.RegionID)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("error describing instance: %s", err))
		}

		if curr.Status != *region.Status {
			return retry.RetryableError(errors.New("operation incomplete"))
		}
		return nil
	}
}
