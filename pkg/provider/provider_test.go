package provider

import (
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
func testAccPreCheck(t *testing.T) {
	t.Logf("Checking BA_API_URI:%s", os.Getenv("BA_API_URI"))
	if os.Getenv("BA_API_URI") == "" {
		t.Fatal("BA_API_URI must be set for acceptance tests")
	}
	t.Logf("Checking BA_BEARER_TOKEN:%s", os.Getenv("BA_BEARER_TOKEN"))
	if os.Getenv("BA_BEARER_TOKEN") == "" {
		t.Fatal("BA_BEARER_TOKEN= must be set for acceptance tests")
	}
}
