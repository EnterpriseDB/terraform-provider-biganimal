package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource              = &tagDataSource{}
	_ datasource.DataSourceWithConfigure = &tagDataSource{}
)

type tagDatasourceModel struct {
	TagResourceModel
}

type tagDataSource struct {
	client *api.TagClient
}

func (d tagDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

// Configure adds the provider configured client to the data source.
func (d *tagDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*api.API).TagClient()
}

func (d tagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tags will enable users to categorize and organize resources across types and improve the efficiency of resource retrieval",
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true}),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"tag_id": schema.StringAttribute{
				Required: true,
			},
			"tag_name": schema.StringAttribute{
				Computed: true,
			},
			"color": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *tagDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data tagDatasourceModel
	diags := req.Config.Get(ctx, &data.TagResourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := readTag(ctx, d.client, &data.TagResourceModel); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading tag", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data.TagResourceModel)...)
}

func NewTagDataSource() datasource.DataSource {
	return &tagDataSource{}
}
