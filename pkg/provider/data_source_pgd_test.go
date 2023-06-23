package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourcePGD_basic(t *testing.T) {
	var (
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_pgd_project_id",
			"BA_TF_ACC_VAR_pgd_name",
		}
		accPGDName   = os.Getenv("BA_TF_ACC_VAR_pgd_name")
		accProjectID = os.Getenv("BA_TF_ACC_VAR_pgd_project_id")
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "pgd", acc_env_vars_checklist)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: pgdDataSourceConfig(accPGDName, accProjectID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.biganimal_pgd.acctest_pgd", "cluster_name", accPGDName),
					resource.TestCheckResourceAttr("data.biganimal_pgd.acctest_pgd", "project_id", accProjectID),
				),
			},
		},
	})
}

func pgdDataSourceConfig(cluster_name, projectID string) string {
	return fmt.Sprintf(`data "biganimal_pgd" "acctest_pgd" {
	cluster_name = "%s"
	project_id   = "%s"
  }`, cluster_name, projectID)
}
