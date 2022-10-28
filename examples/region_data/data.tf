terraform {
  required_providers {
    biganimal = {
      source  = "biganimal"
      version = "0.3.1"
    }
  }
}

data "biganimal_region" "this" {
  cloud_provider = var.cloud_provider
  region_id = "us-east-1"
}

output "regions" {
  value = data.biganimal_region.this.regions
}

output "cloud_provider_id" {
  value = data.biganimal_region.this.cloud_provider
}
