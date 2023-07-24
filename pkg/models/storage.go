package models

type Storage struct {
	Iops               *string `json:"iops,omitempty" mapstructure:"iops" tfsdk:"iops"`
	Size               *string `json:"size" mapstructure:"size" tfsdk:"size"`
	Throughput         *string `json:"throughput,omitempty" mapstructure:"throughput" tfsdk:"throughput"`
	VolumePropertiesId *string `json:"volumePropertiesId" mapstructure:"volume_properties" tfsdk:"volume_properties"`
	VolumeTypeId       *string `json:"volumeTypeId" mapstructure:"volume_type" tfsdk:"volume_type"`
}
