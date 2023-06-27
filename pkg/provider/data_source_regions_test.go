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
			"BA_TF_ACC_VAR_datasource_regions_project_id",
			"BA_TF_ACC_VAR_datasource_regions_provider",
			"BA_TF_ACC_VAR_datasource_regions_region_id",
			"BA_TF_ACC_VAR_datasource_regions_status",
		}
		projectId  = os.Getenv("BA_TF_ACC_VAR_datasource_regions_project_id")
		provider   = os.Getenv("BA_TF_ACC_VAR_datasource_regions_provider")
		regionId   = os.Getenv("BA_TF_ACC_VAR_datasource_regions_region_id")
		regionName = os.Getenv("BA_TF_ACC_VAR_datasource_regions_region_name")
		status     = os.Getenv("BA_TF_ACC_VAR_datasource_regions_status")
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

					resource.TestCheckResourceAttr("data.biganimal_regions.test", "regions.0.name", regionName),
					resource.TestCheckResourceAttr("data.biganimal_regions.test", "regions.0.status", status),
				),
			},
		},
	})
}

func regionsDataSourceConfig(projectId, provider, regionId string) string {
	return fmt.Sprintf(`data "biganimal_regions" "test" {
  project_id     = "%s"
  cloud_provider = "%s"
  region_id      = "%s"
}
`, projectId, provider, regionId)
}
