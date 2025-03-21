terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "2.0.0"
    }
  }
}


variable "external_id" {
  type = string
}
variable "project_id" {
  type = string
}
variable "role_arn" {
  type = string
}


resource "biganimal_aws_connection" "project_aws_conn" {
  project_id  = var.project_id
  role_arn    = var.role_arn
  external_id = var.external_id
}
