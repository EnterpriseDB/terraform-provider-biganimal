package provider_test

import (
	"fmt"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccResourceRegion(t *testing.T) {
	var (
		projectID = envForResourceVar("region", "project")
		provider  = envForResourceVar("region", "provider")
		regionID  = envForResourceVar("region", "region")

		regionConfig = `resource "biganimal_region" "this" {
  status 		 = "%s"
  project_id     = "%s"
  cloud_provider = "%s"
  region_id      = "%s"
}`
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(regionConfig, api.REGION_ACTIVE, projectID, provider, regionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_region.this", "status", api.REGION_ACTIVE),
				),
			},
			{
				Config: fmt.Sprintf(regionConfig, api.REGION_SUSPENDED, projectID, provider, regionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_region.this", "status", api.REGION_SUSPENDED),
				),
			}, {
				Config: fmt.Sprintf(regionConfig, api.REGION_INACTIVE, projectID, provider, regionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_region.this", "status", api.REGION_INACTIVE),
				),
			},
		},
	})

}
