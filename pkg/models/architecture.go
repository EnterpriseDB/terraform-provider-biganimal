package models

type Architecture struct {
	ClusterArchitectureId   string  `json:"clusterArchitectureId" mapstructure:"id"`
	ClusterArchitectureName string  `json:"clusterArchitectureName,omitempty" mapstructure:"name"`
	Nodes                   float64 `json:"nodes" mapstructure:"nodes"`
}
