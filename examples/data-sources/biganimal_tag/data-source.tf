terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "3.0.1"
    }
  }
}

output "tag_name" {
  value = data.biganimal_tag.this.tag_name
}

output "color" {
  value = data.biganimal_tag.this.color
}
