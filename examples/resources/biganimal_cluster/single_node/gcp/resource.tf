terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "2.0.0"
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
  #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
  cluster_architecture = {
    id    = "single"
    nodes = 1
  }
  csp_auth = false

  instance_type = "gcp:e2-highcpu-4"
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
    volume_type       = "pd-ssd"
    volume_properties = "pd-ssd"
    size              = "10 Gi"
  }

  #  wal_storage = {
  #    volume_type       = "pd-ssd"
  #    volume_properties = "pd-ssd"
  #    size              = "10 Gi"
  #  }

  maintenance_window = {
    is_enabled = true
    start_day  = 6
    start_time = "03:00"
  }

  pg_type                = "epas" #valid values ["epas", "pgextended", "postgres]"
  pg_version             = "15"
  private_networking     = false
  cloud_provider         = "bah:gcp" // "bah:gpc" uses BigAnimal's cloud account Google Cloud provider, use "gcp" for your cloud account
  read_only_connections  = false
  region                 = "us-east1"
  superuser_access       = false
  pgvector               = false
  post_gis               = false
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

  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]

  # pe_allowed_principal_ids = [
  #   <example_value> # ex: "development-data-123456"
  # ]

  # service_account_ids = [
  #   <only_needed_for_bah:gcp_clusters> # ex: "test@development-data-123456.iam.gserviceaccount.com"
  # ]

  # transparent_data_encryption = {
  #   key_id = <example_value>
  # }
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.single_node_cluster.password
}

output "faraway_replica_ids" {
  value = biganimal_cluster.single_node_cluster.faraway_replica_ids
}
