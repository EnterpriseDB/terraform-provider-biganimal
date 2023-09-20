package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type ClusterArchitecture struct {
	ClusterArchitectureId   string       `tfsdk:"cluster_architecture_id"`
	ClusterArchitectureName types.String `tfsdk:"cluster_architecture_name"`
	Nodes                   float64      `tfsdk:"nodes"`
	WitnessNodes            types.Int64  `tfsdk:"witness_nodes"`
}
