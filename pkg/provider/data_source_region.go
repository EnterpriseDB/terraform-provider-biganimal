package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigure = &regionsDataSource{}

// NewRegionsDataSource is a helper function to simplify the provider implementation.
func NewRegionsDataSource() datasource.DataSource {
	return &regionsDataSource{}
}

// regionsDataSource is the data source implementation.
type regionsDataSource struct {
	client *api.RegionClient
}

// Configure adds the provider configured client to the data source.
func (r *regionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.API).RegionClient()
}

func (r *regionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

func (r *regionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Datasource ID.",
				Computed:    true,
			},
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
				Description: "Cloud provider to list the regions. For example, \"aws\", \"azure\" or \"bah:aws\".",
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
			"tags": schema.SetNestedAttribute{
				Description: "show tags associated with this resource",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"tag_id": schema.StringAttribute{
							Computed: true,
						},
						"tag_name": schema.StringAttribute{
							Computed: true,
						},
						"color": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type regionsDataSourceModel struct {
	ID            *string          `tfsdk:"id"`
	ProjectId     *string          `tfsdk:"project_id"`
	CloudProvider *string          `tfsdk:"cloud_provider"`
	RegionId      *string          `tfsdk:"region_id"`
	Query         types.String     `tfsdk:"query"`
	Regions       []*models.Region `tfsdk:"regions"`
}

func (r *regionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var cfg regionsDataSourceModel
	diags := req.Config.Get(ctx, &cfg)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	regions := []*models.Region{}
	if cfg.RegionId != nil {
		region, err := r.client.Read(ctx, *cfg.ProjectId, *cfg.CloudProvider, *cfg.RegionId)
		if err != nil {
			if appendDiagFromBAErr(err, &diags) {
				return
			}
			diags.AddError(fmt.Sprintf("Error reading region by id: %q", *cfg.RegionId), err.Error())
			return
		}
		regions = append(regions, region)

	} else {
		respRegions, err := r.client.List(ctx, *cfg.ProjectId, *cfg.CloudProvider, cfg.Query.ValueString())
		if err != nil {
			if appendDiagFromBAErr(err, &diags) {
				return
			}
			diags.AddError(fmt.Sprintf("Error reading region by query: %q", cfg.Query.ValueString()), err.Error())
			return
		}
		regions = respRegions
	}

	cfg.Regions = append(cfg.Regions, regions...)
	resourceID := strconv.FormatInt(time.Now().Unix(), 10)
	cfg.ID = &resourceID
	resp.Diagnostics.Append(resp.State.Set(ctx, &cfg)...)
}
