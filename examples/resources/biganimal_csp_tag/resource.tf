terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.1.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "biganimal_csp_tag" "this" {
  project_id        = "<ex-project-id>"        # ex: "prj_12345"
  cloud_provider_id = "<ex-cloud-provider-id>" # ex cloud-provider-id values ["bah:aws", "bah:azure", "bah:gcp", "aws", "azure", "gcp"]

  add_tags = [
    #{
    #  csp_tag_key   = "<ex-csp-tag-key>"   # ex: "key"
    #  csp_tag_value = "<ex-csp-tag-value>" # ex: "value"
    #},
    #{
    #  csp_tag_key   = "<ex-csp-tag-key>"
    #  csp_tag_value = "<ex-csp-tag-value>"
    #},
  ]

  delete_tags = [
    #"<ex-csp-tag-id>", # ex: "id"
    #"<ex-csp-tag-id>",
  ]

  edit_tags = [
    #{
    #  csp_tag_id    = "<ex-csp-tag-id>"    # ex: "id"
    #  csp_tag_key   = "<ex-csp-tag-key>"   # ex: "key"
    #  csp_tag_value = "<ex-csp-tag-value>" # ex: "value"
    #  status        = "OK"
    #},
    #{
    #  csp_tag_id    = "<ex-csp-tag-id>"
    #  csp_tag_key   = "<ex-csp-tag-key>"
    #  csp_tag_value = "<ex-csp-tag-value>"
    #  status        = "OK"
    #},
  ]
}
