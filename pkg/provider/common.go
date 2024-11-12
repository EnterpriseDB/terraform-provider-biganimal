package provider

import (
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var WeekdaysNumber = map[string]float64{
	"monday":    1.0,
	"tuesday":   2.0,
	"wednesday": 3.0,
	"thursday":  4.0,
	"friday":    5.0,
	"saturday":  6.0,
	"sunday":    0.0,
}

var WeekdaysName = map[float64]string{
	1.0: "Monday",
	2.0: "Tuesday",
	3.0: "Wednesday",
	4.0: "Thursday",
	5.0: "Friday",
	6.0: "Saturday",
	0.0: "Sunday",
}

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

var resourceBackupSchedule = schema.SingleNestedAttribute{
	Description: "Backup schedule.",
	Optional:    true,
	Attributes: map[string]schema.Attribute{
		"start_day": schema.StringAttribute{
			Description: "Backup schedule start day.",
			Required:    true,
		},
		"start_time": schema.StringAttribute{
			Description: "Backup schedule start time.",
			Required:    true,
		},
	},
}
