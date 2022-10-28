terraform {
  required_providers {
    biganimal = {
      source  = "biganimal"
      version = "0.3.1"
    }
  }
}

# resource "biganimal_region" "this" {
#   cloud_provider = "aws"
#   region_id = "us-east-1"
#   status = "SUSPENDED"
# }

# output "region_status" {
#   value = resource.biganimal_region.this.status
# }

# output "region_name" {
#   value = resource.biganimal_region.this.name
# }

# output "region_continent" {
#   value = resource.biganimal_region.this.continent
# }
