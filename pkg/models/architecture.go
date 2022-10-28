package models

type Architecture struct {
	ClusterArchitectureId   string `json:"clusterArchitectureId" mapstructure:"id"`
	ClusterArchitectureName string `json:"clusterArchitectureName,omitempty" mapstructure:"name,omitempty"`
	Nodes                   int    `json:"nodes" mapstructure:"nodes"`
}
