# biganimal_faraway_replica (Resource)

The faraway replica resource is used to manage cluster faraway-replicas on different active regions in the cloud. See [Managing replicas](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/managing_replicas/) for more details.



## Example of Creating a Single Node Cluster and Its Faraway Replica

BigAnimal does not currently support provisioning a replica on a different cloud provider from the source cluster.
That's why both the cluster and the faraway replica in this example are running on the same cloud provider.

```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "3.1.0"
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
  csp_auth = false //can't change once set

  instance_type = "azure:Standard_D2s_v3"
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
    volume_type       = "azurepremiumstorage"
    volume_properties = "P1"
    size              = "4 Gi" # for azurepremiumstorage please check Premium storage disk sizes here: https://learn.microsoft.com/en-us/azure/virtual-machines/premium-storage-performance
  }

  #  wal_storage = {
  #    volume_type       = "azurepremiumstorage"
  #    volume_properties = "P1"
  #    size              = "4 Gi" # for azurepremiumstorage please check Premium storage disk sizes here: https://learn.microsoft.com/en-us/azure/virtual-machines/premium-storage-performance
  #  }

  maintenance_window = {
    is_enabled = true
    start_day  = 6
    start_time = "03:00"
  }

  pg_type                = "epas"      #valid values ["epas", "pgextended", "postgres]" //can't change once set
  pg_version             = "15"        //can't change once set
  private_networking     = false       // field allowed_ip_ranges will need to be set as "allowed_ip_ranges = null" if private_networking = true
  cloud_provider         = "bah:azure" // "bah:azure" uses BigAnimal's cloud account Azure, use "azure" for your cloud account
  read_only_connections  = false
  region                 = "eastus2"
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
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.single_node_cluster.password
}

output "faraway_replica_ids" {
  value = biganimal_cluster.single_node_cluster.faraway_replica_ids
}

resource "biganimal_faraway_replica" "faraway_replica" {
  cluster_name      = "${var.cluster_name}-FAR"
  project_id        = var.project_id
  source_cluster_id = resource.biganimal_cluster.single_node_cluster.cluster_id

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
  private_networking = false // field allowed_ip_ranges will need to be set as "allowed_ip_ranges = null" if private_networking = true
  region             = "centralindia"

  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]

  # transparent_data_encryption = {
  #   key_id = <example_value>
  # }

  volume_snapshot_backup = false
}
```

## Example of promoting a Faraway Replica

Please upgrade terraform version to 1.13.1 before using biganimal_faraway_replica_promoted_cluster resource

To promote a Faraway Replica you have to change the resource "biganimal_faraway_replica" to "biganimal_faraway_replica_promoted_cluster"
in your terraform file, use the same fields and the moved command as shown in the example below. You can overwrite your terraform file
with the example file, but make sure the moved commands corresponds to your previous resource and name.

moved example:
moved {
  from = biganimal_faraway_replica.faraway_replica                    // your previous resource and name
  to   = biganimal_faraway_replica_promoted_cluster.promoted_cluster  // migrate to this resource and name
}

