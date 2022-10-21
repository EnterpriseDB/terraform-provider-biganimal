package utils

import (
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/apiv2"
)

func StringRef(s string) *string {
	return &s
}

func F64Ref(f float64) *float64 {
	return &f
}

// PointInTimeToString transforms an apiv2 PointInTime
// to it's string representation
func PointInTimeToString(p apiv2.PointInTime) string {
	return time.Unix(int64(p.Seconds), int64(p.Nanos)).String()
}
