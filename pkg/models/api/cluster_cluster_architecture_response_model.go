package api

type ClusterClusterArchitectureResponse struct {
	ClusterArchitectureId   string   `json:"clusterArchitectureId"`
	ClusterArchitectureName string   `json:"clusterArchitectureName"`
	Nodes                   float64  `json:"nodes"`
	WitnessNodes            *float64 `json:"witnessNodes,omitempty"`
}
