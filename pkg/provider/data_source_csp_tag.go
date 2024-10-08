package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &cSPTagDataSource{}
	_ datasource.DataSourceWithConfigure = &cSPTagDataSource{}
)

type cSPTagDatasourceModel struct {
	CSPTagResourceModel
}

type cSPTagDataSource struct {
	client *api.CSPTagClient
}

func (c *cSPTagDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csp_tag"
}

// Configure adds the provider configured client to the data source.
func (c *cSPTagDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c.client = req.ProviderData.(*api.API).CSPTagClient()
}

func (c *cSPTagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CSP Tags will enable users to categorize and organize resources across types and improve the efficiency of resource retrieval",
		// using Blocks for backward compatible
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true},
			),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"project_id": schema.StringAttribute{
				Required: true,
			},
			"cloud_provider_id": schema.StringAttribute{
				Required: true,
			},
			"add_tags": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"csp_tag_key": schema.StringAttribute{
							Computed: true,
						},
						"csp_tag_value": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"delete_tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"edit_tags": schema.ListNestedAttribute{
				Computed: true,
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
			"csp_tags": schema.ListNestedAttribute{
				Description: "CSP Tags on cluster",
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

func (c *cSPTagDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data cSPTagDatasourceModel
	diags := req.Config.Get(ctx, &data.CSPTagResourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := readCSPTag(ctx, c.client, &data.CSPTagResourceModel); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data.CSPTagResourceModel)...)
}

func NewCSPTagDataSource() datasource.DataSource {
	return &cSPTagDataSource{}
}
