package api

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"

type WitnessGroup struct {
	GroupId             *string              `json:"groupId,omitempty" tfsdk:"group_id"`
	ClusterArchitecture *ClusterArchitecture `json:"clusterArchitecture,omitempty" tfsdk:"cluster_architecture"`
	ClusterType         *string              `json:"clusterType,omitempty" tfsdk:"cluster_type"`
	InstanceType        *InstanceType        `json:"instanceType,omitempty" tfsdk:"instance_type"`
	Provider            *CloudProvider       `json:"provider,omitempty" tfsdk:"cloud_provider"`
	Region              *Region              `json:"region,omitempty" tfsdk:"region"`
	Storage             *models.Storage      `json:"storage,omitempty" tfsdk:"storage"`
	Phase               *string              `json:"phase,omitempty" tfsdk:"phase"`
}
