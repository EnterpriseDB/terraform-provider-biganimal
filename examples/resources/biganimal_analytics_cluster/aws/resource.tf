terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "3.0.1"
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

  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]

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
  #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
  csp_auth = false

  instance_type = "aws:m6id.12xlarge"
  password      = resource.random_password.password.result

  maintenance_window = {
    is_enabled = false
    start_day  = 0
    start_time = "00:00"
  }

  pg_type            = "epas" #valid values ["epas", "pgextended", "postgres]"
  pg_version         = "16"
  private_networking = false // field allowed_ip_ranges will need to be set as "allowed_ip_ranges = []" if private_networking = true
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
