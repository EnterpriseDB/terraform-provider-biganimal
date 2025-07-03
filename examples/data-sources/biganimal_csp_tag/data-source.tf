terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "3.0.0"
    }
  }
}

data "biganimal_csp_tag" "this" {
  project_id        = "<ex-project-id>"        # ex: "prj_12345"
  cloud_provider_id = "<ex-cloud-provider-id>" # ex: "aws"
}

output "csp_tags" {
  value = data.biganimal_csp_tag.this.csp_tags
}
