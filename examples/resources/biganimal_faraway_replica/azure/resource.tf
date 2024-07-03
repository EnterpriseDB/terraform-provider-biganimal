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

variable "cluster_name" {
  type        = string
  description = "The name of the faraway replica cluster."
}

variable "source_cluster_id" {
  type        = string
  description = "BigAnimal source cluster ID"
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_faraway_replica" "faraway_replica" {
  cluster_name      = var.cluster_name
  project_id        = var.project_id
  source_cluster_id = var.source_cluster_id

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

  backup_retention_period = "8d"
  csp_auth                = false
  instance_type           = "azure:Standard_D2s_v3"

  // only following pg_config parameters are configurable for faraway replica
  // max_connections, max_locks_per_transaction, max_prepared_transactions, max_wal_senders, max_worker_processes.
  // it is highly recommended setting these values to be equal to or greater than the source cluster's.
  // Please visit [this page](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/managing_replicas/#modify-a-faraway-replica)for best practices.
  pg_config = [
    {
      name  = "max_connections"
      value = "100"
    },
    {
      name  = "max_locks_per_transaction"
      value = "64"
    }
  ]

  storage = {
    volume_type       = "azurepremiumstorage"
    volume_properties = "P1"
    size              = "4 Gi"
  }

  private_networking = false
  region             = "australiaeast"
  # pe_allowed_principal_ids = [
  #   <example_value> # ex: "9334e5e6-7f47-aE61-5A4F-ee067daeEf4A"
  # ]
}
