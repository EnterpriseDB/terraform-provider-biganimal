package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type WitnessGroup struct {
	GroupId             types.String `tfsdk:"group_id"`
	ClusterArchitecture types.Object `tfsdk:"cluster_architecture"`
	ClusterType         types.String `tfsdk:"cluster_type"`
	InstanceType        types.Object `tfsdk:"instance_type"`
	Provider            types.Object `tfsdk:"cloud_provider"`
	Region              *Region      `tfsdk:"region"`
	Storage             types.Object `tfsdk:"storage"`
	Phase               types.String `tfsdk:"phase"`
}
