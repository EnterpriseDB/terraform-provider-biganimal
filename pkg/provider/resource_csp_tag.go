package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &cSPTagResource{}
	_ resource.ResourceWithConfigure = &cSPTagResource{}
)

type CSPTagResourceModel struct {
	ID              types.String `tfsdk:"id"`
	ProjectID       types.String `tfsdk:"project_id"`
	CloudProviderID types.String `tfsdk:"cloud_provider_id"`
	AddTags         []addTag     `tfsdk:"add_tags"`
	DeleteTags      types.List   `tfsdk:"delete_tags"`
	EditTags        []CSPTag     `tfsdk:"edit_tags"`
	CSPTags         types.List   `tfsdk:"csp_tags"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

type addTag struct {
	CspTagKey   types.String `tfsdk:"csp_tag_key"`
	CspTagValue types.String `tfsdk:"csp_tag_value"`
}

type CSPTag struct {
	CSPTagID    types.String `tfsdk:"csp_tag_id"`
	CSPTagKey   types.String `tfsdk:"csp_tag_key"`
	CSPTagValue types.String `tfsdk:"csp_tag_value"`
	Status      types.String `tfsdk:"status"`
}

type cSPTagResource struct {
	client *api.CSPTagClient
}

func (tr *cSPTagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	tr.client = req.ProviderData.(*api.API).CSPTagClient()
}

func (tr *cSPTagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csp_tag"
}

func (tf *cSPTagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CSP Tags will enable users to categorize and organize resources across types and improve the efficiency of resource retrieval",
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true}),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"project_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cloud_provider_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"add_tags": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"csp_tag_key": schema.StringAttribute{
							Required: true,
						},
						"csp_tag_value": schema.StringAttribute{
							Required: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"delete_tags": schema.ListAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"edit_tags": schema.ListNestedAttribute{
				Description: "",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"csp_tag_id": schema.StringAttribute{
							Required: true,
						},
						"csp_tag_key": schema.StringAttribute{
							Required: true,
						},
						"csp_tag_value": schema.StringAttribute{
							Required: true,
						},
						"status": schema.StringAttribute{
							Required: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"csp_tags": schema.ListNestedAttribute{
				Description: "",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"csp_tag_id": schema.StringAttribute{
							Computed: true,
						},
						"csp_tag_key": schema.StringAttribute{
							Computed: true,
						},
						"csp_tag_value": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (tr *cSPTagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config CSPTagResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cSPTagRequest := commonApi.CSPTagRequest{}
	cSPTagRequest.AddTags = []commonApi.AddTag{}
	cSPTagRequest.DeleteTags = []string{}
	cSPTagRequest.EditTags = []commonApi.EditTag{}

	for _, addTag := range config.AddTags {
		cSPTagRequest.AddTags = append(cSPTagRequest.AddTags, commonApi.AddTag{
			CspTagKey:   addTag.CspTagKey.ValueString(),
			CspTagValue: addTag.CspTagValue.ValueString(),
		})
	}

	for _, deleteTag := range config.DeleteTags.Elements() {
		cSPTagRequest.DeleteTags = append(cSPTagRequest.DeleteTags, deleteTag.(basetypes.StringValue).ValueString())
	}

	for _, editTag := range config.EditTags {
		cSPTagRequest.EditTags = append(cSPTagRequest.EditTags, commonApi.EditTag{
			CSPTagID:    editTag.CSPTagID.ValueString(),
			CSPTagKey:   editTag.CSPTagKey.ValueString(),
			CSPTagValue: editTag.CSPTagValue.ValueString(),
		})
	}

	_, err := tr.client.Put(ctx, config.ProjectID.ValueString(), config.CloudProviderID.ValueString(), cSPTagRequest)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error creating tag", err.Error())
		}
		return
	}

	config.ID = types.StringValue(fmt.Sprintf("%s/%s", config.ProjectID.ValueString(), config.CloudProviderID.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}

func (tr *cSPTagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CSPTagResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := readCSPTag(ctx, tr.client, &state)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading tag", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func readCSPTag(ctx context.Context, client *api.CSPTagClient, resource *CSPTagResourceModel) error {
	cSPTagResp, err := client.Get(ctx, resource.ProjectID.ValueString(), resource.CloudProviderID.ValueString())
	if err != nil {
		return err
	}

	resource.ID = types.StringValue(fmt.Sprintf("%s/%s", resource.ProjectID.ValueString(), resource.CloudProviderID.ValueString()))
	cSPTagsElems := []attr.Value{}
	for _, respCSPTag := range cSPTagResp.Data {
		cSPTagsElems = append(
			cSPTagsElems, types.ObjectValueMust(resource.CSPTags.ElementType(ctx).(types.ObjectType).AttributeTypes(),
				map[string]attr.Value{
					"csp_tag_id":    types.StringValue(respCSPTag.CSPTagID),
					"csp_tag_key":   types.StringValue(respCSPTag.CSPTagKey),
					"csp_tag_value": types.StringValue(respCSPTag.CSPTagValue),
					"status":        types.StringValue(respCSPTag.Status),
				},
			))
	}

	resource.CSPTags = types.ListValueMust(resource.CSPTags.ElementType(ctx), cSPTagsElems)
	return nil
}

func (tr *cSPTagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CSPTagResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cSPTagRequest := commonApi.CSPTagRequest{}
	cSPTagRequest.AddTags = []commonApi.AddTag{}
	cSPTagRequest.DeleteTags = []string{}
	cSPTagRequest.EditTags = []commonApi.EditTag{}
	for _, addTag := range plan.AddTags {
		cSPTagRequest.AddTags = append(cSPTagRequest.AddTags, commonApi.AddTag{
			CspTagKey:   addTag.CspTagKey.ValueString(),
			CspTagValue: addTag.CspTagValue.ValueString(),
		})
	}
	for _, deleteTag := range plan.DeleteTags.Elements() {
		cSPTagRequest.DeleteTags = append(cSPTagRequest.DeleteTags, deleteTag.(basetypes.StringValue).ValueString())
	}
	for _, editTag := range plan.EditTags {
		cSPTagRequest.EditTags = append(cSPTagRequest.EditTags, commonApi.EditTag{
			CSPTagID:    editTag.CSPTagID.ValueString(),
			CSPTagKey:   editTag.CSPTagKey.ValueString(),
			CSPTagValue: editTag.CSPTagValue.ValueString(),
			Status:      editTag.Status.ValueString(),
		})
	}

	_, err := tr.client.Put(ctx, plan.ProjectID.ValueString(), plan.CloudProviderID.ValueString(), cSPTagRequest)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error creating tag", err.Error())
		}
		return
	}

	err = readCSPTag(ctx, tr.client, &plan)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading tag", err.Error())
		}
		return
	}

	plan.ID = types.StringValue(fmt.Sprintf("%s/%s", plan.ProjectID.ValueString(), plan.CloudProviderID.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (tr *cSPTagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// var state CSPTagResourceModel
	// diags := req.State.Get(ctx, &state)
	// resp.Diagnostics.Append(diags...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// err := tr.client.Delete(ctx, state.TagId.ValueString())
	// if err != nil {
	// 	if !appendDiagFromBAErr(err, &resp.Diagnostics) {
	// 		resp.Diagnostics.AddError("Error deleting tag", err.Error())
	// 	}
	// 	return
	// }
}

func (tr *cSPTagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: project_id/cloud_provider_id. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cloud_provider_id"), idParts[1])...)
}

func NewCSPTagResource() resource.Resource {
	return &cSPTagResource{}
}
