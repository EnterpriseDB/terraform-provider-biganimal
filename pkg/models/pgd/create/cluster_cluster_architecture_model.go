package create

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd"

type ClusterClusterArchitecture struct {
	ClusterArchitectureId string                              `json:"clusterArchitectureId" tfsdk:"cluster_architecture_id"`
	Nodes                 float64                             `json:"nodes" tfsdk:"nodes"`
	Params                *[]pgd.ArrayOfNameValueObjectsInner `json:"params,omitempty" tfsdk:"params"`
	WitnessNodes          *float64                            `json:"witnessNodes,omitempty" tfsdk:"witness_nodes"`
}
