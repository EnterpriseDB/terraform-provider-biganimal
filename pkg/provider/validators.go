package provider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/google/uuid"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

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

func ProjectIdValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile("^prj_[0-9A-Za-z_]{16}$"),
		"Please provide a valid name for the project_id, for example: prj_abcdABCD01234567",
	)
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

func startDayValidator() validator.Float64 {
	return float64validator.Between(0, 6)
}

func startTimeValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d$`),
		"Please provide a valid time for start time from 00:00 to 23:59, for example: 07:00",
	)
}
