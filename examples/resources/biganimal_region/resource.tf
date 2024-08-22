terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.0.0"
    }
  }
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_region" "this" {
  cloud_provider = "aws"
  region_id      = "eu-west-1"
  project_id     = "prj_abc01234567"
  
  tags = [
    {
      tag_name  = "test"
      color = "blue"
    },
    {
      tag_name  = "<ex-tag-name-2>"
    },
  ]
}

output "region_status" {
  value = resource.biganimal_region.this.status
}

output "region_name" {
  value = resource.biganimal_region.this.name
}

output "region_continent" {
  value = resource.biganimal_region.this.continent
}
