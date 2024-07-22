terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.11.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

variable "cluster_name" {
  type        = string
  description = "The name of the cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_cluster" "ha_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  pause        = false

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

  maintenance_window = {
    is_enabled = true
    start_day  = 6
    start_time = "03:00"
  }

  pg_type               = "epas"
  pg_version            = "15"
  private_networking    = false
  cloud_provider        = "aws"
  read_only_connections = true
  region                = "us-east-1"
  superuser_access      = true
  pgvector              = false
  post_gis              = false
  volume_snapshot_backup = false

  pg_bouncer = {
    is_enabled = false
    #  settings = [ # If is_enabled is true, remove the comment and enter the settings. Should you prefer something different from the defaults.
    #    {
    #      name      = "autodb_idle_timeout"
    #      operation = "read-write" #valid values ["read-write", "read-only"]. "read-only" is only valid for ha clusters with read_only_connections set to true
    #      value     = "5000"
    #    },
    #    {
    #      name      = "client_idle_timeout"
    #      operation = "read-write" #valid values ["read-write", "read-only"]. "read-only" is only valid for ha clusters with read_only_connections set to true
    #      value     = "6000"
    #    },
    #  ]
  }
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.ha_cluster.password
}

output "ro_connection_uri" {
  value = resource.biganimal_cluster.ha_cluster.ro_connection_uri
}

output "faraway_replica_ids" {
  value = resource.biganimal_cluster.ha_cluster.faraway_replica_ids
}
