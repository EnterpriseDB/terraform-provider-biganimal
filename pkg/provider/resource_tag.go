package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
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
	TagId   types.String `tfsdk:"tag_id"`
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
		MarkdownDescription: "The PGD cluster data source describes a BigAnimal cluster. The data source requires your PGD cluster name.",
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
			"tag_id": schema.StringAttribute{
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
	config.TagId = types.StringPointerValue(tagId)

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}

func (tr *tagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TagResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := tr.read(ctx, &state)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading tag", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (tr *tagResource) read(ctx context.Context, resource *TagResourceModel) error {
	tagResp, err := tr.client.Get(ctx, resource.TagId.ValueString())
	if err != nil {
		return err
	}

	resource.ID = types.StringValue(tagResp.TagId)
	resource.TagId = types.StringValue(tagResp.TagId)
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

	_, err := tr.client.Update(ctx, plan.TagId.ValueString(), commonApi.TagRequest{
		Color:   plan.Color.ValueStringPointer(),
		TagName: plan.TagName.ValueString(),
	})
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating tag", err.Error())
		}
		return
	}

	err = tr.read(ctx, &plan)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading tag", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (tr *tagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (tr *tagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
}

func NewTagResource() resource.Resource {
	return &tagResource{}
}
