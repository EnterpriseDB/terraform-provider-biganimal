terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.5.0"
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

resource "biganimal_cluster" "single_node_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id

  cluster_architecture {
    id    = "single"
    nodes = 1
  }

  instance_type = "azure:Standard_D2s_v3"
  password      = resource.random_password.password.result

  storage {
    volume_type       = "azurepremiumstorage"
    volume_properties = "P1"
    size              = "4 Gi"
  }

  pg_type        = "epas"
  pg_version     = "15"
  cloud_provider = "azure"
  region         = "eastus"
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.single_node_cluster.password
}

resource "biganimal_faraway_replica" "faraway_replica" {
  cluster_name      = "${var.cluster_name}-FAR"
  project_id        = var.project_id
  source_cluster_id = resource.biganimal_cluster.single_node_cluster.cluster_id

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
  csp_auth                = false
  instance_type           = "azure:Standard_D2s_v3"


  // only following pg_config parameters are configurable for faraway replica
  // max_connections, max_locks_per_transaction, max_prepared_transactions, max_wal_senders, max_worker_processes.
  // it is highly recommended setting these values to be equal to or greater than the source cluster's.
  // Please visit [this page](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/managing_replicas/#modify-a-faraway-replica)for best practices.
  pg_config {
    name  = "max_connections"
    value = "100"
  }

  pg_config {
    name  = "max_locks_per_transaction"
    value = "64"
  }
  pg_config {
    name  = "max_prepared_transactions"
    value = "0"
  }
  pg_config {
    name  = "max_wal_senders"
    value = "10"
  }
  pg_config {
    name  = "max_worker_processes"
    value = "32"
  }

  storage {
    volume_type       = "azurepremiumstorage"
    volume_properties = "P1"
    size              = "4 Gi"
  }

  private_networking = false
  region             = "centralindia"
}
