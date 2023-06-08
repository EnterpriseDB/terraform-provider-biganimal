package api

type ClusterCreateWitnessGroup struct {
	ClusterArchitecture *ClusterClusterArchitecture `json:"clusterArchitecture"`
	ClusterType         string                      `json:"clusterType"`
	InstanceType        *ClusterInstanceType        `json:"instanceType"`
	Provider            *ClusterCloudProvider       `json:"provider"`
	Region              *ClusterRegion              `json:"region"`
	Storage             *ClusterStorage             `json:"storage"`
}
