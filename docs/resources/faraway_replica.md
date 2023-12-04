# biganimal_faraway_replica (Resource)

The faraway replica resource is used to manage cluster faraway-replicas on different active regions in the cloud. See [Managing replicas](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/managing_replicas/) for more details.

## Example Usage

```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.6.1"
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

  allowed_ip_ranges {
    cidr_block  = "127.0.0.1/32"
    description = "localhost"
  }

  allowed_ip_ranges {
    cidr_block  = "192.168.0.1/32"
    description = "description!"
  }

  backup_retention_period = "6d"
  csp_auth                = true
  instance_type           = "aws:m5.large"

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
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }
  private_networking = false
  region             = "eu-west-2"
}
```

## Example of Creating a Single Node Cluster and Its Faraway Replica

BigAnimal does not currently support provisioning a replica on a different cloud provider from the source cluster.
That's why both the cluster and the faraway replica in this example are running on the same cloud provider.

```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.6.1"
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_name` (String) Name of the faraway replica cluster.
- `instance_type` (String) Instance type. For example, "azure:Standard_D2s_v3", "aws:c5.large" or "gcp:e2-highcpu-4".
- `project_id` (String) BigAnimal Project ID.
- `region` (String) Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.
- `source_cluster_id` (String) Source cluster ID.
- `storage` (Block List, Min: 1) Storage. (see [below for nested schema](#nestedblock--storage))

### Optional

- `allowed_ip_ranges` (Block Set) Allowed IP ranges. (see [below for nested schema](#nestedblock--allowed_ip_ranges))
- `backup_retention_period` (String) Backup retention period. For example, "7d", "2w", or "3m".
- `csp_auth` (Boolean) Is authentication handled by the cloud service provider. Available for AWS only, See [Authentication](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/#authentication) for details.
- `pg_config` (Block Set) Database configuration parameters. See [Modifying database configuration parameters](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/03_modifying_your_cluster/05_db_configuration_parameters/) for details. (see [below for nested schema](#nestedblock--pg_config))
- `private_networking` (Boolean) Is private networking enabled.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `cluster_id` (String) Cluster ID.
- `cluster_type` (String) Type of the cluster. For example, "cluster" for biganimal_cluster resources, or "faraway_replica" for biganimal_faraway_replica resources.
- `connection_uri` (String) Cluster connection URI.
- `created_at` (String) Cluster creation time.
- `deleted_at` (String) Cluster deletion time.
- `expired_at` (String) Cluster expiry time.
- `id` (String) The ID of this resource.
- `logs_url` (String) The URL to find the logs of this cluster.
- `metrics_url` (String) The URL to find the metrics of this cluster.
- `phase` (String) Current phase of the cluster.
- `resizing_pvc` (List of String) Resizing PVC.

<a id="nestedblock--storage"></a>
### Nested Schema for `storage`

Required:

- `volume_properties` (String) Volume properties in accordance with the selected volume type.
- `volume_type` (String) Volume type. For Azure: "azurepremiumstorage" or "ultradisk". For AWS: "gp3", "io2", or "io2-block-express". For Google Cloud: only "pd-ssd".

Optional:

- `iops` (String) IOPS for the selected volume. It can be set to different values depending on your volume type and properties.
- `size` (String) Size of the volume. It can be set to different values depending on your volume type and properties.
- `throughput` (String) Throughput is automatically calculated by BigAnimal based on the IOPS input.


<a id="nestedblock--allowed_ip_ranges"></a>
### Nested Schema for `allowed_ip_ranges`

Required:

- `cidr_block` (String) CIDR block.

Optional:

- `description` (String) CIDR block description.


<a id="nestedblock--pg_config"></a>
### Nested Schema for `pg_config`

Required:

- `name` (String) GUC name.
- `value` (String) GUC value.


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
