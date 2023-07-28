package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type Storage struct {
	Iops               types.String `tfsdk:"iops"`
	Size               *string      `tfsdk:"size"`
	Throughput         types.String `tfsdk:"throughput"`
	VolumePropertiesId *string      `tfsdk:"volume_properties"`
	VolumeTypeId       *string      `tfsdk:"volume_type"`
}
