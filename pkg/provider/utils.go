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

func fromErr(err error, summary string, args ...any) frameworkdiag.Diagnostics {
	summary = fmt.Sprintf(summary, args...)
	return frameworkdiag.Diagnostics{
		frameworkdiag.NewErrorDiagnostic(
			summary, err.Error(),
		),
	}
}
