terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.1.1"
    }
  }
}

data "biganimal_tag" "this" {
  tag_id = "tag-id"
}

output "tag_name" {
  value = data.biganimal_tag.this.tag_name
}

output "color" {
  value = data.biganimal_tag.this.color
}
