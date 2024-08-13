package plan_modifier_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func Test_customWitnessGroupDiffModifier_PlanModifySet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	defaultWgAttr := map[string]attr.Value{
		"region": basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"region_id": basetypes.StringType{},
			},
			map[string]attr.Value{
				"region_id": basetypes.NewStringValue("us-east-1"),
			},
		),
		"cloud_provider": basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"cloud_provider_id": basetypes.StringType{},
			},
			map[string]attr.Value{
				"cloud_provider_id": basetypes.NewStringValue("aws"),
			},
		),
	}

	defaultWgAttrTypes := map[string]attr.Type{
		"region":         defaultWgAttr["region"].Type(ctx),
		"cloud_provider": defaultWgAttr["cloud_provider"].Type(ctx),
	}

	defaultWgObject := basetypes.NewObjectValueMust(defaultWgAttrTypes, defaultWgAttr)
	defaultWgObjects := []attr.Value{}
	defaultWgObjects = append(defaultWgObjects, defaultWgObject)
	defaultWgSet := basetypes.NewListValueMust(defaultWgObject.Type(ctx), defaultWgObjects)

	addGroupObject := map[string]attr.Value{
		"region": basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"region_id": defaultWgAttr["region"].(basetypes.ObjectValue).Attributes()["region_id"].Type(ctx),
			},
			map[string]attr.Value{
				"region_id": basetypes.NewStringValue("us-east-2"),
			},
		),
		"cloud_provider": basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"cloud_provider_id": defaultWgAttr["cloud_provider"].(basetypes.ObjectValue).Attributes()["cloud_provider_id"].Type(ctx),
			},
			map[string]attr.Value{
				"cloud_provider_id": basetypes.NewStringValue("aws"),
			},
		),
	}

	type args struct {
		ctx  context.Context
		req  planmodifier.ListRequest
		resp *planmodifier.ListResponse
	}

	tests := []struct {
		name                   string
		m                      plan_modifier.CustomWitnessGroupDiffModifier
		args                   args
		expectedWarningsCount  int
		expectedWarningSummary []string
		expectedPlanElements   []attr.Value
	}{
		{
			name: "Add wg expected success",
			args: args{
				req: planmodifier.ListRequest{
					StateValue: defaultWgSet,
				},
				resp: &planmodifier.ListResponse{
					PlanValue: basetypes.NewListValueMust(defaultWgObject.Type(ctx),
						append(defaultWgObjects, basetypes.NewObjectValueMust(defaultWgAttrTypes,
							addGroupObject,
						)),
					),
				},
			},
			expectedWarningsCount:  1,
			expectedWarningSummary: []string{"Adding new witness group"},
			expectedPlanElements: append(defaultWgObjects, basetypes.NewObjectValueMust(defaultWgAttrTypes,
				addGroupObject,
			)),
		},
		{
			name: "Create new wg expected success",
			args: args{
				req: planmodifier.ListRequest{
					StateValue: basetypes.NewListNull(defaultWgObject.Type(ctx)),
				},
				resp: &planmodifier.ListResponse{
					PlanValue: basetypes.NewListValueMust(defaultWgObject.Type(ctx), defaultWgObjects),
				},
			},
			expectedPlanElements: defaultWgObjects,
		},
		{
			name: "Use state for unknown expected success",
			args: args{
				req: planmodifier.ListRequest{
					StateValue: basetypes.NewListValueMust(defaultWgObject.Type(ctx), defaultWgObjects),
					PlanValue:  basetypes.NewListUnknown(defaultWgObject.Type(ctx)),
				},
				resp: &planmodifier.ListResponse{
					PlanValue: basetypes.NewListUnknown(defaultWgObject.Type(ctx)),
				},
			},
			expectedPlanElements: defaultWgObjects,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.m.PlanModifyList(tt.args.ctx, tt.args.req, tt.args.resp)

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
