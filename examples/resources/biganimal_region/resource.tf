terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.1.2"
    }
  }
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"

  validation {
    condition     = can(regex("^prj_[[:alnum:]]{16}$", var.project_id))
    error_message = "Please provide a valid name for the project_id, for example: prj_abcdABCD01234567."
  }
}

resource "biganimal_region" "this" {
  cloud_provider = "aws"
  region_id      = "eu-west-1"
  project_id     = var.project_id
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
