package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type BackupSchedule struct {
	StartDay  types.String `tfsdk:"start_day"`
	StartTime types.String `tfsdk:"start_time"`
}
