variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}


data "biganimal_aws_connection" "this" {
  project_id = var.project_id
}

output "external_id" {
  value = data.biganimal_aws_connection.this.external_id
}

output "role_arn" {
  value = data.biganimal_aws_connection.this.role_arn
}
