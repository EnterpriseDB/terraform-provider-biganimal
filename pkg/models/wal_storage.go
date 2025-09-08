package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type StorageResourceModel struct {
	VolumeType       types.String `tfsdk:"volume_type"`
	VolumeProperties types.String `tfsdk:"volume_properties"`
	Size             types.String `tfsdk:"size"`
	Iops             types.String `tfsdk:"iops"`
	Throughput       types.String `tfsdk:"throughput"`
}
