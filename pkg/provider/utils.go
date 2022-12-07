package provider

import (
	"errors"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func BigAnimalFromErr(err error) diag.Diagnostics {
	if err == nil {
		return nil
	}

	var baError *api.BigAnimalError
	if errors.As(err, &baError) {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprint(baError.Err),
				Detail:   baError.GetDetails(),
			},
		}
	}
	return diag.FromErr(err)
}
