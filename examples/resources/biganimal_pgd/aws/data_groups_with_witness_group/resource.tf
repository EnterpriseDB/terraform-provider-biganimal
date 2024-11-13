terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.1.1"
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

resource "biganimal_pgd" "pgd_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
  password     = resource.random_password.password.result
  pause        = false
  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]
  data_groups = [
    {
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
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "aws:m6i.large"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "gp3"
        volume_properties = "gp3"
        size              = "32 Gi"
      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
      }
      region = {
        region_id = "eu-west-1"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 1
        start_time = "13:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: 123456789012
      # ]
    },
    {
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
      backup_retention_period = "6d"
      #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "aws:m6i.large"
      }
      pg_config = [
        {
          name  = "application_name"
          value = "created through terraform"
        },
        {
          name  = "array_nulls"
          value = "off"
        },
      ]
      storage = {
        volume_type       = "gp3"
        volume_properties = "gp3"
        size              = "32 Gi"
      }
      pg_type = {
        pg_type_id = "epas" #valid values ["epas", "pgextended", "postgres]"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
      }
      region = {
        region_id = "eu-west-2"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 2
        start_time = "15:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: 123456789012
      # ]
    }
  ]
  witness_groups = [
    {
      region = {
        region_id = "us-east-1"
      }
      cloud_provider = {
        cloud_provider_id = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 3
        start_time = "03:00"
      }
    }
  ]
}

output "password" {
  sensitive = true
  value     = resource.biganimal_pgd.pgd_cluster.password
}
