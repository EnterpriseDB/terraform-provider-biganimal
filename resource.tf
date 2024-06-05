terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.8.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "biganimal_faraway_replica" "faraway_replica" {
  cluster_name      = "wai-far"
  project_id        = "prj_QjA3eosXYUuMrba1"
  source_cluster_id = "p-pp7rqagzhr"

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
  instance_type           = "aws:c6i.large"

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
    iops              = "3000"
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "32 Gi"
  }
  private_networking = false
  region             = "us-east-1"
}
