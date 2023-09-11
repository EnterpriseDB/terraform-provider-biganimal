package plan_modifier_test

import (
	"context"
	"testing"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/provider"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func Test_customPGConfigModifier_PlanModifySet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	pgdSchema := provider.PgdSchema(ctx)
	pgConfigElemType := pgdSchema.Attributes["data_groups"].(schema.NestedAttribute).GetNestedObject().GetAttributes()["pg_config"].(schema.Attribute).GetType().(types.SetType).ElemType
	pgConfigObjectType := pgdSchema.Attributes["data_groups"].(schema.NestedAttribute).GetNestedObject().GetAttributes()["pg_config"].(schema.Attribute).GetType().(types.SetType).ElemType.(types.ObjectType).AttributeTypes()

	defaultPgConfigDefaultsObjects := []attr.Value{}

	for k, v := range plan_modifier.PgConfigDefaults() {
		defaultPgConfigDefaultsObjects = append(defaultPgConfigDefaultsObjects,
			basetypes.NewObjectValueMust(pgConfigObjectType, map[string]attr.Value{
				"name":  basetypes.NewStringValue(k),
				"value": basetypes.NewStringValue(v),
			}),
		)
	}

	// defaultPgConfigSet := basetypes.NewSetValueMust(pgConfigElemType, defaultPgConfigDefaultsObjects)

	type args struct {
		ctx  context.Context
		req  planmodifier.SetRequest
		resp *planmodifier.SetResponse
	}
	tests := []struct {
		name                   string
		m                      plan_modifier.CustomPGConfigModifier
		args                   args
		expectedWarningsCount  int
		expectedWarningSummary []string
		expectedPlanElements   []attr.Value
	}{
		{
			name: "Use defaults only expected success",
			args: args{
				req: planmodifier.SetRequest{
					StateValue: basetypes.NewSetNull(pgConfigElemType),
				},
				resp: &planmodifier.SetResponse{
					PlanValue: basetypes.NewSetNull(pgConfigElemType),
				},
			},
			expectedWarningsCount:  1,
			expectedWarningSummary: []string{"PG config changed"},
			expectedPlanElements:   defaultPgConfigDefaultsObjects,
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

			expectedPlanSet := basetypes.NewSetValueMust(pgConfigElemType, tt.expectedPlanElements)
			respPlanValueSet := basetypes.NewSetValueMust(pgConfigElemType, tt.args.resp.PlanValue.Elements())

			if !expectedPlanSet.Equal(respPlanValueSet) {
				t.Fatalf("expected plan elements: %v, got: %v", expectedPlanSet, respPlanValueSet)
			}
		})
	}
}
