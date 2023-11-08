package plan_modifier_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/provider"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

	// clusterArchPathSteps := tftypes.NewAttributePathWithSteps([]tftypes.AttributePathStep{
	// 	tftypes.AttributeName("cluster_architecture"),
	// })

	// pp, _ := pgdSchema.AttributeAtPath(ctx, path.Root("data_groups"))

	// dgAttrType := pgdSchema.Attributes["data_groups"].(schema.NestedAttribute).GetNestedObject().GetAttributes().Type().ValueType(ctx).Type(ctx).(types.ObjectType).AttributeTypes()

	dgsSchemaAttr := pgdSchema.Attributes["data_groups"].(schema.NestedAttribute).GetNestedObject().GetAttributes()
	regionType := dgsSchemaAttr["region"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	cloudProviderType := dgsSchemaAttr["cloud_provider"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	storageAttrType := dgsSchemaAttr["storage"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	conditionsElemType := dgsSchemaAttr["conditions"].(schema.Attribute).GetType().(types.SetType).ElemType
	resizingPvcElemType := dgsSchemaAttr["resizing_pvc"].(schema.Attribute).GetType().(types.SetType).ElemType
	allowedIpRangesElemType := dgsSchemaAttr["allowed_ip_ranges"].(schema.Attribute).GetType().(types.SetType).ElemType
	allowedIpRangesElemObjectType := dgsSchemaAttr["allowed_ip_ranges"].(schema.Attribute).GetType().(types.SetType).ElemType.(types.ObjectType).AttributeTypes()
	clusterArchAttrType := dgsSchemaAttr["cluster_architecture"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	instanceTypeType := dgsSchemaAttr["instance_type"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	pgTypeType := dgsSchemaAttr["pg_type"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	pgVersionType := dgsSchemaAttr["pg_version"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	cmwType := dgsSchemaAttr["maintenance_window"].(schema.Attribute).GetType().(types.ObjectType).AttributeTypes()
	saElemType := dgsSchemaAttr["service_account_ids"].(schema.Attribute).GetType().(types.SetType).ElemType
	peElemType := dgsSchemaAttr["pe_allowed_principal_ids"].(schema.Attribute).GetType().(types.SetType).ElemType
	pgElemType := dgsSchemaAttr["pg_config"].(schema.Attribute).GetType().(types.SetType).ElemType

	defaultRegion := map[string]attr.Value{
		"region_id": basetypes.NewStringValue("us-east-1"),
	}

	defaultCloudProvider := map[string]attr.Value{
		"cloud_provider_id": basetypes.NewStringValue("aws"),
	}

	defaultStorage := map[string]attr.Value{
		"volume_type":       basetypes.NewStringValue("gp3"),
		"volume_properties": basetypes.NewStringValue("gp3"),
		"size":              basetypes.NewStringValue("4 Gi"),
		"iops":              basetypes.NewStringUnknown(),
		"throughput":        basetypes.NewStringUnknown(),
	}

	defaultAllowedIpRange := []attr.Value{
		basetypes.NewObjectValueMust(allowedIpRangesElemObjectType, map[string]attr.Value{
			"cidr_block":  basetypes.NewStringValue("127.0.0.1/32"),
			"description": basetypes.NewStringValue("test ip 1"),
		}),
		basetypes.NewObjectValueMust(allowedIpRangesElemObjectType, map[string]attr.Value{
			"cidr_block":  basetypes.NewStringValue("192.0.0.1/32"),
			"description": basetypes.NewStringValue("test ip 2"),
		}),
	}

	defaultClusterArch := map[string]attr.Value{
		"cluster_architecture_id":   basetypes.NewStringValue("pgd"),
		"cluster_architecture_name": basetypes.NewStringUnknown(),
		"nodes":                     basetypes.NewFloat64Value(3),
		"witness_nodes":             basetypes.NewInt64Unknown(),
	}

	defaultInstanceType := map[string]attr.Value{
		"instance_type_id": basetypes.NewStringValue("aws:m5.large"),
	}

	defaultPgType := map[string]attr.Value{
		"pg_type_id": basetypes.NewStringValue("epas"),
	}

	defaultPgVersion := map[string]attr.Value{
		"pg_version_id": basetypes.NewStringValue("15"),
	}

	defaultCmw := map[string]attr.Value{
		"is_enabled": basetypes.NewBoolValue(true),
		"start_day":  basetypes.NewFloat64Value(1),
		"start_time": basetypes.NewStringValue("03:00"),
	}

	defaultSa := []attr.Value{}

	defaultPe := []attr.Value{}

	defaultPg := []attr.Value{}

	defaultBackupRetentionPeriod := "3d"

	defaultDgAttr := map[string]attr.Value{
		"region":                   basetypes.NewObjectValueMust(regionType, defaultRegion),
		"cloud_provider":           basetypes.NewObjectValueMust(cloudProviderType, defaultCloudProvider),
		"storage":                  basetypes.NewObjectValueMust(storageAttrType, defaultStorage),
		"cluster_name":             basetypes.NewStringUnknown(),
		"cluster_type":             basetypes.NewStringUnknown(),
		"conditions":               basetypes.NewSetUnknown(conditionsElemType),
		"connection_uri":           basetypes.NewStringUnknown(),
		"created_at":               basetypes.NewStringUnknown(),
		"group_id":                 basetypes.NewStringUnknown(),
		"logs_url":                 basetypes.NewStringUnknown(),
		"metrics_url":              basetypes.NewStringUnknown(),
		"phase":                    basetypes.NewStringUnknown(),
		"resizing_pvc":             basetypes.NewSetUnknown(resizingPvcElemType),
		"allowed_ip_ranges":        basetypes.NewSetValueMust(allowedIpRangesElemType, defaultAllowedIpRange),
		"backup_retention_period":  basetypes.NewStringValue(defaultBackupRetentionPeriod),
		"cluster_architecture":     basetypes.NewObjectValueMust(clusterArchAttrType, defaultClusterArch),
		"csp_auth":                 basetypes.NewBoolValue(false),
		"instance_type":            basetypes.NewObjectValueMust(instanceTypeType, defaultInstanceType),
		"pg_type":                  basetypes.NewObjectValueMust(pgTypeType, defaultPgType),
		"pg_version":               basetypes.NewObjectValueMust(pgVersionType, defaultPgVersion),
		"private_networking":       basetypes.NewBoolValue(false),
		"maintenance_window":       basetypes.NewObjectValueMust(cmwType, defaultCmw),
		"service_account_ids":      basetypes.NewSetValueMust(saElemType, defaultSa),
		"pe_allowed_principal_ids": basetypes.NewSetValueMust(peElemType, defaultPe),
		"pg_config":                basetypes.NewSetValueMust(pgElemType, defaultPg),
	}

	defaultDgAttrTypes := map[string]attr.Type{
		"region":                   defaultDgAttr["region"].Type(ctx),
		"cloud_provider":           defaultDgAttr["cloud_provider"].Type(ctx),
		"storage":                  defaultDgAttr["storage"].Type(ctx),
		"cluster_name":             defaultDgAttr["cluster_name"].Type(ctx),
		"cluster_type":             defaultDgAttr["cluster_type"].Type(ctx),
		"conditions":               defaultDgAttr["conditions"].Type(ctx),
		"connection_uri":           defaultDgAttr["connection_uri"].Type(ctx),
		"created_at":               defaultDgAttr["created_at"].Type(ctx),
		"group_id":                 defaultDgAttr["group_id"].Type(ctx),
		"logs_url":                 defaultDgAttr["logs_url"].Type(ctx),
		"metrics_url":              defaultDgAttr["metrics_url"].Type(ctx),
		"phase":                    defaultDgAttr["phase"].Type(ctx),
		"resizing_pvc":             defaultDgAttr["resizing_pvc"].Type(ctx),
		"allowed_ip_ranges":        defaultDgAttr["allowed_ip_ranges"].Type(ctx),
		"backup_retention_period":  defaultDgAttr["backup_retention_period"].Type(ctx),
		"cluster_architecture":     defaultDgAttr["cluster_architecture"].Type(ctx),
		"csp_auth":                 defaultDgAttr["csp_auth"].Type(ctx),
		"instance_type":            defaultDgAttr["instance_type"].Type(ctx),
		"pg_type":                  defaultDgAttr["pg_type"].Type(ctx),
		"pg_version":               defaultDgAttr["pg_version"].Type(ctx),
		"private_networking":       defaultDgAttr["private_networking"].Type(ctx),
		"maintenance_window":       defaultDgAttr["maintenance_window"].Type(ctx),
		"service_account_ids":      defaultDgAttr["service_account_ids"].Type(ctx),
		"pe_allowed_principal_ids": defaultDgAttr["pe_allowed_principal_ids"].Type(ctx),
		"pg_config":                defaultDgAttr["pg_config"].Type(ctx),
	}

	defaultDgObject := basetypes.NewObjectValueMust(defaultDgAttrTypes, defaultDgAttr)
	defaultDgObjects := []attr.Value{}
	defaultDgObjects = append(defaultDgObjects, defaultDgObject)
	defaultDgSet := basetypes.NewSetValueMust(defaultDgObject.Type(ctx), defaultDgObjects)

	addGroupObject := map[string]attr.Value{
		"region": basetypes.NewObjectValueMust(regionType,
			map[string]attr.Value{
				"region_id": basetypes.NewStringValue("us-east-2"),
			},
		),
		"cloud_provider":          basetypes.NewObjectValueMust(cloudProviderType, defaultCloudProvider),
		"storage":                 basetypes.NewObjectValueMust(storageAttrType, defaultStorage),
		"cluster_name":            basetypes.NewStringUnknown(),
		"cluster_type":            basetypes.NewStringUnknown(),
		"conditions":              basetypes.NewSetUnknown(conditionsElemType),
		"connection_uri":          basetypes.NewStringUnknown(),
		"created_at":              basetypes.NewStringUnknown(),
		"group_id":                basetypes.NewStringUnknown(),
		"logs_url":                basetypes.NewStringUnknown(),
		"metrics_url":             basetypes.NewStringUnknown(),
		"phase":                   basetypes.NewStringUnknown(),
		"resizing_pvc":            basetypes.NewSetUnknown(resizingPvcElemType),
		"allowed_ip_ranges":       basetypes.NewSetValueMust(allowedIpRangesElemType, defaultAllowedIpRange),
		"backup_retention_period": basetypes.NewStringValue(defaultBackupRetentionPeriod),
		"cluster_architecture":    basetypes.NewObjectValueMust(clusterArchAttrType, defaultClusterArch),
		"csp_auth":                basetypes.NewBoolValue(false),
		"instance_type":           basetypes.NewObjectValueMust(instanceTypeType, defaultInstanceType),
		"pg_type":                 basetypes.NewObjectValueMust(pgTypeType, defaultPgType),
		"pg_version":              basetypes.NewObjectValueMust(pgVersionType, defaultPgVersion),
		"private_networking":      basetypes.NewBoolValue(false),
		"maintenance_window": basetypes.NewObjectValueMust(cmwType,
			map[string]attr.Value{
				"is_enabled": basetypes.NewBoolValue(true),
				"start_day":  basetypes.NewFloat64Value(2),
				"start_time": basetypes.NewStringValue("06:00"),
			},
		),
		"service_account_ids":      basetypes.NewSetValueMust(saElemType, defaultSa),
		"pe_allowed_principal_ids": basetypes.NewSetValueMust(peElemType, defaultPe),
		"pg_config":                basetypes.NewSetValueMust(pgElemType, defaultPg),
	}

	_ = addGroupObject

	// ob := basetypes.NewObjectValueMust(map[string]attr.Type{
	// 	"data_groups": basetypes.ObjectType{AttrTypes: defaultDgAttrTypes},
	// },
	// 	map[string]attr.Value{
	// 		"data_groups": basetypes.NewObjectValueMust(defaultDgAttrTypes, defaultDgAttr),
	// 	},
	// )

	// old
	// dgt := basetypes.ObjectType{AttrTypes: defaultDgAttrTypes}

	// // dgv := basetypes.NewObjectValueMust(defaultDgAttrTypes, defaultDgAttr)

	// // ddd := basetypes.NewObjectValueMust(dgt, dgv)

	// tt := basetypes.NewSetValueMust(dgt, defaultDgObjects)

	// xx, _ := basetypes.NewObjectValueMust(defaultDgAttrTypes, defaultDgAttr).ToTerraformValue(ctx)
	// vv := []tftypes.Value{}
	// vv = append(vv, xx)

	// old end

	bb := basetypes.NewObjectValueMust(defaultDgAttrTypes, defaultDgAttr)
	// xx, _ := bb.ToTerraformValue(ctx)

	gg := []attr.Value{}
	gg = append(gg, bb)
	zz := basetypes.NewSetValueMust(bb.Type(ctx), gg)
	yy, _ := zz.ToTerraformValue(ctx)
	vv := map[string]tftypes.Value{
		"data_groups": yy,
	}

	rootOb := basetypes.NewObjectValueMust(map[string]attr.Type{
		"data_groups": zz.Type(ctx),
	},
		map[string]attr.Value{
			"data_groups": zz,
		},
	)

	// sss := schema.Schema{
	// 	Attributes: map[string]schema.Attribute{
	// 		"data_groups": schema.SetNestedAttribute{NestedObject: schema.NestedAttributeObject{
	// 			Attributes: pgdSchema.Attributes["data_groups"].(schema.NestedAttribute),
	// 		}},
	// 	},
	// }

	sss := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"data_groups": pgdSchema.Attributes["data_groups"].(schema.NestedAttribute),
		},
	}

	// rootOb := basetypes.NewObjectValueMust(map[string]attr.Type{
	// 	"data_groups": basetypes.ObjectType{AttrTypes: defaultDgAttrTypes},
	// },
	// 	map[string]attr.Value{
	// 		"data_groups": basetypes.NewObjectValueMust(defaultDgAttrTypes, defaultDgAttr),
	// 	},
	// )

	// sss := schema.Schema{
	// 	Attributes: map[string]schema.Attribute{
	// 		"data_groups": schema.ObjectAttribute{AttributeTypes: defaultDgAttrTypes},
	// 	},
	// }

	// vv := map[string]tftypes.Value{
	// 	"data_groups": tftypes.NewValue(tftypes.Object{
	// 		AttributeTypes: map[string]tftypes.Type{
	// 			"cluster_name": tftypes.String,
	// 		},
	// 	},
	// 		map[string]tftypes.Value{
	// 			"cluster_name": tftypes.NewValue(tftypes.String, "hello"),
	// 		}),
	// }

	// objType := tftypes.Object{
	// 	AttributeTypes: map[string]tftypes.Type{
	// 		"cluster_name": tftypes.String,
	// 	},
	// }

	updateObjectAttr := map[string]attr.Value{}

	for k, v := range defaultDgAttr {
		updateObjectAttr[k] = v
	}

	updateObjectAttr["allowed_ip_ranges"] = basetypes.NewSetValueMust(allowedIpRangesElemType, []attr.Value{
		basetypes.NewObjectValueMust(allowedIpRangesElemObjectType, map[string]attr.Value{
			"cidr_block":  basetypes.NewStringValue("168.0.0.1/32"),
			"description": basetypes.NewStringValue("updated"),
		}),
	})
	updateObjectAttr["backup_retention_period"] = basetypes.NewStringValue("5d")
	updateObjectAttr["cluster_architecture"] = basetypes.NewObjectValueMust(clusterArchAttrType, map[string]attr.Value{
		"cluster_architecture_id":   basetypes.NewStringValue("pgd"),
		"cluster_architecture_name": basetypes.NewStringUnknown(),
		"nodes":                     basetypes.NewFloat64Value(1),
		"witness_nodes":             basetypes.NewInt64Unknown(),
	})

	updateObject := basetypes.NewObjectValueMust(defaultDgAttrTypes, updateObjectAttr)
	updateObjects := []attr.Value{}
	updateObjects = append(updateObjects, updateObject)
	updateSet := basetypes.NewSetValueMust(defaultDgObject.Type(ctx), updateObjects)

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
					Plan:       tfsdk.Plan{Schema: sss, Raw: tftypes.NewValue(rootOb.Type(ctx).TerraformType(ctx), vv)},
					StateValue: defaultDgSet,
				},
				resp: &planmodifier.SetResponse{
					PlanValue: basetypes.NewSetValueMust(defaultDgObject.Type(ctx),
						append(defaultDgObjects, basetypes.NewObjectValueMust(defaultDgAttrTypes, addGroupObject)),
					),
				},
			},
			expectedWarningsCount:  1,
			expectedWarningSummary: []string{"Adding new data group"},
			expectedPlanElements:   append(defaultDgObjects, basetypes.NewObjectValueMust(defaultDgAttrTypes, addGroupObject)),
		},
		{
			name: "Remove dg expect success",
			args: args{
				ctx: ctx,
				req: planmodifier.SetRequest{
					Plan: tfsdk.Plan{Schema: sss, Raw: tftypes.NewValue(rootOb.Type(ctx).TerraformType(ctx), vv)},
					StateValue: basetypes.NewSetValueMust(defaultDgObject.Type(ctx),
						append(defaultDgObjects, basetypes.NewObjectValueMust(defaultDgAttrTypes, addGroupObject)),
					),
				},
				resp: &planmodifier.SetResponse{
					PlanValue: defaultDgSet,
				},
			},
			expectedWarningsCount:  1,
			expectedWarningSummary: []string{"Removing data group"},
			expectedPlanElements:   defaultDgObjects,
		},
		{
			name: "Update object expect success",
			args: args{
				ctx: ctx,
				req: planmodifier.SetRequest{
					Plan:       tfsdk.Plan{Schema: sss, Raw: tftypes.NewValue(rootOb.Type(ctx).TerraformType(ctx), vv)},
					StateValue: defaultDgSet,
				},
				resp: &planmodifier.SetResponse{
					PlanValue: updateSet,
				},
			},
			expectedWarningsCount: 3,
			expectedWarningSummary: []string{
				"Allowed IP ranges changed",
				"Backup retention changed",
				"Cluster architecture changed",
			},
			expectedPlanElements: updateObjects,
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
