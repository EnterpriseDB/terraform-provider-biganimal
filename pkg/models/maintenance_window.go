package models

type MaintenanceWindow struct {
	IsEnabled *bool `json:"isEnabled,omitempty" tfsdk:"is_enabled"`
	// The day of week, 0 represents Sunday, 1 is Monday, and so on.
	StartDay  *float64 `json:"startDay,omitempty" tfsdk:"start_day"`
	StartTime *string  `json:"startTime,omitempty" tfsdk:"start_time"`
}
