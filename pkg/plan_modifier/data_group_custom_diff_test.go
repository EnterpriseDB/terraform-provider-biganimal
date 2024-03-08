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
	// conditionsElemType := dgsSchemaAttr["conditions"].(schema.Attribute).GetType().(types.SetType).ElemType
	// dgsSchemaAttr := pgdSchema.Attributes["data_groups"].(schema.NestedAttribute).GetNestedObject().GetAttributes()
	// allowedIpRangesElemType := dgsSchemaAttr["allowed_ip_ranges"].(schema.Attribute).GetType().(types.SetType).ElemType
	// allowedIpRangesElemObjectType := dgsSchemaAttr["allowed_ip_ranges"].(schema.Attribute).GetType().(types.SetType).ElemType.(types.ObjectType).AttributeTypes()
	// clusterArchAttrType := dgsSchemaAttr["cluster_architecture"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	// dgElemAttrType := pgdSchema.Attributes["data_groups"].(schema.Attribute).GetType().(types.SetType).ElemType.(types.ObjectType).AttributeTypes()
	dgTfValue, _ := pgdSchema.Attributes["data_groups"].GetType().ValueType(ctx).ToTerraformValue(ctx)
	dgType := pgdSchema.Attributes["data_groups"].GetType()

	dgAttrTypes := dgType.(types.SetType).ElemType.(types.ObjectType).AttributeTypes()

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

			AllowedIpRanges: basetypes.NewSetValueMust(
				dgAttrTypes["allowed_ip_ranges"].(types.SetType).ElemType,
				[]attr.Value{
					types.ObjectValueMust(
						dgAttrTypes["allowed_ip_ranges"].(types.SetType).ElemType.(types.ObjectType).AttributeTypes(),
						map[string]attr.Value{
							"cidr_block":  types.StringPointerValue(utils.ToPointer("127.0.0.1/32")),
							"description": types.StringPointerValue(utils.ToPointer("test ip 1")),
						},
					),
					types.ObjectValueMust(
						dgAttrTypes["allowed_ip_ranges"].(types.SetType).ElemType.(types.ObjectType).AttributeTypes(),
						map[string]attr.Value{
							"cidr_block":  types.StringPointerValue(utils.ToPointer("192.0.0.1/32")),
							"description": types.StringPointerValue(utils.ToPointer("test ip 2")),
						},
					),
				},
			),

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
			ServiceAccountIds:     basetypes.NewSetUnknown(dgAttrTypes["service_account_ids"].(types.SetType).ElemType),
			PeAllowedPrincipalIds: basetypes.NewSetUnknown(dgAttrTypes["pe_allowed_principal_ids"].(types.SetType).ElemType),
			PgConfig:              &[]models.KeyValue{},
			Conditions:            basetypes.NewSetUnknown(dgAttrTypes["conditions"].(types.SetType).ElemType),
			ResizingPvc:           basetypes.NewSetUnknown(dgAttrTypes["resizing_pvc"].(types.SetType).ElemType),
		},
	}

	customState := tfsdk.State{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)}
	diag := customState.SetAttribute(ctx, path.Root("data_groups"), defaultDgs)
	if diag.ErrorsCount() > 0 {
		fmt.Printf("set attribute data groups error: %v", diag.Errors())
		return
	}

	tfDefaultDgs := new(types.Set)
	customState.GetAttribute(ctx, path.Root("data_groups"), tfDefaultDgs)

	addedDgs := []terraform.DataGroup(defaultDgs)
	newDg := defaultDgs[0]
	newDg.Region = &api.Region{RegionId: "us-east-2"}
	addDg := terraform.DataGroup(newDg)
	addedDgs = append(addedDgs, addDg)

	customState = tfsdk.State{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)}
	diag = customState.SetAttribute(ctx, path.Root("data_groups"), addedDgs)
	if diag.ErrorsCount() > 0 {
		fmt.Printf("set attribute data groups error %v", diag.Errors())
		return
	}

	tfAddedDgs := new(types.Set)
	customState.GetAttribute(ctx, path.Root("data_groups"), tfAddedDgs)

	updatedDgs := []terraform.DataGroup(defaultDgs)
	updatedDgs[0].AllowedIpRanges = basetypes.NewSetValueMust(
		dgAttrTypes["allowed_ip_ranges"].(types.SetType).ElemType,
		[]attr.Value{
			types.ObjectValueMust(
				dgAttrTypes["allowed_ip_ranges"].(types.SetType).ElemType.(types.ObjectType).AttributeTypes(),
				map[string]attr.Value{
					"cidr_block":  types.StringPointerValue(utils.ToPointer("168.0.0.1/32")),
					"description": types.StringPointerValue(utils.ToPointer("updated")),
				},
			),
		},
	)

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
		fmt.Printf("set attribute data groups error: %v", diag.Errors())
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
		{
			name: "Add dg expect success",
			args: args{
				ctx: ctx,
				req: planmodifier.SetRequest{
					Plan:       tfsdk.Plan{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)},
					StateValue: *tfDefaultDgs,
				},
				resp: &planmodifier.SetResponse{
					PlanValue: *tfAddedDgs,
				},
			},
			expectedWarningsCount:  1,
			expectedWarningSummary: []string{"Adding new data group"},
			expectedPlanElements:   tfAddedDgs.Elements(),
		},
		{
			name: "Remove dg expect success",
			args: args{
				ctx: ctx,
				req: planmodifier.SetRequest{
					Plan:       tfsdk.Plan{Schema: dgSchema, Raw: tftypes.NewValue(rawRootType.TerraformType(ctx), rawRootValue)},
					StateValue: *tfAddedDgs,
				},
				resp: &planmodifier.SetResponse{
					PlanValue: *tfDefaultDgs,
				},
			},
			expectedWarningsCount:  1,
			expectedWarningSummary: []string{"Removing data group"},
			expectedPlanElements:   tfDefaultDgs.Elements(),
		},
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
