terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.0.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "biganimal_csp_tag" "this" {
  project_id        = <example_project_id>
  cloud_provider_id = <example_cloud_provider_id> #ex cloud-provider-id values ["bah:aws", "bah:azure", "bah:gcp", "aws", "azure", "gcp"]
  
  add_tags = [
    #{
    #  csp_tag_key   = <example_csp_tag_key>
    #  csp_tag_value = <example_csp_tag_value>
    #},
    #{
    #  csp_tag_key   = <example_csp_tag_key>
    #  csp_tag_value = <example_csp_tag_value>
    #},
  ]
  
  delete_tags = [
    #<example_csp_tag_id>,
    #<example_csp_tag_id>,
  ]
  
  edit_tags = [
    #{
    #  csp_tag_id    = <example_csp_tag_id>
    #  csp_tag_key   = <example_csp_tag_key>
    #  csp_tag_value = <example_csp_tag_value>
    #  status        = "OK"
    #},
    #{
    #  csp_tag_id    = <example_csp_tag_id>
    #  csp_tag_key   = <example_csp_tag_key>
    #  csp_tag_value = <example_csp_tag_value>
    #  status        = "OK"
    #},
  ]
}
