package provider

import (
	"context"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (

	resourceRegion    = NewRegionResource()
	resourceCluster   = NewClusterResource()
	resourceProject   = NewProjectResource()
	resourceFAReplica = NewFAReplicaResource()
	resourceAWSConnection   = NewAWSConnectionResource()
	resourceAzureConnection = NewAzureConnectionResource()


	dataRegion        = NewRegionData()
	dataCluster       = NewClusterData()
	dataProjects      = NewProjectsData()
	dataAWSConnection = NewAWSConnectionData()
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
				"ba_bearer_token": {
					Type:        schema.TypeString,
					Sensitive:   false,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("BA_BEARER_TOKEN", nil),
				},
				"ba_api_uri": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("BA_API_URI", "https://portal.biganimal.com/api/v3"),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"biganimal_cluster":        dataCluster.Schema(),
				"biganimal_region":         dataRegion.Schema(),
				"biganimal_projects":       dataProjects.Schema(),
				"biganimal_aws_connection": dataAWSConnection.Schema(),
			},

			ResourcesMap: map[string]*schema.Resource{
				"biganimal_cluster":         resourceCluster.Schema(),
				"biganimal_region":          resourceRegion.Schema(),
				"biganimal_project":         resourceProject.Schema(),
				"biganimal_faraway_replica": resourceFAReplica.Schema(),
				"biganimal_aws_connection":   resourceAWSConnection.Schema(),
				"biganimal_azure_connection": resourceAzureConnection.Schema(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {

	return func(ctx context.Context, schema *schema.ResourceData) (any, diag.Diagnostics) {
		ba_bearer_token := schema.Get("ba_bearer_token").(string)
		ba_api_uri := schema.Get("ba_api_uri").(string)
		// set our meta to be a new api.API
		// this can be turned into concrete clients
		// by
		// api.BuildAPI(meta).ClusterClient()
		// or
		// api.BuildAPI(meta).RegionClient()

		userAgent := fmt.Sprintf("%s/%s", "terraform-provider-biganimal", version)
		return api.NewAPI(ba_bearer_token, ba_api_uri, userAgent), nil
	}
}
