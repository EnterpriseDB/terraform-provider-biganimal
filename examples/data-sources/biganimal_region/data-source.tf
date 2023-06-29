variable "cloud_provider" {
  type        = string
  description = "Cloud Provider"

  validation {
    condition     = contains(["aws", "azure", "bah:aws"], var.cloud_provider)
    error_message = "Please select one of the supported regions: aws, azure or bah:aws."
  }
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"

}

data "biganimal_region" "this" {
  cloud_provider = var.cloud_provider
  project_id     = var.project_id
  // region_id   = "us-west-1" //optional
  // query       = "eu" // optional
}

output "regions" {
  value = data.biganimal_region.this.regions
}

output "cloud_provider_id" {
  value = data.biganimal_region.this.cloud_provider
}
