package provider

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

type regionResource struct {
	client *api.RegionClient
}

func (r regionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

func (r regionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The region resource is used to manage regions for a given cloud provider. See [Activating regions](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/activating_regions/) for more details.",
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true}),
		},

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Resource ID of the region.",
				Computed:            true,
			},
			"cloud_provider": schema.StringAttribute{
				MarkdownDescription: "Cloud provider. For example, \"aws\", \"azure\" or \"bah:aws\".",
				Required:            true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "BigAnimal Project ID.",
				Required:            true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
			},
			"region_id": schema.StringAttribute{
				MarkdownDescription: "Region ID of the region. For example, \"germanywestcentral\" in the Azure cloud provider or \"eu-west-1\" in the AWS cloud provider.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Region name of the region. For example, \"Germany West Central\" or \"EU West 1\".",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Region status of the region. For example, \"ACTIVE\", \"INACTIVE\", or \"SUSPENDED\".",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(api.REGION_ACTIVE),
			},
			"continent": schema.StringAttribute{
				MarkdownDescription: "Continent that region belongs to. For example, \"Asia\", \"Australia\", or \"Europe\".",
				Computed:            true,
			},
		},
	}
}

func (r *regionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.API).RegionClient()
}

type Region struct {
	ProjectID     *string      `tfsdk:"project_id"`
	CloudProvider *string      `tfsdk:"cloud_provider"`
	RegionID      *string      `tfsdk:"region_id"`
	ID            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Continent     types.String `tfsdk:"continent"`
	Status        *string      `tfsdk:"status"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r regionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config Region
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = r.update(ctx, config, &resp.State)
	resp.Diagnostics.Append(diags...)
}

func (r *regionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Region
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.read(ctx, state, &resp.State)...)
}

func (r *regionResource) read(ctx context.Context, region Region, state *tfsdk.State) frameworkdiag.Diagnostics {
	read, err := r.client.Read(ctx, *region.ProjectID, *region.CloudProvider, *region.RegionID)
	if err != nil {
		return fromErr(err, "Error reading region %v", region.RegionID)
	}
	region.ID = types.StringValue(fmt.Sprintf("%s/%s/%s", *region.ProjectID, *region.CloudProvider, *region.RegionID))
	region.Name = types.StringValue(read.Name)
	region.Status = &read.Status
	region.Continent = types.StringValue(read.Continent)
	return state.Set(ctx, &region)
}

func (r *regionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan Region
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(r.update(ctx, plan, &resp.State)...)
}

func (r *regionResource) update(ctx context.Context, region Region, state *tfsdk.State) frameworkdiag.Diagnostics {
	if err := r.client.Update(ctx, *region.Status, *region.ProjectID, *region.CloudProvider, *region.RegionID); err != nil {
		return fromErr(err, "Error updating region %v", region.RegionID)
	}

	timeout, diagnostics := region.Timeouts.Create(ctx, 60*time.Minute)
	if diagnostics != nil {
		return diagnostics
	}

	err := retry.RetryContext(
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

	if err := r.client.Update(ctx, api.REGION_INACTIVE, *state.ProjectID, *state.CloudProvider, *state.RegionID); err != nil {
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

func (r *regionResource) retryFunc(ctx context.Context, region Region) retry.RetryFunc {
	return func() *retry.RetryError {
		curr, err := r.client.Read(ctx, *region.ProjectID, *region.CloudProvider, *region.RegionID)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("error describing instance: %s", err))
		}

		if curr.Status != *region.Status {
			return retry.RetryableError(errors.New("operation incomplete"))
		}
		return nil
	}
}

func (r regionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: project_id/cloud_provider/region_id. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cloud_provider"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("region_id"), idParts[2])...)
}

func NewRegionResource() resource.Resource {
	return &regionResource{}
}
