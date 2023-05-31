package provider

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceCluster_basic(t *testing.T) {
	var (
		// Add env vars to fetch to the following checklist
		// The variable naming scheme is as follows:
		// BA_TF_ACC_VAR_<resource_type>_<variable_name>
		// e.g. for biganimal_cluster resource, the project_id variable can be fetched
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_cluster_project_id",
			"BA_TF_ACC_VAR_cluster_region",
		}
		accClusterName = fmt.Sprintf("acctest_cluster_basic_%v", time.Now().Unix())
		accProjectID   = os.Getenv("BA_TF_ACC_VAR_cluster_project_id")
		accRegion      = os.Getenv("BA_TF_ACC_VAR_cluster_region")
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "cluster", acc_env_vars_checklist)
		},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: clusterResourceConfig(accClusterName, accProjectID, accRegion),
				Check:  resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "instance_type", "aws:m5.large"),
			},
		},
	})
}

func clusterResourceConfig(cluster_name, projectID, region string) string {
	return fmt.Sprintf(`resource "biganimal_cluster" "acctest_cluster" {
  cluster_name = "%s"
  project_id   = "%s"

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
  cluster_architecture {
    id    = "single"
    nodes = 1
  }
  csp_auth = true

  instance_type = "aws:m5.large"
  password      = "thisismyverystrongpassword"
  pg_config {
    name  = "application_name"
    value = "created through terraform"
  }

  pg_config {
    name  = "array_nulls"
    value = "off"
  }

  storage {
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }

  pg_type               = "epas"
  pg_version            = "14"
  private_networking    = false
  cloud_provider        = "aws"
  read_only_connections = false
  region                = "%s"
}`, cluster_name, projectID, region)
}
