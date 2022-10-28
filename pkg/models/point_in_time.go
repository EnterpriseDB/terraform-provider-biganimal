package models

import "time"

type PointInTime struct {
	Nanos   float64 `json:"nanos"`
	Seconds float64 `json:"seconds"`
}

// PointInTimeToString transforms an apiv2 PointInTime
// to it's string representation
func (p PointInTime) String() string {
	return time.Unix(int64(p.Seconds), int64(p.Nanos)).String()
}
