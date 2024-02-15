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
	resourceAWSConnection   = NewAWSConnectionResource()
	resourceAzureConnection = NewAzureConnectionResource()
	resourceFAReplica       = NewFAReplicaResource()

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
				"edb_tf_access_key": {
					Type:        sdkschema.TypeString,
					Description: "BigAnimal Access Key",
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
				"biganimal_faraway_replica": dataFaReplica.Schema(),
				"biganimal_aws_connection":  dataAWSConnection.Schema(),
			},

			ResourcesMap: map[string]*sdkschema.Resource{
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
		// If the credential data are provided inside a provider block, get them first
		// If they are not provided, the schema_* credentials will be empty strings
		schema_ba_bearer_token := schema.Get("ba_bearer_token").(string)
		schema_edb_tf_access_key := schema.Get("edb_tf_access_key").(string)
		schema_ba_api_uri := schema.Get("ba_api_uri").(string)

		data := &providerData{BaAPIUri: &schema_ba_api_uri, BaBearerToken: &schema_ba_bearer_token, EdbTFAccessKey: &schema_edb_tf_access_key}
		ok, summary, detail := checkProviderConfig(data)
		if !ok {
			return nil, diag.Diagnostics{diag.Diagnostic{Severity: diag.Error, Summary: summary, Detail: detail}}
		}

		userAgent := fmt.Sprintf("%s/%s", "terraform-provider-biganimal", version)
		return api.NewAPI(*data.EdbTFAccessKey, *data.BaBearerToken, *data.BaAPIUri, userAgent), nil
	}
}

type bigAnimalProvider struct {
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	BaBearerToken  *string `tfsdk:"ba_bearer_token"`
	EdbTFAccessKey *string `tfsdk:"edb_tf_access_key"`
	BaAPIUri       *string `tfsdk:"ba_api_uri"`
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
	client := api.NewAPI(*data.EdbTFAccessKey, *data.BaBearerToken, *data.BaAPIUri, userAgent)
	resp.ResourceData = client
	resp.DataSourceData = client
}

// Checks if the providerData is set.
// If not, checks if the environment variables are set or throws an error
func checkProviderConfig(data *providerData) (ok bool, summary, detail string) {
	if data.BaBearerToken == nil || *data.BaBearerToken == "" {
		token := os.Getenv("BA_BEARER_TOKEN")
		data.BaBearerToken = &token
	}

	// access key environment variable takes precedence over access key schema field
	accessKey := os.Getenv("EDB_TF_ACCESS_KEY")
	if accessKey != "" {
		data.EdbTFAccessKey = &accessKey
	} else if data.EdbTFAccessKey == nil || *data.EdbTFAccessKey == "" {
		data.EdbTFAccessKey = &accessKey
	}

	if *data.EdbTFAccessKey == "" && *data.BaBearerToken == "" {
		return false, "Unable to find EDB_TF_ACCESS_KEY or BA_BEARER_TOKEN", "EDB_TF_ACCESS_KEY and BA_BEARER_TOKEN both cannot be an empty string"
	}

	if data.BaAPIUri == nil || *data.BaAPIUri == "" {
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
			"edb_tf_access_key": frameworkschema.StringAttribute{
				MarkdownDescription: "BigAnimal Access Key",
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
		NewClusterDataSource,
		NewPgdDataSource,
		NewRegionsDataSource,
	}
}

func (b bigAnimalProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProjectResource,
		NewPgdResource,
		NewRegionResource,
		NewClusterResource,
	}
}
