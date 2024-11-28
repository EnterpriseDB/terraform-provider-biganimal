variable "cluster_id" {
  type        = string
  description = "Id of the faraway replica cluster"
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

data "biganimal_faraway_replica" "this" {
  cluster_id = var.cluster_id
  project_id = var.project_id
}

output "source_cluster_id" {
  value = data.biganimal_faraway_replica.this.source_cluster_id
}

output "cluster_type" {
  value = data.biganimal_faraway_replica.this.cluster_type
}

output "cluster_architecture" {
  value = data.biganimal_faraway_replica.this.cluster_architecture
}

output "backup_retention_period" {
  value = data.biganimal_faraway_replica.this.backup_retention_period
}

output "backup_schedule_time" {
  value = data.biganimal_faraway_replica.this.backup_schedule_time
}

output "cluster_name" {
  value = data.biganimal_faraway_replica.this.cluster_name
}

output "created_at" {
  value = data.biganimal_faraway_replica.this.created_at
}

output "csp_auth" {
  value = coalesce(data.biganimal_faraway_replica.this.csp_auth, false)
}

output "instance_type" {
  value = data.biganimal_faraway_replica.this.instance_type
}

output "metrics_url" {
  value = data.biganimal_faraway_replica.this.metrics_url
}

output "logs_url" {
  value = data.biganimal_faraway_replica.this.logs_url
}

output "pg_config" {
  value = data.biganimal_faraway_replica.this.pg_config
}

output "pg_type" {
  value = data.biganimal_faraway_replica.this.pg_type
}

output "pg_version" {
  value = data.biganimal_faraway_replica.this.pg_version
}

output "phase" {
  value = data.biganimal_faraway_replica.this.phase
}

output "private_networking" {
  value = coalesce(data.biganimal_faraway_replica.this.private_networking, false)
}

output "cloud_provider" {
  value = data.biganimal_faraway_replica.this.cloud_provider
}

output "region" {
  value = data.biganimal_faraway_replica.this.region
}

output "resizing_pvc" {
  value = data.biganimal_faraway_replica.this.resizing_pvc
}

output "storage" {
  value = data.biganimal_faraway_replica.this.storage
}

output "wal_storage" {
  value = data.biganimal_faraway_replica.this.wal_storage
}

output "volume_snapshot_backup" {
  value = data.biganimal_faraway_replica.this.volume_snapshot_backup
}
