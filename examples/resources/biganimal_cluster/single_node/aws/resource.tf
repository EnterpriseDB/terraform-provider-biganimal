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

resource "biganimal_cluster" "single_node_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  pause        = false

  allowed_ip_ranges = [
    {
      cidr_block  = "127.0.0.1/32"
      description = "localhost"
    },
    {
      cidr_block  = "192.168.0.1/32"
      description = "description!"
    }
  ]

  backup_retention_period = "6d"
  cluster_architecture = {
    id    = "single"
    nodes = 1
  }
  csp_auth = true

  instance_type = "aws:m5.large"
  password      = resource.random_password.password.result
  pg_config = [
    {
      name  = "application_name"
      value = "created through terraform"
    },
    {
      name  = "array_nulls"
      value = "off"
    }
  ]

  storage = {
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
  cloud_provider        = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
  read_only_connections = false
  region                = "us-east-1"
  superuser_access      = true
  pgvector              = false
  post_gis              = false

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

  # pe_allowed_principal_ids = [
  #   <example_value> # ex: 123456789012
  # ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.single_node_cluster.password
}

output "faraway_replica_ids" {
  value = biganimal_cluster.single_node_cluster.faraway_replica_ids
}
