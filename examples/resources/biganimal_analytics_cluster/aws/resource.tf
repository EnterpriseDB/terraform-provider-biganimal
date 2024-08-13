terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.0.0"
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

resource "biganimal_analytics_cluster" "analytics_cluster" {
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
    },
  ]

  backup_retention_period = "30d"
  csp_auth                = false

  instance_type = "aws:m6id.12xlarge"
  password      = resource.random_password.password.result

  maintenance_window = {
    is_enabled = false
    start_day  = 0
    start_time = "00:00"
  }

  pg_type            = "epas" #valid values ["epas", "pgextended", "postgres]"
  pg_version         = "16"
  private_networking = false
  cloud_provider     = "bah:aws"
  region             = "ap-south-1"
  # pe_allowed_principal_ids = [
  #   <example_value> # ex: 123456789012
  # ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_analytics_cluster.analytics_cluster.password
}
