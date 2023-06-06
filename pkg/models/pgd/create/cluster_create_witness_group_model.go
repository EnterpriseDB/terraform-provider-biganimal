package create

type ClusterCreateWitnessGroup struct {
	ClusterArchitecture *ClusterClusterArchitecture `json:"clusterArchitecture" tfsdk:"cluster_architecture"`
	ClusterType         string                      `json:"clusterType" tfsdk:"cluster_type"`
	InstanceType        *ClusterInstanceType        `json:"instanceType" tfsdk:"instance_type"`
	Provider            *ClusterCloudProvider       `json:"provider" tfsdk:"provider"`
	Region              *ClusterRegion              `json:"region" tfsdk:"region"`
	Storage             *ClusterStorage             `json:"storage" tfsdk:"storage"`
}
