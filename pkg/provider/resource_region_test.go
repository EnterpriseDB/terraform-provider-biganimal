package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceRegion_basic(t *testing.T) {
	var (
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_region_project_id",
			"BA_TF_ACC_VAR_region_provider",
			"BA_TF_ACC_VAR_region_region_id",
		}

		projectID = os.Getenv("BA_TF_ACC_VAR_region_project_id")
		provider  = os.Getenv("BA_TF_ACC_VAR_region_provider")
		regionID  = os.Getenv("BA_TF_ACC_VAR_region_region_id")

		regionConfig = `resource "biganimal_region" "this" {
  status 		 = "%s"
  project_id     = "%s"
  cloud_provider = "%s"
  region_id      = "%s"
}`
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "region", acc_env_vars_checklist)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
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
