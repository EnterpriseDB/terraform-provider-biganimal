package provider

import (
	"context"
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// custom tag validation here
func ValidateTags(ctx context.Context, tagClient *api.TagClient, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
}

func buildRequestAllowedIpRanges(tfAllowedIpRanges basetypes.SetValue) *[]models.AllowedIpRange {
	apiAllowedIpRanges := &[]models.AllowedIpRange{}

	for _, v := range tfAllowedIpRanges.Elements() {
		*apiAllowedIpRanges = append(*apiAllowedIpRanges, models.AllowedIpRange{
			CidrBlock:   v.(types.Object).Attributes()["cidr_block"].(types.String).ValueString(),
			Description: v.(types.Object).Attributes()["description"].(types.String).ValueString(),
		})
	}

	return apiAllowedIpRanges
}

func buildTFRsrcAllowedIpRanges(respAllowedIpRanges *[]models.AllowedIpRange) (basetypes.SetValue, diag.Diagnostics) {
	attributeTypes := map[string]attr.Type{
		"cidr_block":  basetypes.StringType{},
		"description": basetypes.StringType{},
	}

	allowedIpRanges := []attr.Value{}
	if respAllowedIpRanges != nil && len(*respAllowedIpRanges) > 0 {
		for _, v := range *respAllowedIpRanges {
			v := v

			description := v.Description

			// if cidr block is 0.0.0.0/0 then set description to empty string
			// setting private networking = true and leaving allowed ip ranges as empty will return
			// cidr block as 0.0.0.0/0 and description as "To allow all access"
			// so we need to set description to empty string to keep it consistent with the tf resource
			if v.CidrBlock == "0.0.0.0/0" {
				description = ""
			}

			ob, diag := types.ObjectValue(attributeTypes, map[string]attr.Value{
				"cidr_block":  types.StringValue(v.CidrBlock),
				"description": types.StringValue(description),
			})
			if diag.HasError() {
				return basetypes.SetValue{}, diag
			}

			allowedIpRanges = append(allowedIpRanges, ob)
		}
	}

	allwdIpRngsElemType := types.ObjectType{AttrTypes: attributeTypes}
	allwdIpRngsSet := types.SetNull(allwdIpRngsElemType)
	if len(allowedIpRanges) > 0 {
		allwdIpRngsSet = types.SetValueMust(allwdIpRngsElemType, allowedIpRanges)
	}
	return allwdIpRngsSet, nil
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
	PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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

var resourceAllowedIpRanges = resourceSchema.SetNestedAttribute{
	Description: "Allowed IP ranges.",
	Optional:    true,
	Computed:    true,
	NestedObject: resourceSchema.NestedAttributeObject{
		Attributes: map[string]resourceSchema.Attribute{
			"cidr_block": resourceSchema.StringAttribute{
				Description: "CIDR block",
				Required:    true,
			},
			"description": resourceSchema.StringAttribute{
				Description: "Description of CIDR block",
				Required:    true,
			},
		},
	},
	PlanModifiers: []planmodifier.Set{plan_modifier.SetForceUnknownUpdate()},
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

func BuildTfRsrcWalStorage(walStorage *models.Storage) types.Object {
	walStorageAttrType := map[string]attr.Type{
		"iops":              types.StringType,
		"size":              types.StringType,
		"throughput":        types.StringType,
		"volume_properties": types.StringType,
		"volume_type":       types.StringType,
	}

	walStorageValue := map[string]attr.Value{
		"iops":              basetypes.NewStringNull(),
		"size":              basetypes.NewStringNull(),
		"throughput":        basetypes.NewStringNull(),
		"volume_properties": basetypes.NewStringNull(),
		"volume_type":       basetypes.NewStringNull(),
	}

	if walStorage != nil {
		walStorageValue = map[string]attr.Value{
			"iops":              basetypes.NewStringPointerValue(walStorage.Iops),
			"size":              basetypes.NewStringPointerValue(walStorage.Size),
			"throughput":        basetypes.NewStringPointerValue(walStorage.Throughput),
			"volume_properties": basetypes.NewStringPointerValue(walStorage.VolumePropertiesId),
			"volume_type":       basetypes.NewStringPointerValue(walStorage.VolumeTypeId),
		}
	}

	return basetypes.NewObjectValueMust(walStorageAttrType, walStorageValue)
}

func BuildRequestWalStorage(tfRsrcWalStorage types.Object) *models.Storage {
	if tfRsrcWalStorage.IsNull() || tfRsrcWalStorage.IsUnknown() {
		return nil
	}

	walStorage := &models.Storage{
		Iops:               tfRsrcWalStorage.Attributes()["iops"].(types.String).ValueStringPointer(),
		Size:               tfRsrcWalStorage.Attributes()["size"].(types.String).ValueStringPointer(),
		Throughput:         tfRsrcWalStorage.Attributes()["throughput"].(types.String).ValueStringPointer(),
		VolumePropertiesId: tfRsrcWalStorage.Attributes()["volume_properties"].(types.String).ValueStringPointer(),
		VolumeTypeId:       tfRsrcWalStorage.Attributes()["volume_type"].(types.String).ValueStringPointer(),
	}

	return walStorage
}

// json marshals to bytes then unmarshals into struct to fill a struct. Out has to be a pointer
func CopyObjectJson(in any, out any) error {
	bytes, err := json.Marshal(in)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return err
	}

	return nil
}
