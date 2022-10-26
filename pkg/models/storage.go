package models

type Storage struct {
	Iops               string `json:"iops,omitempty" mapstructure:"iops"`
	Size               string `json:"size" mapstructure:"size"`
	Throughput         string `json:"throughput,omitempty" mapstructure:"throughput"`
	VolumePropertiesId string `json:"volumePropertiesId" mapstructure:"volume_properties"`
	VolumeTypeId       string `json:"volumeTypeId" mapstructure:"volume_type"`
}

func (s Storage) PropList() PropList {
	propMap := map[string]interface{}{}
	propMap["iops"] = s.Iops
	propMap["size"] = s.Size
	propMap["throughput"] = s.Throughput
	propMap["volume_properties"] = s.VolumePropertiesId
	propMap["volume_type"] = s.VolumeTypeId
	return PropList{propMap}
}
