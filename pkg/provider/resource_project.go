package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

type Project struct {
	ID             *string               `tfsdk:"id"`
	ProjectID      *string               `tfsdk:"project_id"`
	ProjectName    *string               `tfsdk:"project_name"`
	UserCount      *int                  `tfsdk:"user_count"`
	ClusterCount   *int                  `tfsdk:"cluster_count"`
	CloudProviders types.Set             `tfsdk:"cloud_providers"`
	Tags           []commonTerraform.Tag `tfsdk:"tags"`
}

type projectResource struct {
	client *api.ProjectClient
}

func (p projectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (p projectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The project resource is used to manage projects in your organization. " +
			"See [Managing projects](https://www.enterprisedb.com/docs/biganimal/latest/administering_cluster/projects/) for more details.\n\n" +
			"Newly created projects are not automatically connected to your cloud providers. " +
			"Please visit [Connecting your cloud](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/02_connecting_to_your_cloud/) for more details.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Resource ID of the project.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "Project ID of the project.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				DeprecationMessage: "The usage of 'project_id' is no longer recommended and has been deprecated. We suggest using 'id' instead.",
			},
			"project_name": schema.StringAttribute{
				MarkdownDescription: "Project Name of the project.",
				Required:            true,
			},
			"user_count": schema.Int64Attribute{
				MarkdownDescription: "User Count of the project.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"cluster_count": schema.Int64Attribute{
				MarkdownDescription: "Cluster Count of the project.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},

			// We don't have a mechanism to automate the csp connection right now
			// So, the `cloud_providers` value is computed only.
			"cloud_providers": schema.SetNestedAttribute{
				MarkdownDescription: "Enabled Cloud Providers.",
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					PlanModifiers: []planmodifier.Object{
						objectplanmodifier.UseStateForUnknown(),
					},
					Attributes: map[string]schema.Attribute{
						"cloud_provider_id": schema.StringAttribute{
							MarkdownDescription: "Cloud Provider ID.",
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"cloud_provider_name": schema.StringAttribute{
							MarkdownDescription: "Cloud Provider Name.",
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				Description:   "Assign existing tags or create tags to assign to this resource",
				Optional:      true,
				Computed:      true,
				NestedObject:  ResourceTagNestedObject,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

// modify plan on at runtime
func (p *projectResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	ValidateTags(ctx, p.client.TagClient(), req, resp)
}

// Configure adds the provider configured client to the data source.
func (p *projectResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p.client = req.ProviderData.(*api.API).ProjectClient()
}

// Create creates the resource and sets the initial Terraform state.
func (p projectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config Project
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectReqModel := models.Project{
		ProjectName: *config.ProjectName,
		Tags:        buildApiReqTags(config.Tags),
	}

	projectId, err := p.client.Create(ctx, projectReqModel)
	if err != nil {
		resp.Diagnostics.AddError("Error creating project", "Could not create project, unexpected error: "+err.Error())
		return
	}

	project, err := p.client.Read(ctx, projectId)
	if err != nil {
		resp.Diagnostics.AddError("Error reading project", "Could not read project, unexpected error: "+err.Error())
		return
	}

	config.ID = &project.ProjectId
	config.ProjectID = &project.ProjectId
	config.ProjectName = &project.ProjectName
	config.UserCount = &project.UserCount
	config.ClusterCount = &project.ClusterCount
	config.CloudProviders = BuildTfRsrcCloudProviders(project.CloudProviders)

	buildTfRsrcTagsAs(&config.Tags, project.Tags)

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (p projectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Project
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	project, err := p.client.Read(ctx, *state.ID)
	if err != nil {
		resp.Diagnostics.AddError("Error reading project", "Could not read project, unexpected error: "+err.Error())
		return
	}

	state.ProjectID = &project.ProjectId
	state.ProjectName = &project.ProjectName
	state.UserCount = &project.UserCount
	state.ClusterCount = &project.ClusterCount
	state.CloudProviders = BuildTfRsrcCloudProviders(project.CloudProviders)

	buildTfRsrcTagsAs(&state.Tags, project.Tags)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (p projectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan Project
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectReqModel := models.Project{
		ProjectName: *plan.ProjectName,
		Tags:        buildApiReqTags(plan.Tags),
	}

	_, err := p.client.Update(ctx, *plan.ProjectID, projectReqModel)
	if err != nil {
		resp.Diagnostics.AddError("Error updating project", "Could not update project, unexpected error: "+err.Error())
		return
	}

	buildTfRsrcTagsAs(&plan.Tags, projectReqModel.Tags)

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (p projectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state Project
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := p.client.Delete(ctx, *state.ProjectID); err != nil {
		resp.Diagnostics.AddError("Error deleting project", "Could not delete project, unexpected error: "+err.Error())
		return
	}
}

func (p projectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewProjectResource() resource.Resource {
	return &projectResource{}
}
