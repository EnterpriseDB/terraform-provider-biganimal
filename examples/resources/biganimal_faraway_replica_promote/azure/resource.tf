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

variable "cluster_name" {
  type        = string
  description = "The name of the faraway replica cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_faraway_replica_promote" "promote" {
  cluster_name      = var.cluster_name
  project_id        = var.project_id

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
  #  backup_schedule_time = "0 5 1 * * *" //24 hour format cron expression e.g. "0 5 1 * * *" is 01:05
  csp_auth      = false
  instance_type = "azure:Standard_D2s_v3"

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
    size              = "4 Gi" # for azurepremiumstorage please check Premium storage disk sizes here: https://learn.microsoft.com/en-us/azure/virtual-machines/premium-storage-performance
  }
  #  wal_storage = {
  #    volume_type       = "azurepremiumstorage"
  #    volume_properties = "P1"
  #    size              = "4 Gi" # for azurepremiumstorage please check Premium storage disk sizes here: https://learn.microsoft.com/en-us/azure/virtual-machines/premium-storage-performance
  #  }
  private_networking = false // field allowed_ip_ranges will need to be set as "allowed_ip_ranges = []" if private_networking = true
  region             = "australiaeast"

  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]

  # pe_allowed_principal_ids = [
  #   <example_value> # ex: "9334e5e6-7f47-aE61-5A4F-ee067daeEf4A"
  # ]

  # transparent_data_encryption = {
  #   key_id = <example_value>
  # }

  volume_snapshot_backup = false
  password               = resource.random_password.password.result

  cluster_architecture = {
      id    = "single"
      nodes = 1
  }
}

moved {
  from = biganimal_faraway_replica.faraway_replica
  to   = biganimal_faraway_replica_promote.promote
}
