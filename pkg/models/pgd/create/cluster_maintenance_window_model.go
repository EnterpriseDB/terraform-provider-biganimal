package create

type ClusterMaintenanceWindow struct {
	IsEnabled bool `json:"isEnabled"`
	// The day of week, 0 represents Sunday, 1 is Monday, and so on.
	StartDay  *float64 `json:"startDay,omitempty"`
	StartTime *string  `json:"startTime,omitempty"`
}
