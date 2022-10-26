package models

// this is probably just a kube condition, and we might
// be able to import that struct from the kube libs
type Condition struct {
	ConditionStatus    *string      `json:"conditionStatus,omitempty"`
	LastTransitionTime *PointInTime `json:"lastTransitionTime,omitempty"`
	Message            *string      `json:"message,omitempty"`
	Reason             *string      `json:"reason,omitempty"`
	Type_              *string      `json:"type,omitempty"`
}
