package pgd

type Condition struct {
	ConditionStatus *string `json:"conditionStatus,omitempty" tfsdk:"condition_status"`
	Type_           *string `json:"type,omitempty" tfsdk:"type"`
}
