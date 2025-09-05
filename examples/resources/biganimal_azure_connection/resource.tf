terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "3.1.0"
    }
  }
}


variable "subscription_id" {
  type = string
}

variable "project_id" {
  type = string
}

variable "client_id" {
  type = string
}

variable "client_secret" {
  type = string
}

variable "tenant_id" {
  type = string
}


resource "biganimal_azure_connection" "project_azure_conn" {
  project_id      = var.project_id
  tenant_id       = var.tenant_id
  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
}
