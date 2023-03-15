package provider

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func validateProjectId(v interface{}, path cty.Path) diag.Diagnostics {
	value := v.(string)
	var diags diag.Diagnostics
	//if value != can(regex("^prj_[[:alnum:]]{16}$", value)) {
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
			Detail:   fmt.Sprintf("%v is a invliad aws role arn", v),
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
