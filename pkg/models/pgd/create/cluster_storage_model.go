package create

type ClusterStorage struct {
	Iops *string `json:"iops,omitempty" tfsdk:"iops"`
	Size *string `json:"size,omitempty" tfsdk:"size"`
	// Unused
	Throughput         *string `json:"throughput,omitempty" tfsdk:"throughput"`
	VolumePropertiesId string  `json:"volumePropertiesId" tfsdk:"volume_properties_id"`
	VolumeTypeId       string  `json:"volumeTypeId" tfsdk:"volume_type_id"`
}
