package provider

import (
	"context"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// validate config tags. Add error if invalid
func ValidateTags(ctx context.Context, tagClient *api.TagClient, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	set := new(types.Set)
	req.Config.GetAttribute(ctx, path.Root("tags"), set)

	configTags := []terraform.Tag{}
	diag := set.ElementsAs(ctx, &configTags, false)
	if diag.ErrorsCount() > 0 {
		resp.Diagnostics.Append(diag...)
		return
	}

	// check for tag duplicates in config
	checkDupes := make(map[string]struct{})
	for _, configTag := range configTags {
		checkDupes[configTag.TagName.ValueString()] = struct{}{}
	}

	if len(checkDupes) != len(configTags) {
		resp.Diagnostics.AddError("Duplicate tag_name not allowed", "Please remove duplicate tag_name in tags")
		return
	}

	// Validate existing tag. Existing tag colors cannot be changed in a cluster request and must be removed.
	// To change tag color, use tag request
	existingTags, err := tagClient.TagClient().List(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error fetching existing tags", err.Error())
		return
	}
	for _, configTag := range configTags {
		for _, existingTag := range existingTags {
			// this sets color to empty string if color is nil so we don't need to handle nil case separately
			if existingTag.Color == nil {
				existingTag.Color = utils.ToPointer("")
			}

			// if config tag matches existing tag, then config tags color has to match existing tag color or
			// config tag color should be removed, otherwise throw a validation error
			// color is a computed value so color unknown means color is removed from config
			if existingTag.TagName == configTag.TagName.ValueString() && !configTag.Color.IsNull() &&
				existingTag.Color != nil && *existingTag.Color != configTag.Color.ValueString() {

				resp.Diagnostics.AddError("An existing tag's color cannot be changed to another color when using this resource.",
					fmt.Sprintf("Please remove the color field for tag: \"%v\" or set it to the existing tag's color: \"%v\".\nTo change an existing tag's color please use resource `biganimal_tag`.",
						configTag.TagName.ValueString(), *existingTag.Color))
				return
			}
		}
	}
}

// build tag assign terraform resource as, using api response as input
func buildTfRsrcTagsAs(tfRsrcTagsOut *[]commonTerraform.Tag, apiRespTags []commonApi.Tag) {
	*tfRsrcTagsOut = []commonTerraform.Tag{}
	for _, v := range apiRespTags {
		color := v.Color
		// if color is nil, set it to empty string. This is to handle the case where color is nil in api response
		if v.Color == nil {
			color = utils.ToPointer("")
		}
		*tfRsrcTagsOut = append(*tfRsrcTagsOut, commonTerraform.Tag{
			TagName: types.StringValue(v.TagName),
			Color:   basetypes.NewStringPointerValue(color),
		})
	}
}

// build tag assign request using terraform resource as input
func buildApiReqTags(tfRsrcTags []commonTerraform.Tag) []commonApi.Tag {
	tags := []commonApi.Tag{}
	for _, tag := range tfRsrcTags {
		tags = append(tags, commonApi.Tag{
			Color:   tag.Color.ValueStringPointer(),
			TagName: tag.TagName.ValueString(),
		})
	}
	return tags
}

var ResourceBackupScheduleTime = resourceSchema.StringAttribute{
	MarkdownDescription: "Backup schedule time in 24 hour cron expression format.",
	Optional:            true,
	Computed:            true,
}

var resourceWal = resourceSchema.SingleNestedAttribute{
	Description: "Use a separate storage volume for Write-Ahead Logs (Recommended for high write workloads)",
	Optional:    true,
	Attributes: map[string]resourceSchema.Attribute{
		"iops": resourceSchema.StringAttribute{
			Description:   "IOPS for the selected volume. It can be set to different values depending on your volume type and properties.",
			Optional:      true,
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"size": resourceSchema.StringAttribute{
			Description:   "Size of the volume. It can be set to different values depending on your volume type and properties.",
			Required:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"throughput": resourceSchema.StringAttribute{
			Description:   "Throughput is automatically calculated by BigAnimal based on the IOPS input if it's not provided.",
			Optional:      true,
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"volume_properties": resourceSchema.StringAttribute{
			Description: "Volume properties in accordance with the selected volume type.",
			Required:    true,
		},
		"volume_type": resourceSchema.StringAttribute{
			Description: "Volume type. For Azure: \"azurepremiumstorage\" or \"ultradisk\". For AWS: \"gp3\", \"io2\", or \"io2-block-express\". For Google Cloud: only \"pd-ssd\".",
			Required:    true,
		},
	},
}

var DataSourceTagNestedObject = dataSourceSchema.NestedAttributeObject{
	Attributes: map[string]dataSourceSchema.Attribute{
		"tag_name": dataSourceSchema.StringAttribute{
			Computed: true,
		},
		"color": dataSourceSchema.StringAttribute{
			Computed: true,
		},
	},
}

var ResourceTagNestedObject = resourceSchema.NestedAttributeObject{
	PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
	Attributes: map[string]resourceSchema.Attribute{
		"tag_name": resourceSchema.StringAttribute{
			Required: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"color": resourceSchema.StringAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

func BuildTfRsrcCloudProviders(CloudProviders []models.CloudProvider) types.Set {
	cloudProviderAttrType := map[string]attr.Type{"cloud_provider_id": types.StringType, "cloud_provider_name": types.StringType}
	cloudProvidersValue := []attr.Value{}
	for _, provider := range CloudProviders {
		cloudProviderElem := basetypes.NewObjectValueMust(cloudProviderAttrType, map[string]attr.Value{
			"cloud_provider_id":   basetypes.NewStringValue(provider.CloudProviderId),
			"cloud_provider_name": basetypes.NewStringValue(provider.CloudProviderName),
		})
		cloudProvidersValue = append(cloudProvidersValue, cloudProviderElem)
	}

	return basetypes.NewSetValueMust(basetypes.ObjectType{AttrTypes: cloudProviderAttrType}, cloudProvidersValue)
}
