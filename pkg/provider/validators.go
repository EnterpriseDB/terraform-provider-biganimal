package provider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/google/uuid"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// Please use the ProjectIdValidator validator.String when you migrate any SDKv2 resource/data-source to the Framework Library.
func validateProjectId(v interface{}, path cty.Path) diag.Diagnostics {
	value := v.(string)
	var diags diag.Diagnostics
	// if value != can(regex("^prj_[[:alnum:]]{16}$", value)) {
	matched, _ := regexp.MatchString("^prj_[0-9A-Za-z_]{16}$", value)
	if !matched {
		diag := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "invalid value for variable",
			Detail:   fmt.Sprintf("%q is not valid. Please provide a valid name for the project_id, for example: prj_abcdABCD01234567.", value),
		}
		diags = append(diags, diag)
	}
	return diags
}

func validateARN(v interface{}, _ cty.Path) diag.Diagnostics {
	a, err := arn.Parse(v.(string))
	if err != nil || a.Service != "iam" || !strings.HasPrefix(a.Resource, "role") {
		return diag.Diagnostics{{
			Severity: diag.Error,
			Summary:  "invalid arn",
			Detail:   fmt.Sprintf("%v is a invalid aws role arn", v),
		}}
	}

	return nil
}

func validateUUID(v interface{}, _ cty.Path) diag.Diagnostics {
	_, err := uuid.Parse(v.(string))
	if err != nil {
		return diag.Diagnostics{{
			Severity: diag.Error,
			Summary:  "invalid uuid",
			Detail:   fmt.Sprintf("%v is a invalid uuid", v),
		}}
	}
	return nil
}

//////////////////////////////////
// Framework type of Validators //
//////////////////////////////////

// Project_id should start with prj_ and then 16 alphanumeric characters.
func ProjectIdValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile("^prj_[0-9A-Za-z_]{16}$"),
		"Please provide a valid name for the project_id, for example: prj_abcdABCD01234567",
	)
}

// Backup Retention Period should be a value between one of the
// * 1d and 180d
// * 1w and 25w
// * 1m and 6m
func BackupRetentionPeriodValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile("^[1-9][0-9]?|1[0-7][0-9]|180d|[1-9]|1[0-9]|2[0-5]w|[1-6]m$"),
		"Please provide a valid value for the backup retention period, for example: 7d, 2w, or 3m.",
	)
}

func startTimeValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`),
		"Please provide a valid time for start time from 00:00 to 23:59, for example: 07:00",
	)
}
