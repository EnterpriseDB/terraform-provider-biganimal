package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBiganimalProjectResource(t *testing.T) {
	projectName := fmt.Sprintf("acc_test_project_%v", time.Now().Unix())

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: projectConfig(projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_project.test_project", "project_name", projectName),
					resource.TestCheckResourceAttrSet("biganimal_project.test_project", "project_id"),
					resource.TestCheckResourceAttrSet("biganimal_project.test_project", "user_count"),
					resource.TestCheckResourceAttrSet("biganimal_project.test_project", "cluster_count"),
					resource.TestCheckResourceAttrSet("biganimal_project.test_project", "cloud_providers"),
				),
			},
			{
				Config: projectConfig(projectName + "_updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("biganimal_project.test_project", "project_name", projectName+"_updated"),
				),
			},
		},
	})
}

func projectConfig(projectName string) string {
	return fmt.Sprintf(`
resource "biganimal_project" "test_project" {
		project_name = "%s"
	}`, projectName)
}
