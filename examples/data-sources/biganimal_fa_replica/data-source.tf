variable "cluster_name" {
  type        = string
  description = "The name of the cluster"
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

data "biganimal_fa_replica" "this" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
}

output "source_cluster_id" {
  value = data.biganimal_fa_replica.this.source_cluster_id
}

output "cluster_type" {
  value = data.biganimal_fa_replica.this.cluster_type
}

output "cluster_architecture" {
  value = data.biganimal_fa_replica.this.cluster_architecture
}

output "backup_retention_period" {
  value = data.biganimal_fa_replica.this.backup_retention_period
}

output "cluster_name" {
  value = data.biganimal_fa_replica.this.cluster_name
}

output "created_at" {
  value = data.biganimal_fa_replica.this.created_at
}

output "csp_auth" {
  value = coalesce(data.biganimal_fa_replica.this.csp_auth, false)
}

output "deleted_at" {
  value = data.biganimal_fa_replica.this.deleted_at
}

output "expired_at" {
  value = data.biganimal_fa_replica.this.expired_at
}

output "instance_type" {
  value = data.biganimal_fa_replica.this.instance_type
}

output "metrics_url" {
  value = data.biganimal_fa_replica.this.metrics_url
}

output "logs_url" {
  value = data.biganimal_fa_replica.this.logs_url
}

output "pg_config" {
  value = data.biganimal_fa_replica.this.pg_config
}

output "pg_type" {
  value = data.biganimal_fa_replica.this.pg_type
}

output "pg_version" {
  value = data.biganimal_fa_replica.this.pg_version
}

output "phase" {
  value = data.biganimal_fa_replica.this.phase
}

output "private_networking" {
  value = coalesce(data.biganimal_fa_replica.this.private_networking, false)
}

output "cloud_provider" {
  value = data.biganimal_fa_replica.this.cloud_provider
}

//output "read_only_connections" {
//  value = coalesce(data.biganimal_fa_replica.this.read_only_connections, false)
//}
//
//output "ro_connection_uri" {
//  value = data.biganimal_fa_replica.this.ro_connection_uri
//}

output "region" {
  value = data.biganimal_fa_replica.this.region
}

output "resizing_pvc" {
  value = data.biganimal_fa_replica.this.resizing_pvc
}

output "storage" {
  value = data.biganimal_fa_replica.this.storage
}
