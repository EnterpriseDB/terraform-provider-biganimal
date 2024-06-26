# biganimal_cluster (Data Source)
The cluster data source describes a BigAnimal cluster. The data source requires your cluster name.

## Example Usage
```terraform
variable "cluster_name" {
  type        = string
  description = "The name of the cluster"
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

data "biganimal_cluster" "this" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
}

output "cluster_architecture" {
  value = data.biganimal_cluster.this.cluster_architecture
}

output "backup_retention_period" {
  value = data.biganimal_cluster.this.backup_retention_period
}

output "cluster_name" {
  value = data.biganimal_cluster.this.cluster_name
}

output "created_at" {
  value = data.biganimal_cluster.this.created_at
}

output "csp_auth" {
  value = coalesce(data.biganimal_cluster.this.csp_auth, false)
}

output "deleted_at" {
  value = data.biganimal_cluster.this.deleted_at
}

output "expired_at" {
  value = data.biganimal_cluster.this.expired_at
}

output "instance_type" {
  value = data.biganimal_cluster.this.instance_type
}

output "metrics_url" {
  value = data.biganimal_cluster.this.metrics_url
}

output "logs_url" {
  value = data.biganimal_cluster.this.logs_url
}

output "pg_config" {
  value = data.biganimal_cluster.this.pg_config
}

output "pg_type" {
  value = data.biganimal_cluster.this.pg_type
}

output "pg_version" {
  value = data.biganimal_cluster.this.pg_version
}

output "phase" {
  value = data.biganimal_cluster.this.phase
}

output "private_networking" {
  value = coalesce(data.biganimal_cluster.this.private_networking, false)
}

output "cloud_provider" {
  value = data.biganimal_cluster.this.cloud_provider
}

output "read_only_connections" {
  value = coalesce(data.biganimal_cluster.this.read_only_connections, false)
}

output "ro_connection_uri" {
  value = data.biganimal_cluster.this.ro_connection_uri
}

output "region" {
  value = data.biganimal_cluster.this.region
}

output "resizing_pvc" {
  value = data.biganimal_cluster.this.resizing_pvc
}

output "storage" {
  value = data.biganimal_cluster.this.storage
}

output "superuser_access" {
  value = coalesce(data.biganimal_cluster.this.superuser_access, false)
}

output "pgvector" {
  value = coalesce(data.biganimal_cluster.this.pgvector, false)
}

output "post_gis" {
  value = coalesce(data.biganimal_cluster.this.post_gis, false)
}

output "faraway_replica_ids" {
  value = data.biganimal_cluster.this.faraway_replica_ids
}

output "pe_allowed_principal_ids" {
  value = data.biganimal_cluster.this.pe_allowed_principal_ids
}

output "service_account_ids" {
  value = data.biganimal_cluster.this.service_account_ids
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_name` (String) Name of the cluster.
- `project_id` (String) BigAnimal Project ID.

### Optional

- `allowed_ip_ranges` (Attributes Set) Allowed IP ranges. (see [below for nested schema](#nestedatt--allowed_ip_ranges))
- `faraway_replica_ids` (Set of String)
- `most_recent` (Boolean) Show the most recent cluster when there are multiple clusters with the same name.
- `pause` (Boolean) Pause cluster. If true it will put the cluster on pause and set the phase as paused, if false it will resume the cluster and set the phase as healthy
- `pe_allowed_principal_ids` (Set of String)
- `pg_bouncer` (Attributes) Pg bouncer. (see [below for nested schema](#nestedatt--pg_bouncer))
- `service_account_ids` (Set of String)

### Read-Only

- `backup_retention_period` (String) Backup retention period.
- `cloud_provider` (String) Cloud provider.
- `cluster_architecture` (Attributes) Cluster architecture. (see [below for nested schema](#nestedatt--cluster_architecture))
- `cluster_id` (String) Cluster ID.
- `cluster_type` (String) Type of the Specified Cluster.
- `connection_uri` (String) Cluster connection URI.
- `created_at` (String) Cluster creation time.
- `csp_auth` (Boolean) Is authentication handled by the cloud service provider.
- `deleted_at` (String) Cluster deletion time.
- `expired_at` (String) Cluster expiry time.
- `first_recoverability_point_at` (String) Earliest backup recover time.
- `id` (String) Datasource ID.
- `instance_type` (String) Instance type.
- `logs_url` (String) The URL to find the logs of this cluster.
- `maintenance_window` (Attributes) Custom maintenance window. (see [below for nested schema](#nestedatt--maintenance_window))
- `metrics_url` (String) The URL to find the metrics of this cluster.
- `pg_config` (Attributes Set) Database configuration parameters. (see [below for nested schema](#nestedatt--pg_config))
- `pg_type` (String) Postgres type.
- `pg_version` (String) Postgres version.
- `pgvector` (Boolean) Is pgvector extension enabled. Adds support for vector storage and vector similarity search to Postgres.
- `phase` (String) Current phase of the cluster.
- `post_gis` (Boolean) Is postGIS extension enabled. PostGIS extends the capabilities of the PostgreSQL relational database by adding support storing, indexing and querying geographic data.
- `private_networking` (Boolean) Is private networking enabled.
- `read_only_connections` (Boolean) Is read only connection enabled.
- `region` (String) Region to deploy the cluster.
- `resizing_pvc` (List of String) Resizing PVC.
- `ro_connection_uri` (String) Cluster read-only connection URI. Only available for high availability clusters.
- `storage` (Attributes) Storage. (see [below for nested schema](#nestedatt--storage))
- `superuser_access` (Boolean) Is superuser access enabled.

<a id="nestedatt--allowed_ip_ranges"></a>
### Nested Schema for `allowed_ip_ranges`

Read-Only:

- `cidr_block` (String) CIDR block.
- `description` (String) CIDR block description.


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



<a id="nestedatt--cluster_architecture"></a>
### Nested Schema for `cluster_architecture`

Read-Only:

- `id` (String) Cluster architecture ID.
- `name` (String) Name.
- `nodes` (Number) Node count.


<a id="nestedatt--maintenance_window"></a>
### Nested Schema for `maintenance_window`

Read-Only:

- `is_enabled` (Boolean) Is maintenance window enabled.
- `start_day` (Number) The day of week, 0 represents Sunday, 1 is Monday, and so on.
- `start_time` (String) Start time. "hh:mm", for example: "23:59".


<a id="nestedatt--pg_config"></a>
### Nested Schema for `pg_config`

Read-Only:

- `name` (String) GUC name.
- `value` (String) GUC value.


<a id="nestedatt--storage"></a>
### Nested Schema for `storage`

Read-Only:

- `iops` (String) IOPS for the selected volume.
- `size` (String) Size of the volume.
- `throughput` (String) Throughput.
- `volume_properties` (String) Volume properties.
- `volume_type` (String) Volume type.
