package models

type Architecture struct {
	ClusterArchitectureId   string `json:"clusterArchitectureId" mapstructure:"id"`
	ClusterArchitectureName string `json:"clusterArchitectureName,omitempty"`
	Nodes                   int    `json:"nodes" mapstructure:"nodes"`
}
