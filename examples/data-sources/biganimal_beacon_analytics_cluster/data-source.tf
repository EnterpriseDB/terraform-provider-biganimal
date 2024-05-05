variable "cluster_name" {
  type        = string
  description = "The name of the cluster"
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

data "biganimal_beacon_analytics_cluster" "this" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
}

output "backup_retention_period" {
  value = data.biganimal_beacon_analytics_cluster.this.backup_retention_period
}

output "cluster_name" {
  value = data.biganimal_beacon_analytics_cluster.this.cluster_name
}

output "created_at" {
  value = data.biganimal_beacon_analytics_cluster.this.created_at
}

output "csp_auth" {
  value = coalesce(data.biganimal_beacon_analytics_cluster.this.csp_auth, false)
}

output "deleted_at" {
  value = data.biganimal_beacon_analytics_cluster.this.deleted_at
}

output "expired_at" {
  value = data.biganimal_beacon_analytics_cluster.this.expired_at
}

output "instance_type" {
  value = data.biganimal_beacon_analytics_cluster.this.instance_type
}

output "metrics_url" {
  value = data.biganimal_beacon_analytics_cluster.this.metrics_url
}

output "logs_url" {
  value = data.biganimal_beacon_analytics_cluster.this.logs_url
}

output "pg_type" {
  value = data.biganimal_beacon_analytics_cluster.this.pg_type
}

output "pg_version" {
  value = data.biganimal_beacon_analytics_cluster.this.pg_version
}

output "phase" {
  value = data.biganimal_beacon_analytics_cluster.this.phase
}

output "private_networking" {
  value = coalesce(data.biganimal_beacon_analytics_cluster.this.private_networking, false)
}

output "cloud_provider" {
  value = data.biganimal_beacon_analytics_cluster.this.cloud_provider
}

output "region" {
  value = data.biganimal_beacon_analytics_cluster.this.region
}

output "resizing_pvc" {
  value = data.biganimal_beacon_analytics_cluster.this.resizing_pvc
}

output "storage" {
  value = data.biganimal_beacon_analytics_cluster.this.storage
}

output "pe_allowed_principal_ids" {
  value = data.biganimal_cluster.this.pe_allowed_principal_ids
}

output "service_account_ids" {
  value = data.biganimal_cluster.this.service_account_ids
}
