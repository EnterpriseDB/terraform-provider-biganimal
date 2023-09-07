package api

type MaintenanceWindow struct {
	IsEnabled *bool `json:"isEnabled,omitempty" mapstructure:"isEnabled"`
	// The day of week, 0 represents Sunday, 1 is Monday, and so on.
	StartDay  *float64 `json:"startDay,omitempty" mapstructure:"startDay"`
	StartTime *string  `json:"startTime,omitempty" mapstructure:"startTime"`
}
