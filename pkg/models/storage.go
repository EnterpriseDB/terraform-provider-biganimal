package models

type Storage struct {
	Iops               string `json:"iops,omitempty" mapstructure:"iops"`
	Size               string `json:"size" mapstructure:"size"`
	Throughput         string `json:"throughput,omitempty" mapstructure:"throughput"`
	VolumePropertiesId string `json:"volumePropertiesId" mapstructure:"volume_properties"`
	VolumeTypeId       string `json:"volumeTypeId" mapstructure:"volume_type"`
}
