package plan_modifier

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func Test_customWitnessGroupDiffModifier_PlanModifySet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type args struct {
		ctx  context.Context
		req  planmodifier.SetRequest
		resp *planmodifier.SetResponse
	}

	test := map[string]attr.Type{"region_id": types.StringValue("").Type(ctx)}

	defaultAttrs := map[string]attr.Value{
		"region": basetypes.NewObjectValueMust(
			map[string]attr.Type{
				"region_id": basetypes.StringType{},
			},
			map[string]attr.Value{
				"region_id": basetypes.NewStringValue("us-east-1"),
			},
		),
		"cloud_provider": basetypes.NewStringValue(""),
	}
	// defaultAttrTypes := map[string]attr.Type{"cidr_block": defaultAttrs["cidr_block"].Type(ctx), "description": defaultAttrs["description"].Type(ctx)}

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
				resp: &planmodifier.SetResponse{
					PlanValue: basetypes.NewSetValueMust(),
				},
			},
			expectedWarningsCount: 0,
			expectedPlanElements: []attr.Value{
				basetypes.NewStringValue(""),
			},
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
