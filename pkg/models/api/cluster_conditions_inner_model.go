package api

type ClusterConditionsInner struct {
	ConditionStatus    *string      `json:"conditionStatus,omitempty"`
	LastTransitionTime *PointInTime `json:"lastTransitionTime,omitempty"`
	Message            *string      `json:"message,omitempty"`
	Reason             *string      `json:"reason,omitempty"`
	Type_              *string      `json:"type,omitempty"`
}
