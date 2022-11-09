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
