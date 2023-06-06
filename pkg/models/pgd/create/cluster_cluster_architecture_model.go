package create

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd"

type ClusterClusterArchitecture struct {
	ClusterArchitectureId string                              `json:"clusterArchitectureId"`
	Nodes                 float64                             `json:"nodes"`
	Params                *[]pgd.ArrayOfNameValueObjectsInner `json:"params,omitempty"`
	WitnessNodes          *float64                            `json:"witnessNodes,omitempty"`
}
