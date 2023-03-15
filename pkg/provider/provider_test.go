package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"biganimal": func() (*schema.Provider, error) {
			return New("test")(), nil
		},
	}

}
func testAccPreCheck(t *testing.T) {
	if os.Getenv("BA_API_URI") == "" {
		t.Fatal("BA_API_URI must be set for acceptance tests")
	}
	if os.Getenv("BA_BEARER_TOKEN") == "" {
		t.Fatal("BA_BEARER_TOKEN= must be set for acceptance tests")
	}
}
