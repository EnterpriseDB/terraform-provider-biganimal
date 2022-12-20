variable "cluster_name" {
  type        = string
  description = "The name of the cluster"
}

data "biganimal_cluster" "this" {
  cluster_name = var.cluster_name
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
