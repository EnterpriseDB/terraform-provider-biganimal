terraform {
  required_providers {
    biganimal = {
      source  = "biganimal"
      version = "0.3.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.3.1"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

resource "biganimal_cluster" "this_resource" {
  cluster_name = var.cluster_name

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
  cluster_architecture {
    id    = "ha"
    nodes = 3
  }

  instance_type = "aws:c5.large"
  password      = resource.random_password.password.result
  pg_config {
    name  = "application_name"
    value = "created through terraform"
  }

  pg_config {
    name  = "array_nulls"
    value = "off"
  }

  storage {
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }

  pg_type               = "epas"
  pg_version            = "14"
  private_networking    = false
  cloud_provider        = "aws"
  read_only_connections = true
  region                = "us-east-1"
  replicas              = 1
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.this_resource.password
}

output "ro_connection_uri" {
  value = resource.biganimal_cluster.this_resource.ro_connection_uri
}