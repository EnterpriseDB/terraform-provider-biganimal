package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func TestAccBiganimalAWSConnectionResource_basic(t *testing.T) {
	t.Skip()
	var (
		acc_env_vars_checklist = []string{
			"BA_TF_ACC_VAR_aws_connection_project_id",
			"BA_TF_ACC_VAR_aws_connection_role_arn",
			"BA_TF_ACC_VAR_aws_connection_region_id",
		}

		projectID  = os.Getenv("BA_TF_ACC_VAR_aws_connection_project_id")
		roleARN    = os.Getenv("BA_TF_ACC_VAR_aws_connection_role_arn")
		externalID = os.Getenv("BA_TF_ACC_VAR_aws_connection_external_id")
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccResourcePreCheck(t, "aws_connection", acc_env_vars_checklist)
		},
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: awsConnConfig(projectID, roleARN, externalID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_aws_connection.test_aws_conn", "role_arn", roleARN),
					resource.TestCheckResourceAttr("biganimal_aws_connection.test_aws_conn", "external_id", externalID),
				),
			},
		},
	})
}

func awsConnConfig(projectID, roleARN, externalID string) string { // nolint:staticcheck
	return fmt.Sprintf(`resource "biganimal_aws_connection" "test_aws_conn" {
		project_id  = "%s"
		role_arn    = "%s"
		external_id = "%s"
	}`, projectID, roleARN, externalID)
}
