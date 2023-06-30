package provider

import (
	"errors"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"

	sdkdiag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func fromBigAnimalErr(err error) sdkdiag.Diagnostics {
	if err == nil {
		return nil
	}

	var baError *api.BigAnimalError
	if errors.As(err, &baError) {
		return sdkdiag.Diagnostics{
			sdkdiag.Diagnostic{
				Severity: sdkdiag.Error,
				Summary:  baError.Error(),
				Detail:   baError.GetDetails(),
			},
		}
	}
	return sdkdiag.FromErr(err)
}

func unsupportedWarning(message string) sdkdiag.Diagnostics {
	return sdkdiag.Diagnostics{
		sdkdiag.Diagnostic{
			Severity: sdkdiag.Warning,
			Summary:  "Unsupported",
			Detail:   message,
		},
	}
}

/*
Please use this function for error check after client API calls.
This function returns a boolean representing if passed error is as *api.BigAnimalError, for example:

	r.client.Read(ctx, ...)
	if err != nil {
		if appendDiagFromBAErr(err){
			return
		}

		resp.Diagnostics.AddError(summary, detail)
		return
	}
*/
func appendDiagFromBAErr(err error, diags *frameworkdiag.Diagnostics) bool {
	var baError *api.BigAnimalError
	if errors.As(err, &baError) {
		diags.AddError(baError.Error(), baError.GetDetails())
		return true
	}
	return false
}
