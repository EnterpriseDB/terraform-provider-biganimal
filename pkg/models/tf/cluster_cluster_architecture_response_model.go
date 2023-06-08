package tf

type ClusterClusterArchitectureResponse struct {
	ClusterArchitectureId   string   `tfsdk:"cluster_architecture_id"`
	ClusterArchitectureName string   `tfsdk:"cluster_architecture_name"`
	Nodes                   float64  `tfsdk:"nodes"`
	WitnessNodes            *float64 `tfsdk:"witness_nodes"`
}
