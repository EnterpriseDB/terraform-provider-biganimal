package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type Storage struct {
	Iops               types.String `tfsdk:"iops"`
	Size               types.String `tfsdk:"size"`
	Throughput         types.String `tfsdk:"throughput"`
	VolumePropertiesId types.String `tfsdk:"volume_properties"`
	VolumeTypeId       types.String `tfsdk:"volume_type"`
}
