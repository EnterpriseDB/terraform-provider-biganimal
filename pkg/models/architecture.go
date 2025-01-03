package models

type Architecture struct {
	ClusterArchitectureId string `json:"clusterArchitectureId" mapstructure:"id"`
	Nodes                 int    `json:"nodes" mapstructure:"nodes"`
}
