package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceRegions_basic(t *testing.T) {
	var (
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_regions_project_id",
			"BA_TF_ACC_VAR_regions_provider",
			"BA_TF_ACC_VAR_regions_region_id",
		}
		projectId = os.Getenv("BA_TF_ACC_VAR_regions_project_id")
		provider  = os.Getenv("BA_TF_ACC_VAR_regions_provider")
		regionId  = os.Getenv("BA_TF_ACC_VAR_regions_region_id")
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "datasource regions", acc_env_vars_checklist)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: regionsDataSourceConfig(projectId, provider, regionId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.biganimal_regions.test", "regions.0.name", "biganimal_regions.test", "name"),
					resource.TestCheckResourceAttrPair("data.biganimal_regions.test", "regions.0.status", "biganimal_regions.test", "status"),
				),
			},
		},
	})
}

func regionsDataSourceConfig(projectId, provider, regionId string) string {
	return fmt.Sprintf(`data "biganimal_regions" "test" {
  project_id     = "%[1]s"
  cloud_provider = "%[2]s"
  region_id      = "%[3]s"
}

resource "biganimal_region" "test" {
  project_id     = "%[1]s"
  cloud_provider = "%[2]s"
  region_id      = "%[3]s"
  # no need to active region actually for saving time'
  status         = "INACTIVE"
}



`, projectId, provider, regionId)
}
