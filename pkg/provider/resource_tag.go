package provider

import (
	"context"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &tagResource{}
	_ resource.ResourceWithConfigure = &tagResource{}
)

type TagResourceModel struct {
	ID      types.String `tfsdk:"id"`
	TagName types.String `tfsdk:"tag_name"`
	Color   types.String `tfsdk:"color"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

type tagResource struct {
	client *api.TagClient
}

func (tr *tagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	tr.client = req.ProviderData.(*api.API).TagClient()
}

func (tr *tagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

func (tf *tagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tags will enable users to categorize and organize resources across types and improve the efficiency of resource retrieval",
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true}),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tag_name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"color": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (tr *tagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config TagResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tagId, err := tr.client.Create(ctx, commonApi.TagRequest{
		Color:   config.Color.ValueStringPointer(),
		TagName: config.TagName.ValueString(),
	})
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error creating tag", err.Error())
		}
		return
	}

	config.ID = types.StringPointerValue(tagId)

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}

func (tr *tagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TagResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := readTag(ctx, tr.client, &state)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading tag", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func readTag(ctx context.Context, client *api.TagClient, resource *TagResourceModel) error {
	tagResp, err := client.Get(ctx, resource.TagName.ValueString())
	if err != nil {
		return err
	}

	resource.ID = types.StringValue(tagResp.TagId)
	resource.TagName = types.StringValue(tagResp.TagName)
	resource.Color = types.StringPointerValue(tagResp.Color)
	return nil
}

func (tr *tagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TagResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := tr.client.Update(ctx, plan.ID.ValueString(), commonApi.TagRequest{
		Color:   plan.Color.ValueStringPointer(),
		TagName: plan.TagName.ValueString(),
	})
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating tag", err.Error())
		}
		return
	}

	// wait to be safe so that the changes reflect as there is no phase to check the state
	time.Sleep(5 * time.Second)

	err = readTag(ctx, tr.client, &plan)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading tag", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (tr *tagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TagResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := tr.client.Delete(ctx, state.TagName.ValueString())
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error deleting tag", err.Error())
		}
		return
	}
}

func (tr *tagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tag_name"), req.ID)...)
}

func NewTagResource() resource.Resource {
	return &tagResource{}
}
