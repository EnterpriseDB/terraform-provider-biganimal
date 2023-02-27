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

//resource "random_password" "password" {
//  length           = 16
//  special          = true
//  override_special = "!#$%&*()-_=+[]{}<>:?"
//}

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
    description = "localhoster"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
//  cluster_architecture {
//    id    = "single"
//    nodes = 1
//  }

  csp_auth = true

  instance_type = "aws:m5.large"
//  password      = resource.random_password.password.result

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
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }

//  pg_type               = "epas"
//  pg_version            = "14"
  private_networking    = false
//  cloud_provider        = "aws"
//  read_only_connections = false
  region                = "eu-west-2"
}

//output "password" {
//  sensitive = true
//  value     = resource.biganimal_cluster.single_node_cluster.password
//}
