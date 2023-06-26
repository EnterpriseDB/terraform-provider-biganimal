package pgd

type ClusterArchitecture struct {
	ClusterArchitectureId   string   `json:"clusterArchitectureId" tfsdk:"cluster_architecture_id"`
	ClusterArchitectureName *string  `json:"clusterArchitectureName,omitempty" tfsdk:"cluster_architecture_name"`
	Nodes                   float64  `json:"nodes" tfsdk:"nodes"`
	WitnessNodes            *float64 `json:"witnessNodes,omitempty" tfsdk:"witness_nodes"`
}
