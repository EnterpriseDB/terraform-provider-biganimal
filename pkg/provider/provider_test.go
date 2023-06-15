package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	testAccProviderFactories        map[string]func() (*schema.Provider, error)
	testAccProtoV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error)
)

func init() {
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"biganimal": func() (*schema.Provider, error) {
			return NewSDKProvider("test")(), nil
		},
	}

	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"biganimal": providerserver.NewProtocol6WithError(NewFrameworkProvider("test")()),
	}

}

// The following check can be added as a Resource specific PreCheck during Acceptance Tests
// A similar usage would be something like:
//
//		resource.Test(t, resource.TestCase{
//			PreCheck: func() {
//				testAccPreCheck(t)
//				testAccResourcePreCheck(t, "cluster", acc_env_vars_checklist)
//			},
//	  ...
func testAccResourcePreCheck(t *testing.T, resource_type string, checklist []string) {
	t.Logf("Checking env variables for the %s resource", resource_type)
	for _, envVar := range checklist {
		if os.Getenv(envVar) == "" {
			t.Fatalf(fmt.Sprintf("%s must be set to run the %s acceptance tests", envVar, resource_type))
		}
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("BA_API_URI") == "" {
		t.Fatal("BA_API_URI must be set for acceptance tests")
	}
	if os.Getenv("BA_BEARER_TOKEN") == "" {
		t.Fatal("BA_BEARER_TOKEN must be set for acceptance tests")
	}
}

func Test_checkProviderConfig(t *testing.T) {
	configToken := "config_token"
	configUrl := "config_url"

	type args struct {
		data     *providerData
		envURL   bool
		envToken bool
	}
	tests := []struct {
		name      string
		args      args
		wantOk    bool
		wantToken string
		wantUrl   string
	}{
		{
			name: "failure due to an empty token",
			args: args{data: &providerData{}},
		},

		{
			name:      "From environment variables",
			args:      args{envToken: true, envURL: true, data: &providerData{}},
			wantOk:    true,
			wantToken: "env_token",
			wantUrl:   "env_url",
		},

		{
			name:      "From configuration",
			args:      args{envToken: true, envURL: true, data: &providerData{BaAPIUri: &configUrl, BaBearerToken: &configToken}},
			wantOk:    true,
			wantUrl:   configUrl,
			wantToken: configToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envToken := ""
			if tt.args.envToken {
				envToken = "env_token"
			}
			if err := os.Setenv("BA_BEARER_TOKEN", envToken); err != nil {
				t.Fatal(err)
			}

			envURL := ""
			if tt.args.envURL {
				envURL = "env_url"
			}
			if err := os.Setenv("BA_API_URI", envURL); err != nil {
				t.Fatal(err)
			}

			gotOk, _, _ := checkProviderConfig(tt.args.data)
			if gotOk != tt.wantOk {
				t.Errorf("checkProviderConfig() gotOk = %v, want %v", gotOk, tt.wantOk)
			}

			if gotOk {
				gotUrl := *tt.args.data.BaAPIUri
				if gotUrl != tt.wantUrl {
					t.Errorf("checkProviderConfig() gotURL = %v, want %v", gotUrl, tt.wantUrl)
				}

				gotToken := *tt.args.data.BaBearerToken
				if gotToken != tt.wantToken {
					t.Errorf("checkProviderConfig() gotToken = %v, want %v", gotToken, tt.wantToken)
				}

			}
		})
	}
}
