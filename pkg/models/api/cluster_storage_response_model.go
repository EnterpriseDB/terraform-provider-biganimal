package api

type ClusterStorageResponse struct {
	Iops                 string `json:"iops"`
	Size                 string `json:"size"`
	Throughput           string `json:"throughput"`
	VolumePropertiesId   string `json:"volumePropertiesId"`
	VolumePropertiesName string `json:"volumePropertiesName"`
	VolumeTypeId         string `json:"volumeTypeId"`
	VolumeTypeName       string `json:"volumeTypeName"`
}
