package api

type ClusterClusterArchitecture struct {
	ClusterArchitectureId string                          `json:"clusterArchitectureId"`
	Nodes                 float64                         `json:"nodes"`
	Params                *[]ArrayOfNameValueObjectsInner `json:"params,omitempty"`
	WitnessNodes          *float64                        `json:"witnessNodes,omitempty"`
}
