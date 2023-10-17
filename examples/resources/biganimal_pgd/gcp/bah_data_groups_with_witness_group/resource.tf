terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.6.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.5.1"
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
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "gcp:e2-highcpu-4"
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
        volume_type       = "pd-ssd"
        volume_properties = "pd-ssd"
        size              = "10 Gi"
      }
      pg_type = {
        pg_type_id = "epas"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:gcp"
      }
      region = {
        region_id = "us-east1"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 6
        start_time = "13:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: "development-data-123456"
      # ]

      # service_account_ids = [
      #   <only_needed_for_bah:gcp_clusters> # ex: "test@development-data-123456.iam.gserviceaccount.com"
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
      cluster_architecture = {
        cluster_architecture_id = "pgd"
        nodes                   = 3
      }
      csp_auth = false
      instance_type = {
        instance_type_id = "gcp:e2-highcpu-4"
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
        volume_type       = "pd-ssd"
        volume_properties = "pd-ssd"
        size              = "10 Gi"
      }
      pg_type = {
        pg_type_id = "epas"
      }
      pg_version = {
        pg_version_id = "15"
      }
      private_networking = false
      cloud_provider = {
        cloud_provider_id = "bah:gcp"
      }
      region = {
        region_id = "europe-west1"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 5
        start_time = "12:00"
      }
      # pe_allowed_principal_ids = [
      #   <example_value> # ex: "development-data-123456"
      # ]

      # service_account_ids = [
      #   <only_needed_for_bah:gcp_clusters> # ex: "test@development-data-123456.iam.gserviceaccount.com"
      # ]
    }
  ]
  witness_groups = [
    {
      region = {
        region_id = "asia-south1"
      }
      cloud_provider = {
        cloud_provider_id = "bah:gcp"
      }
      maintenance_window = {
        is_enabled = true
        start_day  = 3
        start_time = "03:00"
      }
    }
  ]
}
