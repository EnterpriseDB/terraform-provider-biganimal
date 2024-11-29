package provider

import (
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// build tag assign terraform resource as, using api response as input
func buildTFRsrcAssignTagsAs(tfRsrcTagsOut *[]commonTerraform.Tag, apiRespTags []commonApi.Tag) {
	*tfRsrcTagsOut = []commonTerraform.Tag{}
	for _, v := range apiRespTags {
		*tfRsrcTagsOut = append(*tfRsrcTagsOut, commonTerraform.Tag{
			TagId:   types.StringValue(v.TagId),
			TagName: types.StringValue(v.TagName),
			Color:   basetypes.NewStringPointerValue(v.Color),
		})
	}
}

// build tag assign request using terraform resource as input
func buildAPIReqAssignTags(tfRsrcTags []commonTerraform.Tag) []commonApi.Tag {
	tags := []commonApi.Tag{}
	for _, tag := range tfRsrcTags {
		tags = append(tags, commonApi.Tag{
			Color:   tag.Color.ValueStringPointer(),
			TagId:   tag.TagId.ValueString(),
			TagName: tag.TagName.ValueString(),
		})
	}
	return tags
}

var ResourceBackupScheduleTime = schema.StringAttribute{
	MarkdownDescription: "Backup schedule time in 24 hour cron expression format.",
	Optional:            true,
	Computed:            true,
}

var resourceWal = schema.SingleNestedAttribute{
	Description: "Use a separate storage volume for Write-Ahead Logs (Recommended for high write workloads)",
	Optional:    true,
	Attributes: map[string]schema.Attribute{
		"iops": schema.StringAttribute{
			Description:   "IOPS for the selected volume. It can be set to different values depending on your volume type and properties.",
			Optional:      true,
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"size": schema.StringAttribute{
			Description:   "Size of the volume. It can be set to different values depending on your volume type and properties.",
			Required:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"throughput": schema.StringAttribute{
			Description:   "Throughput is automatically calculated by BigAnimal based on the IOPS input if it's not provided.",
			Optional:      true,
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"volume_properties": schema.StringAttribute{
			Description: "Volume properties in accordance with the selected volume type.",
			Required:    true,
		},
		"volume_type": schema.StringAttribute{
			Description: "Volume type. For Azure: \"azurepremiumstorage\" or \"ultradisk\". For AWS: \"gp3\", \"io2\", or \"io2-block-express\". For Google Cloud: only \"pd-ssd\".",
			Required:    true,
		},
	},
}
