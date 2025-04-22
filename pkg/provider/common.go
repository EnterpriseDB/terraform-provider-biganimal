package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// custom tag validaiton here
func ValidateTags(ctx context.Context, tagClient *api.TagClient, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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
