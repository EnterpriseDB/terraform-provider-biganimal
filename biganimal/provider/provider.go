package provider

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				"ba_token": &schema.Schema{
					Type:        schema.TypeString,
					Sensitive:   false,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("BA_BEARER_TOKEN", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"biganimal_data_source": dataSourceScaffolding(),
			},
			ResourcesMap: map[string]*schema.Resource{
				//"scaffolding_resource": resourceScaffolding(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
	BaURL      string
	BaToken    string
	HTTPClient *http.Client
}

func NewClient(ba_token string) (*apiClient, error) {
	c := apiClient{
		BaURL:      "https://portal.biganimal.com/api/v2",
		BaToken:    ba_token,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
	return &c, nil
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, schema *schema.ResourceData) (any, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent

		ba_token := schema.Get("ba_token").(string)

		client, _ := NewClient(ba_token)
		return client, nil
	}
}
