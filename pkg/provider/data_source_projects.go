package provider

import (
	"context"
	"fmt"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type projectsDataSource struct {
	client *api.API
}

func (p projectsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_projects"

}

// Configure adds the provider configured client to the data source.
func (d *projectsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*api.API)
}

type projectsDataSourceData struct {
	Query    *string           `tfsdk:"query"`
	Projects []*models.Project `tfsdk:"projects"`
}

func (p projectsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The projects data source shows the BigAnimal Projects.",
		Attributes: map[string]schema.Attribute{
			"projects": schema.SetNestedAttribute{
				Description: "List of the organization's projects.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
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
						"cloud_providers": schema.SetNestedAttribute{
							Description: "Enabled Cloud Providers.",
							Computed:    true,
							Optional:    true,
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
				},
			},
			"query": schema.StringAttribute{
				Description: "Query to filter project list.",
				Optional:    true,
			},
		},
	}
}

func (p projectsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data projectsDataSourceData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "read projects data source")
	var query string
	if data.Query != nil {
		query = *data.Query
	}
	list, err := p.client.ProjectClient().List(ctx, query)
	if err != nil {
		resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Unable to call read project, got error: %s", err))
		return
	}

	data.Projects = list
	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func NewProjectsDataSource() datasource.DataSource {
	return &projectsDataSource{}
}
