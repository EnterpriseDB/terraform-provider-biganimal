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
			return New("test")(), nil
		},
	}

	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"biganimal": providerserver.NewProtocol6WithError(NewProvider("debug")()),
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
