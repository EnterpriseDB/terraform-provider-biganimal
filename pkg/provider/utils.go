package provider

import (
	"errors"

	"fmt"

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
Please use this function for error check after client API calls, for example:

	r.client.Read(ctx, ...)
	if err != nil {
		summary, detail := extractSumAndDetailfromBAErr(err)
		resp.Diagnostics.AddError(summary, detail)
		return
	}
*/
func extractSumAndDetailfromBAErr(err error) (summary, detail string) {
	var baError *api.BigAnimalError
	if errors.As(err, &baError) {
		return baError.Error(), baError.GetDetails()
	}
	return
}

func fromErr(err error, summary string, args ...any) frameworkdiag.Diagnostics {
	summary = fmt.Sprintf(summary, args...)
	return frameworkdiag.Diagnostics{
		frameworkdiag.NewErrorDiagnostic(
			summary, err.Error(),
		),
	}
}
