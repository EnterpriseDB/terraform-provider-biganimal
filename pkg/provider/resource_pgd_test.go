package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourcePGD_basic(t *testing.T) {
	var (
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_pgd_project_id",
			"BA_TF_ACC_VAR_pgd_name",
		}
		accPGDName   = os.Getenv("BA_TF_ACC_VAR_pgd_name")
		accProjectID = os.Getenv("BA_TF_ACC_VAR_pgd_project_id")
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "pgd_resource", acc_env_vars_checklist)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: pgdResourceConfig(accPGDName, accProjectID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_pgd.acctest_pgd", "project_id", accProjectID),
					resource.TestCheckResourceAttr("biganimal_pgd.acctest_pgd", "cluster_name", accPGDName),
					resource.TestCheckResourceAttrSet("biganimal_pgd.acctest_pgd", "cluster_id"),
				),
			},
		},
	})
}

func pgdResourceConfig(cluster_name, projectID string) string {
	return fmt.Sprintf(`resource "biganimal_pgd" "acctest_pgd" {
	cluster_name = "%s"
	project_id   = "%s"
	password     = "thisismyverystrongpassword"
	data_groups = [
    {
      allowed_ip_ranges = [
        {
          cidr_block  = "127.0.0.1/32"
          description = "localhost"
        },
        {
          cidr_block  = "192.168.0.1/32"
          description = "description!"
        },
      ]
      backup_retention_period = "6d"
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 2
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "azure:Standard_D2s_v3"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "azurepremiumstorage"
        volume_properties = "P1"
        size              = "4 Gi"
      }
      pg_type = {
        pg_type_id = "epas"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "azure"
      }
      region = {
        region_id = "northeurope"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 1
        start_time = "13:00"
      }
    },
  ]
  }`, cluster_name, projectID)
}
