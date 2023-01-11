package provider

import (
	"fmt"
	"regexp"

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
