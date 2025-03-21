# biganimal_cluster (Resource)

The cluster resource is used to manage BigAnimal clusters. See [Creating a cluster](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/) for more details.



## Single Node Cluster Example

Please visit the [examples page](https://github.com/EnterpriseDB/terraform-provider-biganimal/tree/main/examples#biganimal_cluster-resource-examples) for more single node cluster examples on various cloud service providers.

```terraform
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

  instance_type = "aws:m6i.large"
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
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }

  #  wal_storage = {
  #    volume_type       = "gp3"
  #    volume_properties = "gp3"
  #    size              = "4 Gi"
  #  }

  maintenance_window = {
    is_enabled = true
    start_day  = 6
    start_time = "03:00"
  }

  pg_type                = "epas" #valid values ["epas", "pgextended", "postgres]"
  pg_version             = "15"
  private_networking     = false
  cloud_provider         = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
  read_only_connections  = false
  region                 = "us-east-1"
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
  #   <example_value> # ex: 123456789012
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
```

## Primary/Standby High Availability Cluster Example
```terraform
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

resource "biganimal_cluster" "ha_cluster" {
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
    id    = "ha"
    nodes = 3
  }

  instance_type = "aws:c6i.large"
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
    volume_type       = "gp3"
    volume_properties = "gp3"
    size              = "4 Gi"
  }

  #  wal_storage = {
  #    volume_type       = "gp3"
  #    volume_properties = "gp3"
  #    size              = "4 Gi"
  #  }

  maintenance_window = {
    is_enabled = true
    start_day  = 6
    start_time = "03:00"
  }

  pg_type                = "epas" #valid values ["epas", "pgextended", "postgres]"
  pg_version             = "15"
  private_networking     = false
  cloud_provider         = "bah:aws" // "bah:aws" uses BigAnimal's cloud account AWS, use "aws" for your cloud account
  read_only_connections  = false
  region                 = "us-east-1"
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

  # transparent_data_encryption = {
  #   key_id = <example_value>
  # }
}

output "password" {
  sensitive = true
  value     = resource.biganimal_cluster.ha_cluster.password
}

output "ro_connection_uri" {
  value = resource.biganimal_cluster.ha_cluster.ro_connection_uri
}

output "faraway_replica_ids" {
  value = resource.biganimal_cluster.ha_cluster.faraway_replica_ids
}
```

## Distributed High Availability Cluster Example

-> Please use the `biganimal_pgd` resource to manage the Distributed High Availability clusters.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_provider` (String) Cloud provider. For example, "aws", "azure", "gcp" or "bah:aws", "bah:gcp".
- `cluster_architecture` (Attributes) Cluster architecture. (see [below for nested schema](#nestedatt--cluster_architecture))
- `cluster_name` (String) Name of the cluster.
- `instance_type` (String) Instance type. For example, "azure:Standard_D2s_v3", "aws:c6i.large" or "gcp:e2-highcpu-4".
- `password` (String) Password for the user edb_admin. It must be 12 characters or more.
- `pg_type` (String) Postgres type. For example, "epas", "pgextended", or "postgres".
- `pg_version` (String) Postgres version. See [Supported Postgres types and versions](https://www.enterprisedb.com/docs/biganimal/latest/overview/05_database_version_policy/#supported-postgres-types-and-versions) for supported Postgres types and versions.
- `project_id` (String) BigAnimal Project ID.
- `region` (String) Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.
- `storage` (Attributes) Storage. (see [below for nested schema](#nestedatt--storage))

### Optional

- `allowed_ip_ranges` (Attributes Set) Allowed IP ranges. (see [below for nested schema](#nestedatt--allowed_ip_ranges))
- `backup_retention_period` (String) Backup retention period. For example, "7d", "2w", or "3m".
- `backup_schedule_time` (String) Backup schedule time in 24 hour cron expression format.
- `csp_auth` (Boolean) Is authentication handled by the cloud service provider. Available for AWS only, See [Authentication](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/#authentication) for details.
- `maintenance_window` (Attributes) Custom maintenance window. (see [below for nested schema](#nestedatt--maintenance_window))
- `pause` (Boolean) Pause cluster. If true it will put the cluster on pause and set the phase as paused, if false it will resume the cluster and set the phase as healthy. Pausing a cluster allows you to save on compute costs without losing data or cluster configuration settings. While paused, clusters aren't upgraded or patched, but changes are applied when the cluster resumes. Pausing a high availability cluster shuts down all cluster nodes
- `pe_allowed_principal_ids` (Set of String) Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.
- `pg_bouncer` (Attributes) Pg bouncer. (see [below for nested schema](#nestedatt--pg_bouncer))
- `pg_config` (Attributes Set) Database configuration parameters. See [Modifying database configuration parameters](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/03_modifying_your_cluster/05_db_configuration_parameters/) for details. (see [below for nested schema](#nestedatt--pg_config))
- `pgvector` (Boolean) Is pgvector extension enabled. Adds support for vector storage and vector similarity search to Postgres.
- `post_gis` (Boolean) Is postGIS extension enabled. PostGIS extends the capabilities of the PostgreSQL relational database by adding support storing, indexing and querying geographic data.
- `private_networking` (Boolean) Is private networking enabled.
- `read_only_connections` (Boolean) Is read only connection enabled.
- `service_account_ids` (Set of String) A Google Cloud Service Account is used for logs. If you leave this blank, then you will be unable to access log details for this cluster. Required when cluster is deployed on BigAnimal's cloud account.
- `superuser_access` (Boolean) Enable to grant superuser access to the edb_admin role.
- `tags` (Attributes Set) Assign existing tags or create tags to assign to this resource (see [below for nested schema](#nestedatt--tags))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `transparent_data_encryption` (Attributes) Transparent Data Encryption (TDE) key (see [below for nested schema](#nestedatt--transparent_data_encryption))
- `volume_snapshot_backup` (Boolean) Enable to take a snapshot of the volume.
- `wal_storage` (Attributes) Use a separate storage volume for Write-Ahead Logs (Recommended for high write workloads) (see [below for nested schema](#nestedatt--wal_storage))

### Read-Only

- `cluster_id` (String) Cluster ID.
- `cluster_type` (String) Type of the cluster. For example, "cluster" for biganimal_cluster resources, or "faraway_replica" for biganimal_faraway_replica resources.
- `connection_uri` (String) Cluster connection URI.
- `created_at` (String) Cluster creation time.
- `faraway_replica_ids` (Set of String)
- `first_recoverability_point_at` (String) Earliest backup recover time.
- `id` (String) Resource ID of the cluster.
- `logs_url` (String) The URL to find the logs of this cluster.
- `metrics_url` (String) The URL to find the metrics of this cluster.
- `pg_identity` (String) PG Identity required to grant key permissions to activate the cluster.
- `phase` (String) Current phase of the cluster.
- `private_link_service_alias` (String) Private link service alias.
- `private_link_service_name` (String) private link service name.
- `resizing_pvc` (List of String) Resizing PVC.
- `ro_connection_uri` (String) Cluster read-only connection URI. Only available for high availability clusters.
- `service_name` (String) Cluster connection service name.
- `transparent_data_encryption_action` (String) Transparent data encryption action.

<a id="nestedatt--cluster_architecture"></a>
### Nested Schema for `cluster_architecture`

Required:

- `id` (String) Cluster architecture ID. For example, "single" or "ha".For Extreme High Availability clusters, please use the [biganimal_pgd](https://registry.terraform.io/providers/EnterpriseDB/biganimal/latest/docs/resources/pgd) resource.
- `nodes` (Number) Node count.


<a id="nestedatt--storage"></a>
### Nested Schema for `storage`

Required:

- `size` (String) Size of the volume. It can be set to different values depending on your volume type and properties.
- `volume_properties` (String) Volume properties in accordance with the selected volume type.
- `volume_type` (String) Volume type. For Azure: "azurepremiumstorage" or "ultradisk". For AWS: "gp3", "io2", or "io2-block-express". For Google Cloud: only "pd-ssd".

Optional:

- `iops` (String) IOPS for the selected volume. It can be set to different values depending on your volume type and properties.
- `throughput` (String) Throughput is automatically calculated by BigAnimal based on the IOPS input if it's not provided.


<a id="nestedatt--allowed_ip_ranges"></a>
### Nested Schema for `allowed_ip_ranges`

Required:

- `cidr_block` (String) CIDR block

Optional:

- `description` (String) Description of CIDR block


<a id="nestedatt--maintenance_window"></a>
### Nested Schema for `maintenance_window`

Required:

- `is_enabled` (Boolean) Is maintenance window enabled.

Optional:

- `start_day` (Number) The day of week, 0 represents Sunday, 1 is Monday, and so on.
- `start_time` (String) Start time. "hh:mm", for example: "23:59".


<a id="nestedatt--pg_bouncer"></a>
### Nested Schema for `pg_bouncer`

Required:

- `is_enabled` (Boolean) Is pg bouncer enabled.

Optional:

- `settings` (Attributes Set) PgBouncer Configuration Settings. (see [below for nested schema](#nestedatt--pg_bouncer--settings))

<a id="nestedatt--pg_bouncer--settings"></a>
### Nested Schema for `pg_bouncer.settings`

Required:

- `name` (String) Name.
- `operation` (String) Operation.
- `value` (String) Value.



<a id="nestedatt--pg_config"></a>
### Nested Schema for `pg_config`

Required:

- `name` (String) GUC name.
- `value` (String) GUC value.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `tag_name` (String)

Optional:

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

## Import

Import is supported using the following syntax:

```shell
# terraform import biganimal_cluster.<resource_name> <project_id>/<cluster_id>
terraform import biganimal_cluster.single_node_cluster prj_deadbeef01234567/p-abcd123456
```
