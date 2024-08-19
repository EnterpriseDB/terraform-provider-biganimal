terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.0.0"
    }
  }
}

data "biganimal_csp_tag" "this" {
  project_id = "ex-project-id"
  cloud_provider_id = "ex-cloud-provider-id"
}

output "csp_tags" {
  value = data.biganimal_csp_tag.this.csp_tags
}