```terraform
// To promote biganimal_faraway_replica resource use the biganimal_faraway_replica_promoted_cluster resource. You will have to change your biganimal_faraway_replica resource to biganimal_faraway_replica_promoted_cluster and use the "moved" command as shown in this example.

terraform {
  required_version = "= 1.13.1"
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "3.1.0"
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
  description = "The name of the faraway replica cluster."
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

resource "biganimal_faraway_replica_promoted_cluster" "promoted_cluster" {
  cluster_name = var.cluster_name
  project_id   = var.project_id

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
  instance_type = "aws:c6i.large"

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
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }
  #  wal_storage = {
  #    volume_type       = "gp3"
  #    volume_properties = "gp3"
  #    size              = "4 Gi"
  #    #iops             = "3000" # optional
  #    #throughput       = "125" # optional
  #  }
  private_networking = false // field allowed_ip_ranges will need to be set as "allowed_ip_ranges = null" if private_networking = true
  region             = "ap-south-1"

  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]

  # pe_allowed_principal_ids = [
  #   <example_value> # ex: 123456789012
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
  to   = biganimal_faraway_replica_promoted_cluster.promoted_cluster
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_name` (String) Name of the faraway replica cluster.
- `instance_type` (String) Instance type. For example, "azure:Standard_D2s_v3", "aws:c6i.large" or "gcp:e2-highcpu-4".
- `region` (String) Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.
- `source_cluster_id` (String) Source cluster ID.
- `storage` (Attributes) Storage. (see [below for nested schema](#nestedatt--storage))

### Optional

- `allowed_ip_ranges` (Attributes Set) Allowed IP ranges. (see [below for nested schema](#nestedatt--allowed_ip_ranges))
- `backup_retention_period` (String) Backup retention period. For example, "7d", "2w", or "3m".
- `backup_schedule_time` (String) Backup schedule time in 24 hour cron expression format.
- `csp_auth` (Boolean) Is authentication handled by the cloud service provider.
- `pe_allowed_principal_ids` (Set of String) Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.
- `pg_config` (Attributes Set) Database configuration parameters. (see [below for nested schema](#nestedatt--pg_config))
- `private_networking` (Boolean) Is private networking enabled.
- `project_id` (String) BigAnimal Project ID.
- `service_account_ids` (Set of String) A Google Cloud Service Account is used for logs. If you leave this blank, then you will be unable to access log details for this cluster. Required when cluster is deployed on BigAnimal's cloud account.
- `tags` (Attributes Set) Assign existing tags or create tags to assign to this resource (see [below for nested schema](#nestedatt--tags))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `transparent_data_encryption` (Attributes) Transparent Data Encryption (TDE) key (see [below for nested schema](#nestedatt--transparent_data_encryption))
- `volume_snapshot_backup` (Boolean) Enable to take a snapshot of the volume.
- `wal_storage` (Attributes) Use a separate storage volume for Write-Ahead Logs (Recommended for high write workloads) (see [below for nested schema](#nestedatt--wal_storage))

### Read-Only

- `cloud_provider` (String) Cloud provider. For example, "aws", "azure", "gcp" or "bah:aws", "bah:gcp".
- `cluster_architecture` (Attributes) Cluster architecture. (see [below for nested schema](#nestedatt--cluster_architecture))
- `cluster_id` (String) Cluster ID.
- `cluster_type` (String) Type of the cluster. For example, "cluster" for biganimal_cluster resources, or "faraway_replica" for biganimal_faraway_replica resources.
- `connection_uri` (String) Cluster connection URI.
- `created_at` (String) Cluster creation time.
- `id` (String) The ID of this resource.
- `logs_url` (String) The URL to find the logs of this cluster.
- `metrics_url` (String) The URL to find the metrics of this cluster.
- `pg_identity` (String) PG Identity required to grant key permissions to activate the cluster.
- `pg_type` (String) Postgres type. For example, "epas", "pgextended", or "postgres".
- `pg_version` (String) Postgres version. See [Supported Postgres types and versions](https://www.enterprisedb.com/docs/biganimal/latest/overview/05_database_version_policy/#supported-postgres-types-and-versions) for supported Postgres types and versions.
- `phase` (String) Current phase of the cluster.
- `private_link_service_alias` (String) Private link service alias.
- `private_link_service_name` (String) private link service name.
- `resizing_pvc` (List of String) Resizing PVC.
- `transparent_data_encryption_action` (String) Transparent data encryption action.

<a id="nestedatt--storage"></a>
### Nested Schema for `storage`

Required:

- `volume_properties` (String) Volume properties.
- `volume_type` (String) Volume type.

Optional:

- `iops` (String) IOPS for the selected volume.
- `size` (String) Size of the volume.
- `throughput` (String) Throughput.


<a id="nestedatt--allowed_ip_ranges"></a>
### Nested Schema for `allowed_ip_ranges`

Required:

- `cidr_block` (String) CIDR block
- `description` (String) Description of CIDR block


<a id="nestedatt--pg_config"></a>
### Nested Schema for `pg_config`

Required:

- `name` (String) GUC name.
- `value` (String) GUC value.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `tag_name` (String)

Read-Only:

- `color` (String)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).


<a id="nestedatt--transparent_data_encryption"></a>
### Nested Schema for `transparent_data_encryption`

Required:

- `key_id` (String) Transparent Data Encryption (TDE) key ID.

Read-Only:

- `key_name` (String) Key name.
- `status` (String) Status.


<a id="nestedatt--wal_storage"></a>
### Nested Schema for `wal_storage`

Required:

- `size` (String) Size of the volume. It can be set to different values depending on your volume type and properties.
- `volume_properties` (String) Volume properties in accordance with the selected volume type.
- `volume_type` (String) Volume type. For Azure: "azurepremiumstorage" or "ultradisk". For AWS: "gp3", "io2", or "io2-block-express". For Google Cloud: only "pd-ssd".

Optional:

- `iops` (String) IOPS for the selected volume. It can be set to different values depending on your volume type and properties.
- `throughput` (String) Throughput is automatically calculated by BigAnimal based on the IOPS input if it's not provided.


<a id="nestedatt--cluster_architecture"></a>
### Nested Schema for `cluster_architecture`

Required:

- `id` (String) Cluster architecture ID.
- `nodes` (Number) Node count.

## Import

Import is supported using the following syntax:

```shell
# terraform import biganimal_faraway_replica.<resource_name> <project_id>/<cluster_id>
terraform import biganimal_faraway_replica.faraway_replica prj_deadbeef01234567/p-abcd123456
```
