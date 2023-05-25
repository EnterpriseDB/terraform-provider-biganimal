package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"biganimal": func() (*schema.Provider, error) {
			return New("test")(), nil
		},
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
//
// TODO: document the usage in a separate ACCEPTANCE_TESTING.md somewhere in the repo
func testAccResourcePreCheck(t *testing.T, resource_type string, checklist []string) {
	t.Logf("Checking env variables for the %s resource", resource_type)
	for _, envVar := range checklist {
		if os.Getenv(envVar) == "" {
			// Question:
			// also consider `t.Skip` here in some cases. It might make more sense than t.Fatal
			//   or, alternatively, we can implement them both, in case needed.
			t.Fatalf(fmt.Sprintf("%s must be set to run the %s acceptance tests", envVar, resource_type))
		}
	}
}

func testAccPreCheck(t *testing.T) {
	// Question:
	// Maybe we can use BA_TF_ACC_API_URI env var or something similar to that for the Acceptance tests.
	// same goes for BA_TF_ACC_BEARER_TOKEN
	if os.Getenv("BA_API_URI") == "" {
		t.Fatal("BA_API_URI must be set for acceptance tests")
	}
	if os.Getenv("BA_BEARER_TOKEN") == "" {
		t.Fatal("BA_BEARER_TOKEN= must be set for acceptance tests")
	}
}
