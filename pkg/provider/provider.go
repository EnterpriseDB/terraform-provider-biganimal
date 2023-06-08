package provider

import (
	"context"
	"fmt"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	schema2 "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
)

var (
	resourceRegion          = NewRegionResource()
	resourceCluster         = NewClusterResource()
	resourceProject         = NewProjectResource()
	resourceAWSConnection   = NewAWSConnectionResource()
	resourceAzureConnection = NewAzureConnectionResource()
	resourceFAReplica       = NewFAReplicaResource()

	dataRegion        = NewRegionData()
	dataCluster       = NewClusterData()
	dataAWSConnection = NewAWSConnectionData()
	dataFaReplica     = NewFAReplicaData()
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown
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
				"biganimal_cluster":         dataCluster.Schema(),
				"biganimal_region":          dataRegion.Schema(),
				"biganimal_faraway_replica": dataFaReplica.Schema(),
				"biganimal_aws_connection":  dataAWSConnection.Schema(),
			},

			ResourcesMap: map[string]*schema.Resource{
				"biganimal_cluster":          resourceCluster.Schema(),
				"biganimal_region":           resourceRegion.Schema(),
				"biganimal_project":          resourceProject.Schema(),
				"biganimal_aws_connection":   resourceAWSConnection.Schema(),
				"biganimal_azure_connection": resourceAzureConnection.Schema(),
				"biganimal_faraway_replica":  resourceFAReplica.Schema(),
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
		// this can be turned into concrete cba_api_uri := schema.Get("ba_api_uri").(string)lients
		// by
		// api.BuildAPI(meta).ClusterClient()
		// or
		// api.BuildAPI(meta).RegionClient()

		userAgent := fmt.Sprintf("%s/%s", "terraform-provider-biganimal", version)
		return api.NewAPI(ba_bearer_token, ba_api_uri, userAgent), nil
	}
}

type bigAnimalProvider struct {
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	BaBearerToken string `tfsdk:"ba_bearer_token"`
	BaAPIUri      string `tfsdk:"ba_api_uri"`
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &bigAnimalProvider{
			version: version,
		}
	}
}

func (b bigAnimalProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "bigAnimal"
	resp.Version = b.version
}

func (b bigAnimalProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// get providerData
	var data providerData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	var token = os.Getenv("BA_BEARER_TOKEN")
	if data.BaBearerToken != "" {
		token = data.BaBearerToken
	}

	if token == "" {
		resp.Diagnostics.AddError(
			"Unable to find ba_nearer_token",
			"ba_nearer_token cannot be an empty string",
		)
		return
	}

	var host = "https://portal.biganimal.com/api/v3"
	if os.Getenv("BA_API_URI") != "" {
		host = os.Getenv("BA_API_URI")
	}
	if data.BaAPIUri != "" {
		host = data.BaAPIUri
	}

	userAgent := fmt.Sprintf("%s/%s", "terraform-provider-biganimal", b.version)
	client := api.NewAPI(token, host, userAgent)
	resp.ResourceData = client
	resp.DataSourceData = client
}

func (b bigAnimalProvider) Schema(ctx context.Context, request provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema2.Schema{
		Attributes: map[string]schema2.Attribute{
			"ba_bearer_token": schema2.StringAttribute{
				MarkdownDescription: "BigAnimal Bearer Token",
				Sensitive:           false,
				Required:            true,
			},
			"ba_api_uri": schema2.StringAttribute{
				MarkdownDescription: "BigAnimal API URL",
				Optional:            true,
			},
		},
	}
}

func (b bigAnimalProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProjectsDataSource,
	}
}

func (b bigAnimalProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}
