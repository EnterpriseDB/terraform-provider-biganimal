package api

type ClusterWitnessGroup struct {
	ClusterName  string                           `json:"clusterName"`
	ClusterType  string                           `json:"clusterType"`
	Conditions   *[]ClusterConditionsInner        `json:"conditions,omitempty"`
	CreatedAt    *PointInTime                     `json:"createdAt,omitempty"`
	DeletedAt    *PointInTime                     `json:"deletedAt,omitempty"`
	GroupId      string                           `json:"groupId"`
	InstanceType *CloudProviderRegionInstanceType `json:"instanceType,omitempty"`
	Phase        *string                          `json:"phase,omitempty"`
	Provider     *CloudProvider                   `json:"provider"`
	Region       *CloudProviderRegion             `json:"region"`
	Storage      *ClusterStorageResponse          `json:"storage,omitempty"`
}
