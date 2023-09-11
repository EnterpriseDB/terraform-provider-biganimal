package plan_modifier

import (
	"context"
	"reflect"
	"testing"

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
	defaultWgSet := basetypes.NewSetValueMust(defaultWgObject.Type(ctx), defaultWgObjects)

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
		req  planmodifier.SetRequest
		resp *planmodifier.SetResponse
	}

	tests := []struct {
		name                   string
		m                      customWitnessGroupDiffModifier
		args                   args
		expectedWarningsCount  int
		expectedWarningSummary []string
		expectedPlanElements   []attr.Value
	}{
		{
			name: "Add wg success",
			args: args{
				req: planmodifier.SetRequest{
					StateValue: defaultWgSet,
				},
				resp: &planmodifier.SetResponse{
					PlanValue: basetypes.NewSetValueMust(defaultWgObject.Type(ctx),
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
			name: "create new wg success",
			args: args{
				req: planmodifier.SetRequest{
					StateValue: basetypes.NewSetNull(defaultWgObject.Type(ctx)),
				},
				resp: &planmodifier.SetResponse{
					PlanValue: basetypes.NewSetValueMust(defaultWgObject.Type(ctx), defaultWgObjects),
				},
			},
			expectedPlanElements: defaultWgObjects,
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
