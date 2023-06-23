variable "cluster_name" {
  type        = string
  description = "The name of the cluster"
}

variable "project_id" {
  type        = string
  description = "BigAnimal Project ID"
}

data "biganimal_pgd" "this" {
  cluster_name = var.cluster_name
  project_id   = var.project_id
}

output "project_id" {
  value = data.biganimal_pgd.this.project_id
}

output "cluster_name" {
  value = data.biganimal_pgd.this.cluster_name
}

output "data_groups" {
  value = data.biganimal_pgd.this.data_groups
}

output "witness_groups" {
  value = data.biganimal_pgd.this.witness_groups
}
