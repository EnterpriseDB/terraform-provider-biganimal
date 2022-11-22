variable "cloud_provider" {
  type        = string
  description = "Cloud Provider"

  validation {
    condition     = contains(["aws", "azure"], var.cloud_provider)
    error_message = "Please select one of the supported regions: aws, azure."
  }
}

variable "region_id" {
  type        = string
  description = "region id"
}

data "biganimal_region" "this" {
  cloud_provider = var.cloud_provider
  // region_id   = var.region_id //optional
  // query       = "eu" // optional
}

output "regions" {
  value = data.biganimal_region.this.regions
}

output "cloud_provider_id" {
  value = data.biganimal_region.this.cloud_provider
}
