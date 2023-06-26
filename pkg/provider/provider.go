package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	frameworkschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DefaultAPIURL = "https://portal.biganimal.com/api/v3"

var (
	resourceCluster         = NewClusterResource()
	resourceAWSConnection   = NewAWSConnectionResource()
	resourceAzureConnection = NewAzureConnectionResource()
	resourceFAReplica       = NewFAReplicaResource()

	dataCluster       = NewClusterData()
	dataAWSConnection = NewAWSConnectionData()
	dataFaReplica     = NewFAReplicaData()
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	sdkschema.DescriptionKind = sdkschema.StringMarkdown
}

func NewSDKProvider(version string) func() *sdkschema.Provider {
	return func() *sdkschema.Provider {
		p := &sdkschema.Provider{
			Schema: map[string]*sdkschema.Schema{
				"ba_bearer_token": {
					Type:        sdkschema.TypeString,
					Description: "BigAnimal Bearer Token",
					Sensitive:   false,
					Optional:    true,
				},
				"ba_api_uri": {
					Type:        sdkschema.TypeString,
					Description: "BigAnimal API URL",
					Optional:    true,
				},
			},
			DataSourcesMap: map[string]*sdkschema.Resource{
				"biganimal_cluster":         dataCluster.Schema(),
				"biganimal_faraway_replica": dataFaReplica.Schema(),
				"biganimal_aws_connection":  dataAWSConnection.Schema(),
			},

			ResourcesMap: map[string]*sdkschema.Resource{
				"biganimal_cluster":          resourceCluster.Schema(),
				"biganimal_aws_connection":   resourceAWSConnection.Schema(),
				"biganimal_azure_connection": resourceAzureConnection.Schema(),
				"biganimal_faraway_replica":  resourceFAReplica.Schema(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *sdkschema.Provider) func(context.Context, *sdkschema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, schema *sdkschema.ResourceData) (any, diag.Diagnostics) {
		ba_bearer_token := schema.Get("ba_bearer_token").(string)
		ba_api_uri := schema.Get("ba_api_uri").(string)
		// set our meta to be a new api.API
		// this can be turned into concrete clients
		// by
		// api.BuildAPI(meta).ClusterClient()
		// or
		// api.BuildAPI(meta).RegionClient()

		data := &providerData{BaAPIUri: &ba_api_uri, BaBearerToken: &ba_bearer_token}
		ok, summary, detail := checkProviderConfig(data)
		if !ok {
			return nil, diag.Diagnostics{diag.Diagnostic{Severity: diag.Error, Summary: summary, Detail: detail}}
		}

		userAgent := fmt.Sprintf("%s/%s", "terraform-provider-biganimal", version)
		return api.NewAPI(ba_bearer_token, ba_api_uri, userAgent), nil
	}
}

type bigAnimalProvider struct {
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	BaBearerToken *string `tfsdk:"ba_bearer_token"`
	BaAPIUri      *string `tfsdk:"ba_api_uri"`
}

func NewFrameworkProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &bigAnimalProvider{
			version: version,
		}
	}
}

func (b bigAnimalProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "biganimal"
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

	ok, summary, detail := checkProviderConfig(&data)
	if !ok {
		resp.Diagnostics.AddError(summary, detail)
		return
	}

	userAgent := fmt.Sprintf("%s/%s", "terraform-provider-biganimal", b.version)
	client := api.NewAPI(*data.BaBearerToken, *data.BaAPIUri, userAgent)
	resp.ResourceData = client
	resp.DataSourceData = client
}

func checkProviderConfig(data *providerData) (ok bool, summary, detail string) {
	if data.BaBearerToken == nil || *data.BaBearerToken == "" {
		token := os.Getenv("BA_BEARER_TOKEN")
		data.BaBearerToken = &token
	}

	if *data.BaBearerToken == "" {
		return false, "Unable to find BA_BEARER_TOKEN", "BA_BEARER_TOKEN cannot be an empty string"
	}

	if data.BaAPIUri == nil {
		url := os.Getenv("BA_API_URI")
		data.BaAPIUri = &url
	}

	if *data.BaAPIUri == "" {
		*data.BaAPIUri = DefaultAPIURL
	}
	return true, "", ""
}

func (b bigAnimalProvider) Schema(ctx context.Context, request provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = frameworkschema.Schema{
		Attributes: map[string]frameworkschema.Attribute{
			"ba_bearer_token": frameworkschema.StringAttribute{
				MarkdownDescription: "BigAnimal Bearer Token",
				Sensitive:           false,
				Optional:            true,
			},
			"ba_api_uri": frameworkschema.StringAttribute{
				MarkdownDescription: "BigAnimal API URL",
				Optional:            true,
			},
		},
	}
}

func (b bigAnimalProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProjectsDataSource,
		NewPgdDataSource,
		NewRegionsDataSource,
	}
}

func (b bigAnimalProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProjectResource,
		NewRegionResource,
	}
}
