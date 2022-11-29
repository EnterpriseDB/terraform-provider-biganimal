terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.1.0"
    }
  }
}

resource "biganimal_region" "this" {
  cloud_provider = "aws"
  region_id      = "eu-west-1"
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
