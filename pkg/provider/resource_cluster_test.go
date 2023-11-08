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
			"BA_TF_ACC_VAR_cluster_provider",
		}
		accClusterName = fmt.Sprintf("acctest_cluster_basic_%v", time.Now().Unix())
		accProjectID   = os.Getenv("BA_TF_ACC_VAR_cluster_project_id")
		accRegion      = os.Getenv("BA_TF_ACC_VAR_cluster_region")
		accProvider    = os.Getenv("BA_TF_ACC_VAR_cluster_provider")
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "cluster", acc_env_vars_checklist)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: clusterResourceConfig(accClusterName, accProjectID, accProvider, accRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "instance_type", "aws:m5.large"),
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "storage.volume_type", "gp3"),
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "storage.size", "4 Gi"),
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "pg_type", "epas"),
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "pg_version", "15"),
				),
			},
			{
				Config: clusterResourceConfigForUpdate(accClusterName, accProjectID, accProvider, accRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "storage.size", "6 Gi"),
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "backup_retention_period", "10d"),
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "cluster_architecture.id", "ha"),
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "cluster_architecture.nodes", "2"),
					resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "cluster_architecture.name", "Primary/Standby High Availability"),
				),
			},
		},
	})
}

func clusterResourceConfig(cluster_name, projectID, provider, region string) string {
	return fmt.Sprintf(`resource "biganimal_cluster" "acctest_cluster" {
  cluster_name = "%s"
  project_id   = "%s"
  cloud_provider        = "%s"
  region                = "%s"
  password      = "thisismyverystrongpassword"

  backup_retention_period = "6d"
  cluster_architecture {
    id    = "single"
    nodes = 1
  }
  csp_auth = false

  instance_type = "aws:m5.large"

  storage {
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
	iops              = "3000"
  }

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  pg_config {
    name  = "application_name"
    value = "created through terraform"
  }

  pg_config {
    name  = "array_nulls"
    value = "off"
  }
  pg_type               = "epas"
  pg_version            = "15"
  private_networking    = false
  read_only_connections = false
  superuser_access      = true
}`, cluster_name, projectID, provider, region)
}

func clusterResourceConfigForUpdate(cluster_name, projectID, provider, region string) string {
	return fmt.Sprintf(`resource "biganimal_cluster" "acctest_cluster" {
  cluster_name = "%s"
  project_id   = "%s"
  cloud_provider        = "%s"
  region                = "%s"
  password      = "thisismyverystrongpassword"

  backup_retention_period = "10d"
  cluster_architecture {
    id    = "ha"
    nodes = 2
    name = "Primary/Standby High Availability"
  }
  csp_auth = true

  instance_type = "aws:m5.large"

  storage {
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "6 Gi"
    iops              = "3000"
  }

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  pg_config {
    name  = "application_name"
    value = "created through terraform"
  }

  pg_config {
    name  = "array_nulls"
    value = "off"
  }
  pg_type               = "epas"
  pg_version            = "15"
  private_networking    = false
  read_only_connections = false
  superuser_access      = true
}`, cluster_name, projectID, provider, region)
}
