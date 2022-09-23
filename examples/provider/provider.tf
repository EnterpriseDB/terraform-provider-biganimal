provider "biganimal" {
  # example configuration here
  version = "0.3.1"
}

data "biganimal_cluster" "example" {
  name = "nicktest"
}

output "curr_pri" {
  value = data.biganimal_cluster.example.current_primary
}

output "phase" {
  value = data.biganimal_cluster.example.phase
}