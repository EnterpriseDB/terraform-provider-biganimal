package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	resourceCluster = NewClusterResource()
	dataCluster     = NewClusterData()
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"ba_token": {
					Type:        schema.TypeString,
					Sensitive:   false,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("BA_BEARER_TOKEN", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"biganimal_cluster": dataCluster.Schema(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"biganimal_cluster": resourceCluster.Schema(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, schema *schema.ResourceData) (any, diag.Diagnostics) {
		// set our meta to be a new api.API
		// this can be turned into concrete clients
		// by
		// api.BuildAPI(meta).ClusterClient()
		// or
		// api.BuildAPI(meta).RegionClient()
		return api.NewAPI(), nil
	}
}
