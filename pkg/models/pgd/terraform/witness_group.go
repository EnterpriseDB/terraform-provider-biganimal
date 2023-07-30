package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type WitnessGroup struct {
	GroupId             types.String         `tfsdk:"group_id"`
	ClusterArchitecture *ClusterArchitecture `tfsdk:"cluster_architecture"`
	ClusterType         types.String         `tfsdk:"cluster_type"`
	InstanceType        *InstanceType        `tfsdk:"instance_type"`
	Provider            *CloudProvider       `tfsdk:"cloud_provider"`
	Region              *Region              `tfsdk:"region"`
	Storage             *Storage             `tfsdk:"storage"`
	Phase               types.String         `tfsdk:"phase"`
}
