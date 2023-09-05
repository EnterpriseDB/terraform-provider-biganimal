package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MaintenanceWindow struct {
	IsEnabled *bool `tfsdk:"is_enabled"`
	// The day of week, 0 represents Sunday, 1 is Monday, and so on.
	StartDay  types.Int64  `tfsdk:"start_day"`
	StartTime types.String `tfsdk:"start_time"`
}
