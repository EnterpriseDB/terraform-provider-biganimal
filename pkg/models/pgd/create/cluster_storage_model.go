package create

type ClusterStorage struct {
	Iops *string `json:"iops,omitempty"`
	Size *string `json:"size,omitempty"`
	// Unused
	Throughput         *string `json:"throughput,omitempty"`
	VolumePropertiesId string  `json:"volumePropertiesId"`
	VolumeTypeId       string  `json:"volumeTypeId"`
}
