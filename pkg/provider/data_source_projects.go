package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &projectsDataSource{}
	_ datasource.DataSourceWithConfigure = &projectsDataSource{}
)

type projectsDataSource struct {
	client *api.ProjectClient
}

func (p projectsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_projects"
}

// Configure adds the provider configured client to the data source.
func (p *projectsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p.client = req.ProviderData.(*api.API).ProjectClient()
}

type projectsDataSourceData struct {
	ID       *string    `tfsdk:"id"`
	Query    *string    `tfsdk:"query"`
	Projects []*Project `tfsdk:"projects"`
}

func (p projectsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The projects data source shows the BigAnimal Projects.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Datasource ID.",
				Computed:    true,
			},
			"projects": schema.SetNestedAttribute{
				Description: "List of the organization's projects.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Resource ID of the project.",
							Computed:            true,
						},
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
						"tags": schema.SetNestedAttribute{
							Description:  "Show existing tags associated with this resource",
							Optional:     true,
							Computed:     true,
							NestedObject: DataSourceTagNestedObject,
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
	list, err := p.client.List(ctx, query)
	if err != nil {
		resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Unable to call read project, got error: %s", err))
		return
	}

	for _, project := range list {
		appendProj := &Project{
			ID:           &project.ProjectId,
			ProjectID:    &project.ProjectId,
			ProjectName:  &project.ProjectName,
			UserCount:    &project.UserCount,
			ClusterCount: &project.ClusterCount,
		}

		appendProj.CloudProviders = BuildTfRsrcCloudProviders(project.CloudProviders)

		tags := []commonTerraform.Tag{}
		for _, tag := range project.Tags {
			tags = append(tags, commonTerraform.Tag{
				TagName: types.StringValue(tag.TagName),
				Color:   types.StringPointerValue(tag.Color),
			})
		}
		appendProj.Tags = tags

		data.Projects = append(data.Projects, appendProj)
	}

	resourceID := strconv.FormatInt(time.Now().Unix(), 10)
	data.ID = &resourceID
	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func NewProjectsDataSource() datasource.DataSource {
	return &projectsDataSource{}
}
