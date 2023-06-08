package provider

import (
	"context"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type projectResource struct {
	client *api.API
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
				MarkdownDescription: "Project ID of the project.",
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
				MarkdownDescription: "User Count of the project.",
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

type cloudProvider struct {
	CloudProviderId   string `tfsdk:"cloud_provider_id"`
	CloudProviderName string `tfsdk:"cloud_provider_name"`
}

type Project struct {
	ID             types.String    `tfsdk:"id"`
	ProjectID      types.String    `tfsdk:"project_id"`
	ProjectName    types.String    `tfsdk:"project_name"`
	UserCount      types.Int64     `tfsdk:"user_count"`
	ClusterCount   types.Int64     `tfsdk:"cluster_count"`
	CloudProviders []cloudProvider `tfsdk:"cloud_providers"`
}

// Create creates the resource and sets the initial Terraform state.
func (p projectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config Project
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectId, err := p.client.ProjectClient().Create(ctx, config.ProjectName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error creating project", "Could not create project, unexpected error: "+err.Error())
		return
	}

	project, err := p.client.ProjectClient().Read(ctx, projectId)
	if err != nil {
		resp.Diagnostics.AddError("Error reading project", "Could not read project, unexpected error: "+err.Error())
		return
	}

	config.ID = types.StringValue(project.ProjectId)
	config.ProjectID = types.StringValue(project.ProjectId)
	config.ProjectName = types.StringValue(project.ProjectName)
	config.UserCount = types.Int64Value(int64(project.UserCount))
	config.ClusterCount = types.Int64Value(int64(project.ClusterCount))
	config.CloudProviders = []cloudProvider{}

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

	project, err := p.client.ProjectClient().Read(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading project", "Could not read project, unexpected error: "+err.Error())
		return
	}

	state.ProjectName = types.StringValue(project.ProjectName)
	state.UserCount = types.Int64Value(int64(project.UserCount))
	state.ClusterCount = types.Int64Value(int64(project.ClusterCount))
	if cps := project.CloudProviders; cps != nil {
		for _, provider := range cps {
			state.CloudProviders = append(state.CloudProviders, cloudProvider{
				CloudProviderId:   provider.CloudProviderId,
				CloudProviderName: provider.CloudProviderName,
			})
		}
	}

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

	_, err := p.client.ProjectClient().Update(ctx, plan.ID.ValueString(), plan.ProjectName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error updating project", "Could not update project, unexpected error: "+err.Error())
		return
	}

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

	if err := p.client.ProjectClient().Delete(ctx, state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Error deleting project", "Could not delete project, unexpected error: "+err.Error())
		return
	}
}

func NewProjectResource() resource.Resource {
	return &projectResource{}
}
