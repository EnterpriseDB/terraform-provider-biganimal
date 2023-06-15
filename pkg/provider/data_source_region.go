package provider

import (
	"context"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NewRegionDataSource is a helper function to simplify the provider implementation.
func NewRegionDataSource() datasource.DataSource {
	return &regionDataSource{}
}

// regionDataSource is the data source implementation.
type regionDataSource struct {
	client *api.API
}

// Configure adds the provider configured client to the data source.
func (r *regionDataSource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.API)
}

func (r *regionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_regions"
}

func (r *regionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"regions": schema.ListNestedAttribute{
				Description: "Region information.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"region_id": schema.StringAttribute{
							Description: "Region ID of the region.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Region name of the region.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Region status of the region.",
							Computed:    true,
						},
						"continent": schema.StringAttribute{
							Description: "Continent that region belongs to.",
							Computed:    true,
						},
					},
				},
			},

			"cloud_provider": schema.StringAttribute{
				Description: "Cloud provider to list the regions. For example, \"aws\" or \"azure\".",
				Required:    true,
			},
			"project_id": schema.StringAttribute{
				Description: "BigAnimal Project ID.",
				Required:    true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
			},
			"query": schema.StringAttribute{
				Description: "Query to filter region list.",
				Optional:    true,
			},
			"region_id": schema.StringAttribute{
				Description: "Unique region ID. For example, \"germanywestcentral\" in the Azure cloud provider, \"eu-west-1\" in the AWS cloud provider.",
				Optional:    true,
			},
		},
	}
}

type regionDatasource struct {
	Regions       []Region `tfsdk:"regions,omitempty"`
	CloudProvider string   `tfsdk:"cloudProvider,omitempty"`
	ProjectId     string   `tfsdk:"projectId,omitempty"`
	Query         string   `tfsdk:"query,omitempty"`
	RegionId      string   `tfsdk:"regionId,omitempty"`
}

func (r *regionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var cfg regionDatasource
	diags := req.Config.Get(ctx, &cfg)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	regions := []*models.Region{}
	if cfg.RegionId != "" {
		region, err := r.client.RegionClient().Read(ctx, cfg.ProjectId, cfg.CloudProvider, cfg.RegionId)
		if err != nil {
			resp.Diagnostics.Append(fromErr(err, "Error reading region by id: %v", cfg.RegionId)...)
			return
		}
		regions = append(regions, region)

	} else {
		respRegions, err := r.client.RegionClient().List(ctx, cfg.ProjectId, cfg.CloudProvider, cfg.Query)
		if err != nil {
			return
		}
		regions = respRegions
	}

	for _, region := range regions {
		cfg.Regions = append(cfg.Regions, Region{
			ProjectID:     cfg.ProjectId,
			CloudProvider: cfg.CloudProvider,
			RegionID:      region.Id,
			Name:          region.Name,
			Status:        region.Status,
			Continent:     region.Continent,
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &cfg)...)
}
