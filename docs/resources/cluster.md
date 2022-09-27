---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "biganimal_cluster Resource - terraform-provider-biganimal"
subcategory: ""
description: |-
  Create a Postgres Cluster
---

# biganimal_cluster (Resource)

Create a Postgres Cluster



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_provider` (String) Cloud Provider
- `cluster_architecture` (Block List, Min: 1) Cluster Architecture (see [below for nested schema](#nestedblock--cluster_architecture))
- `cluster_name` (String) Name of the cluster.
- `instance_type_id` (String) Cluster Expiry Time
- `password` (String, Sensitive) Password
- `pg_type` (String) Postgres type
- `pg_version` (String) Postgres Version
- `region` (String) Region
- `replicas` (Number) Replicas
- `storage` (Block List, Min: 1) Storage (see [below for nested schema](#nestedblock--storage))

### Optional

- `allowed_ip_ranges` (Block List) Allowed IP ranges (see [below for nested schema](#nestedblock--allowed_ip_ranges))
- `backup_retention_period` (String) Backup Retention Period.
- `pg_config` (Block List) Instance Type (see [below for nested schema](#nestedblock--pg_config))
- `private_networking` (Boolean) Is private networking enabled
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `first_recoverability_point_at` (String) Cluster Expiry Time
- `id` (String) cluster ID

<a id="nestedblock--cluster_architecture"></a>
### Nested Schema for `cluster_architecture`

Required:

- `id` (String) ID
- `nodes` (Number) Node Count

Read-Only:

- `name` (String) Name


<a id="nestedblock--storage"></a>
### Nested Schema for `storage`

Required:

- `size` (String) Size
- `volume_properties` (String) Volume Properties
- `volume_type` (String) Volume Type

Optional:

- `iops` (String) IOPS
- `throughput` (String) Throughput


<a id="nestedblock--allowed_ip_ranges"></a>
### Nested Schema for `allowed_ip_ranges`

Required:

- `cidr_block` (String) CIDR Block

Optional:

- `description` (String) CIDR Block Description


<a id="nestedblock--pg_config"></a>
### Nested Schema for `pg_config`

Required:

- `name` (String) GUC Name
- `value` (String) GUC Value


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)

