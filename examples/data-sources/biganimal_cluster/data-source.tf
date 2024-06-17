variable "cluster_id" {
  type        = string
  description = "The id of the cluster"
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

data "biganimal_cluster" "this" {
  cluster_id = var.cluster_id
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
