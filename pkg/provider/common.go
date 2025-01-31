package provider

import (
	"context"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// validate config tags. Add error if invalid
func ValidateTags(ctx context.Context, tagClient *api.TagClient, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	set := new(types.Set)
	req.Plan.GetAttribute(ctx, path.Root("tags"), set)

	configTags := []terraform.Tag{}
	diag := set.ElementsAs(ctx, &configTags, false)
	if diag.ErrorsCount() > 0 {
		resp.Diagnostics.Append(diag...)
		return
	}

	// Validate existing tag. Existing tag colors cannot be changed in a cluster create request and must be removed.
	// To change tag color, use tag request
	existingTags, err := tagClient.TagClient().List(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error fetching existing tags", err.Error())
	}
	for _, configTag := range configTags {
		for _, existingTag := range existingTags {
			// if config tag matches existing tag, then config tags color has to match existing tag color or
			// config tag color should be set to nil, other throw validation error
			if existingTag.TagName == configTag.TagName.ValueString() &&
				(existingTag.Color != nil && *existingTag.Color != configTag.Color.ValueString() ||
					configTag.Color.ValueStringPointer() != nil) {

				existingColor := "nil"
				if existingTag.Color != nil {
					existingColor = *existingTag.Color
				}

				resp.Diagnostics.AddError("An existing tag's color cannot be changed",
					fmt.Sprintf("Please remove the color field for tag: `%v` or set it to the existing tag's color: `%v`.\nTo change an existing tag's color please use resource `biganimal_tag`",
						configTag.TagName.ValueString(), existingColor))
			}
		}
	}
}

// build tag assign terraform resource as, using api response as input
func buildTfRsrcTagsAs(tfRsrcTagsOut *[]commonTerraform.Tag, apiRespTags []commonApi.Tag) {
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
func buildApiReqTags(tfRsrcTags []commonTerraform.Tag) []commonApi.Tag {
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
