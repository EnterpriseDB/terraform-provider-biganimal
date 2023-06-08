package provider

import (
	"context"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type projectResource struct {
	client *api.API
}

func (p projectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (p projectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The project resource is used to manage projects in your organization. " +
			"See [Managing projects](https://www.enterprisedb.com/docs/biganimal/latest/administering_cluster/projects/) for more details.\n\n" +
			"Newly created projects are not automatically connected to your cloud providers. " +
			"Please visit [Connecting your cloud](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/02_connecting_to_your_cloud/) for more details.",
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Description: "Project ID of the project.",
				Computed:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Project Name of the project.",
				Required:    true,
			},
			"user_count": schema.Int64Attribute{
				Description: "User Count of the project.",
				Computed:    true,
			},
			"cluster_count": schema.Int64Attribute{
				Description: "User Count of the project.",
				Computed:    true,
			},

			// We don't have a mechanism to automate the csp connection right now
			// So, the `cloud_providers` value is computed only.
			"cloud_providers": schema.SetNestedAttribute{
				Description: "Enabled Cloud Providers.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cloud_provider_id": schema.StringAttribute{
							Description: "Cloud Provider ID.",
							Computed:    true,
						},
						"cloud_provider_name": schema.StringAttribute{
							Description: "Cloud Provider Name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (r *projectResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.API)
}

// Create creates the resource and sets the initial Terraform state.
func (p projectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.Project
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectId, err := p.client.ProjectClient().Create(ctx, plan.ProjectName)
	if err != nil {
		resp.Diagnostics.AddError("Error creating project", "Could not create project, unexpected error: "+err.Error())
		return
	}

	project, err := p.client.ProjectClient().Read(ctx, projectId)
	if err != nil {
		resp.Diagnostics.AddError("Error reading project", "Could not read project, unexpected error: "+err.Error())
		return
	}

	diags = resp.State.Set(ctx, project)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (p projectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.Project
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	project, err := p.client.ProjectClient().Read(ctx, state.ProjectId)
	if err != nil {
		resp.Diagnostics.AddError("Error reading project", "Could not read project, unexpected error: "+err.Error())
		return
	}

	diags = resp.State.Set(ctx, project)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (p projectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.Project
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := p.client.ProjectClient().Update(ctx, plan.ProjectId, plan.ProjectName)
	if err != nil {
		resp.Diagnostics.AddError("Error updating project", "Could not update project, unexpected error: "+err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (p projectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state models.Project
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := p.client.ProjectClient().Delete(ctx, state.ProjectId); err != nil {
		resp.Diagnostics.AddError("Error deleting project", "Could not delete project, unexpected error: "+err.Error())
		return
	}
}

func NewProjectResource() resource.Resource {
	return &projectResource{}
}
