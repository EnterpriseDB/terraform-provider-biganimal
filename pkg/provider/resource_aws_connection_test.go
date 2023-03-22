package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBiganimalAWSConnectionResource(t *testing.T) {
	var (
		projectID  = os.Getenv("BA_TF_PROJECT")
		roleARN    = os.Getenv("BA_TF_AWS_ROLE_ARN")
		externalID = os.Getenv("BA_TF_AWS_EXTERNAL_ID")
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
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
func awsConnConfig(projectID, roleARN, externalID string) string {
	return fmt.Sprintf(`resource "biganimal_aws_connection" "test_aws_conn" {
		project_id  = "%s"
		role_arn    = "%s"
		external_id = "%s"
	}`, projectID, roleARN, externalID)
}
