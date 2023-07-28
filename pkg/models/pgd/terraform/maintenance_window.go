package terraform

type MaintenanceWindow struct {
	IsEnabled *bool `tfsdk:"is_enabled"`
	// The day of week, 0 represents Sunday, 1 is Monday, and so on.
	StartDay  *float64 `tfsdk:"start_day"`
	StartTime *string  `tfsdk:"start_time"`
}
