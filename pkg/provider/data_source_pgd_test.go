package provider

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourcePGD_basic(t *testing.T) {
	var (
		// Add env vars to fetch to the following checklist
		// The variable naming scheme is as follows:
		// BA_TF_ACC_VAR_<resource_type>_<variable_name>
		// e.g. for biganimal_cluster resource, the project_id variable can be fetched
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_pgd_project_id",
		}
		accPGDName   = fmt.Sprintf("acctest_pgd_basic_%v", time.Now().Unix())
		accProjectID = os.Getenv("BA_TF_ACC_VAR_pgd_project_id")
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "pgd", acc_env_vars_checklist)
		},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: pgdDataSourceConfig(accPGDName, accProjectID),
				Check:  resource.TestCheckResourceAttr("biganimal_pgd.acctest_pgd", "cluster_name", accPGDName),
			},
		},
	})
}

func pgdDataSourceConfig(cluster_name, projectID string) string {
	return fmt.Sprintf(`resource "biganimal_pgd" "acctest_pgd" {
    name = "%s"
    project_id = "%s"
  }

  data "biganimal_pgd" "acctest_pgd" {
	cluster_name = biganimal_pgd.acctest_pgd.name
	project_id   = biganimal_pgd.acctest_pgd.project_id
  }`, cluster_name, projectID)
}
