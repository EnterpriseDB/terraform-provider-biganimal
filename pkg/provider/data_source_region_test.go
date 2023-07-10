package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceRegion_basic(t *testing.T) {
	var (
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_region_project_id",
			"BA_TF_ACC_VAR_region_provider",
			"BA_TF_ACC_VAR_region_region_id",
		}
		projectId = os.Getenv("BA_TF_ACC_VAR_region_project_id")
		provider  = os.Getenv("BA_TF_ACC_VAR_region_provider")
		regionId  = os.Getenv("BA_TF_ACC_VAR_region_region_id")
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "datasource region", acc_env_vars_checklist)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: regionDataSourceConfig(projectId, provider, regionId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.biganimal_region.test", "region.0.name", "biganimal_region.test", "name"),
					resource.TestCheckResourceAttrPair("data.biganimal_region.test", "region.0.status", "biganimal_region.test", "status"),
				),
			},
		},
	})
}

func regionDataSourceConfig(projectId, provider, regionId string) string {
	return fmt.Sprintf(`data "biganimal_region" "test" {
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
