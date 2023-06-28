package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// NewRegionDataSource is a helper function to simplify the provider implementation.
func NewRegionDataSource() datasource.DataSource {
	return &regionDataSource{}
}

// regionDataSource is the data source implementation.
type regionDataSource struct {
	regionsDataSource
}

func (r *regionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

func (r *regionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.regionsDataSource.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "The datasource 'region' is deprecated and will be removed in the next major version. Please use 'regions' instead."
}
