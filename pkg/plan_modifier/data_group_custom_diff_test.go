package plan_modifier_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/provider"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func Test_customDataGroupDiffModifier_PlanModifySet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	pgdSchema := provider.PgdSchema(ctx)

	dgSchema := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"data_groups": pgdSchema.Attributes["data_groups"].(schema.NestedAttribute),
		},
	}

	// dgsSchemaAttr := pgdSchema.Attributes["data_groups"].(schema.NestedAttribute).GetNestedObject().GetAttributes()
	// allowedIpRangesElemType := dgsSchemaAttr["allowed_ip_ranges"].(schema.Attribute).GetType().(types.SetType).ElemType
	// allowedIpRangesElemObjectType := dgsSchemaAttr["allowed_ip_ranges"].(schema.Attribute).GetType().(types.SetType).ElemType.(types.ObjectType).AttributeTypes()
	// clusterArchAttrType := dgsSchemaAttr["cluster_architecture"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	// dgElemAttrType := pgdSchema.Attributes["data_groups"].(schema.Attribute).GetType().(types.SetType).ElemType.(types.ObjectType).AttributeTypes()
	dgTfValue, _ := pgdSchema.Attributes["data_groups"].GetType().ValueType(ctx).ToTerraformValue(ctx)
	dgType := pgdSchema.Attributes["data_groups"].GetType()

	rawRootValue := map[string]tftypes.Value{
		"data_groups": dgTfValue,
	}

	rawRootType := basetypes.ObjectType{AttrTypes: map[string]attr.Type{
		"data_groups": dgType,
	}}

	defaultDgs := []terraform.DataGroup{
		{
			Region:   &api.Region{RegionId: "us-east-1"},
			Provider: &api.CloudProvider{CloudProviderId: utils.ToPointer("aws")},
			Storage: &terraform.Storage{
				VolumeTypeId:       basetypes.NewStringValue("gp3"),
				VolumePropertiesId: basetypes.NewStringValue("gp3"),
				Size:               basetypes.NewStringValue("4 Gi"),
				Iops:               basetypes.NewStringUnknown(),
				Throughput:         basetypes.NewStringUnknown(),
			},
			AllowedIpRanges: &[]models.AllowedIpRange{
				{CidrBlock: "127.0.0.1/32", Description: "test ip 1"},
				{CidrBlock: "192.0.0.1/32", Description: "test ip 2"},
			},
			ClusterArchitecture: &terraform.ClusterArchitecture{
				ClusterArchitectureId:   "pgd",
				ClusterArchitectureName: basetypes.NewStringUnknown(),
				Nodes:                   3,
				WitnessNodes:            basetypes.NewInt64Unknown(),
			},
			InstanceType: &api.InstanceType{InstanceTypeId: "aws:m5.large"},
			PgType:       &api.PgType{PgTypeId: "epas"},
			PgVersion:    &api.PgVersion{PgVersionId: "15"},
			MaintenanceWindow: &models.MaintenanceWindow{
				IsEnabled: utils.ToPointer(true),
				StartDay:  utils.ToPointer(float64(1)),
				StartTime: utils.ToPointer("03:00"),
			},
			BackupRetentionPeriod: utils.ToPointer("3d"),
			ServiceAccountIds:     basetypes.SetValue{},
			PeAllowedPrincipalIds: basetypes.SetValue{},
			PgConfig:              &[]models.KeyValue{},
		},
	}

	customState := tfsdk.State{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)}
	diag := customState.SetAttribute(ctx, path.Root("data_groups"), defaultDgs)
	if diag.ErrorsCount() > 0 {
		_ = fmt.Errorf("set attribute data groups error")
		return
	}

	tfDefaultDgs := new(types.Set)
	customState.GetAttribute(ctx, path.Root("data_groups"), tfDefaultDgs)

	// addGroupObject := map[string]attr.Value{
	// 	"region": basetypes.NewObjectValueMust(regionType,
	// 		map[string]attr.Value{
	// 			"region_id": basetypes.NewStringValue("us-east-2"),
	// 		},
	// 	),
	// 	"cloud_provider":          basetypes.NewObjectValueMust(cloudProviderType, defaultCloudProvider),
	// 	"storage":                 basetypes.NewObjectValueMust(storageAttrType, defaultStorage),
	// 	"cluster_name":            basetypes.NewStringUnknown(),
	// 	"cluster_type":            basetypes.NewStringUnknown(),
	// 	"conditions":              basetypes.NewSetUnknown(conditionsElemType),
	// 	"connection_uri":          basetypes.NewStringUnknown(),
	// 	"created_at":              basetypes.NewStringUnknown(),
	// 	"group_id":                basetypes.NewStringUnknown(),
	// 	"logs_url":                basetypes.NewStringUnknown(),
	// 	"metrics_url":             basetypes.NewStringUnknown(),
	// 	"phase":                   basetypes.NewStringUnknown(),
	// 	"resizing_pvc":            basetypes.NewSetUnknown(resizingPvcElemType),
	// 	"allowed_ip_ranges":       basetypes.NewSetValueMust(allowedIpRangesElemType, defaultAllowedIpRange),
	// 	"backup_retention_period": basetypes.NewStringValue(defaultBackupRetentionPeriod),
	// 	"cluster_architecture":    basetypes.NewObjectValueMust(clusterArchAttrType, defaultClusterArch),
	// 	"csp_auth":                basetypes.NewBoolValue(false),
	// 	"instance_type":           basetypes.NewObjectValueMust(instanceTypeType, defaultInstanceType),
	// 	"pg_type":                 basetypes.NewObjectValueMust(pgTypeType, defaultPgType),
	// 	"pg_version":              basetypes.NewObjectValueMust(pgVersionType, defaultPgVersion),
	// 	"private_networking":      basetypes.NewBoolValue(false),
	// 	"maintenance_window": basetypes.NewObjectValueMust(cmwType,
	// 		map[string]attr.Value{
	// 			"is_enabled": basetypes.NewBoolValue(true),
	// 			"start_day":  basetypes.NewFloat64Value(2),
	// 			"start_time": basetypes.NewStringValue("06:00"),
	// 		},
	// 	),
	// 	"service_account_ids":      basetypes.NewSetValueMust(saElemType, []attr.Value{}),
	// 	"pe_allowed_principal_ids": basetypes.NewSetValueMust(peElemType, []attr.Value{}),
	// 	"pg_config":                basetypes.NewSetValueMust(pgElemType, []attr.Value{}),
	// }

	updatedDgs := []terraform.DataGroup(defaultDgs)
	updatedDgs[0].AllowedIpRanges = &[]models.AllowedIpRange{
		{
			CidrBlock:   "168.0.0.1/32",
			Description: "updated",
		},
	}
	updatedDgs[0].BackupRetentionPeriod = utils.ToPointer("5d")
	updatedDgs[0].ClusterArchitecture = &terraform.ClusterArchitecture{
		ClusterArchitectureId:   "pgd",
		ClusterArchitectureName: basetypes.NewStringUnknown(),
		Nodes:                   1,
		WitnessNodes:            basetypes.NewInt64Unknown(),
	}

	customState = tfsdk.State{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)}
	diag = customState.SetAttribute(ctx, path.Root("data_groups"), updatedDgs)
	if diag.ErrorsCount() > 0 {
		_ = fmt.Errorf("set attribute data groups error")
		return
	}

	tfUpdatedDgs := new(types.Set)
	customState.GetAttribute(ctx, path.Root("data_groups"), tfUpdatedDgs)

	type args struct {
		ctx  context.Context
		req  planmodifier.SetRequest
		resp *planmodifier.SetResponse
	}
	tests := []struct {
		name                   string
		m                      plan_modifier.CustomDataGroupDiffModifier
		args                   args
		expectedWarningsCount  int
		expectedWarningSummary []string
		expectedPlanElements   []attr.Value
	}{
		// {
		// 	name: "Add dg expect success",
		// 	args: args{
		// 		ctx: ctx,
		// 		req: planmodifier.SetRequest{
		// 			Plan:       tfsdk.Plan{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)},
		// 			StateValue: defaultDgSet,
		// 		},
		// 		resp: &planmodifier.SetResponse{
		// 			PlanValue: basetypes.NewSetValueMust(defaultDgObject.Type(ctx),
		// 				append(defaultDgObjects, basetypes.NewObjectValueMust(dgElemAttrType, addGroupObject)),
		// 			),
		// 		},
		// 	},
		// 	expectedWarningsCount:  1,
		// 	expectedWarningSummary: []string{"Adding new data group"},
		// 	expectedPlanElements:   append(defaultDgObjects, basetypes.NewObjectValueMust(dgElemAttrType, addGroupObject)),
		// },
		// {
		// 	name: "Remove dg expect success",
		// 	args: args{
		// 		ctx: ctx,
		// 		req: planmodifier.SetRequest{
		// 			Plan: tfsdk.Plan{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)},
		// 			StateValue: basetypes.NewSetValueMust(defaultDgObject.Type(ctx),
		// 				append(defaultDgObjects, basetypes.NewObjectValueMust(dgElemAttrType, addGroupObject)),
		// 			),
		// 		},
		// 		resp: &planmodifier.SetResponse{
		// 			PlanValue: defaultDgSet,
		// 		},
		// 	},
		// 	expectedWarningsCount:  1,
		// 	expectedWarningSummary: []string{"Removing data group"},
		// 	expectedPlanElements:   defaultDgObjects,
		// },
		{
			name: "Update object expect success",
			args: args{
				ctx: ctx,
				req: planmodifier.SetRequest{
					Plan:       tfsdk.Plan{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)},
					StateValue: *tfDefaultDgs,
				},
				resp: &planmodifier.SetResponse{
					PlanValue: *tfUpdatedDgs,
				},
			},
			expectedWarningsCount: 3,
			expectedWarningSummary: []string{
				"Allowed IP ranges changed",
				"Backup retention changed",
				"Cluster architecture changed",
			},
			expectedPlanElements: tfUpdatedDgs.Elements(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.m.PlanModifySet(tt.args.ctx, tt.args.req, tt.args.resp)

			if tt.args.resp.Diagnostics.WarningsCount() != tt.expectedWarningsCount {
				t.Fatalf("expected warning count: %v, got: %v", tt.expectedWarningsCount, tt.args.resp.Diagnostics.WarningsCount())
			}

			if tt.args.resp.Diagnostics.WarningsCount() != 0 {
				for k, v := range tt.args.resp.Diagnostics.Warnings() {
					if tt.expectedWarningSummary[k] != v.Summary() {
						t.Fatalf("expected warning summary: %v, got: %v", tt.expectedWarningSummary[k], v.Summary())
					}
				}
			}

			if !reflect.DeepEqual(tt.expectedPlanElements, tt.args.resp.PlanValue.Elements()) {
				t.Fatalf("expected plan elements: %v, got: %v", tt.expectedPlanElements, tt.args.resp.PlanValue.Elements())
			}
		})
	}
}
