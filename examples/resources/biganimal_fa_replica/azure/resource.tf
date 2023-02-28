terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.3.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.3.1"
    }
  }
}

variable "cluster_name" {
  type        = string
  description = "The name of the faraway replica cluster."
}

variable "source_cluster_id" {
  type        = string
  description = "BigAnimal source cluster ID"
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_fareplica" "faraway_replica" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  source_cluster_id = var.source_cluster_id

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
  csp_auth = false
  instance_type = "azure:Standard_D2s_v3"

  // todo: add some other example for pgConfig, below values are not modifyable for farep
//  pg_config {
//    name  = "application_name"
//    value = "created through terraform"
//  }
//
//  pg_config {
//    name  = "array_nulls"
//    value = "off"
//  }

  storage {
    volume_type       = "azurepremiumstorage"
    volume_properties = "P1"
    size              = "4 Gi"
  }

  private_networking    = false
  region                = "australiaeast"
}
