package api

type BackupSchedule struct {
	StartDay  *float64 `json:"startDay,omitempty"`
	StartTime *string  `json:"startTime,omitempty"`
}
