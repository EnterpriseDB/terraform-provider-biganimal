package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const cluster_name = "TF-ACCTestCluster"

func TestAccResourceCluster(t *testing.T) {
	var (
		projectID = os.Getenv("TF_VAR_project_id")
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: clusterResourceConfig(projectID),
				Check:  resource.TestCheckResourceAttr("biganimal_cluster.acctest_cluster", "instance_type", "aws:m5.large"),
				// Otherwise, it gives the following error:
				// resource_cluster_test.go:16: Step 1/1 error: After applying this test step, the plan was not empty.
				// Remove this ExpectNonEmptyPlan: true, when the following PR is merged:
				// https://github.com/EnterpriseDB/terraform-provider-biganimal/pull/128
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func clusterResourceConfig(projectID string) string {
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
  region                = "us-east-1"
}`, cluster_name, projectID)
}
