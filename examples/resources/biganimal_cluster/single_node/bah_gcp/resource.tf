terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.6.0"
    }
  }
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

  allowed_ip_ranges {
    cidr_block  = "0.0.0.0/0"
    description = "To allow all access"
  }

  backup_retention_period = "6d"
  cluster_architecture {
    id    = "single"
    nodes = 1
  }
  csp_auth = false

  instance_type = "gcp:e2-highcpu-4"
  password      = "thisismyverystrongpassword"
  pg_config {
    name  = "application_name"
    value = "created through terraform"
  }

  pg_config {
    name  = "array_nulls"
    value = "off"
  }

  storage {
    volume_type       = "pd-ssd"
    volume_properties = "pd-ssd"
    size              = "10 Gi"
  }

  maintenance_window = {
    is_enabled = false
    start_day  = 0
    start_time = "00:00"
  }

  # pe_allowed_principal_ids = [
  #   <example_value>
  # ]

  # service_account_ids = [
  #   <only_needed_for_bah:gcp_clusters>
  # ]

  pg_type               = "epas"
  pg_version            = "15"
  private_networking    = true
  cloud_provider        = "bah:gcp"
  read_only_connections = false
  region                = "europe-west1"
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.single_node_cluster.password
}

output "faraway_replica_ids" {
  value = biganimal_cluster.single_node_cluster.faraway_replica_ids
}
