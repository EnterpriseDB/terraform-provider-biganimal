package provider

import (
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/constants"
)

func TdeActionInfo(provider string) string {
	switch provider {
	case "aws", "bah:aws":
		return fmt.Sprintf(constants.TDE_KEY_AWS_ACTION, provider)
	case "azure", "bah:azure":
		return fmt.Sprintf(constants.TDE_KEY_AZURE_ACTION, provider)
	case "gcp", "bah:gcp":
		return fmt.Sprintf(constants.TDE_KEY_GCP_ACTION, provider)
	default:
		return fmt.Sprintf(constants.TDE_KEY_ACTION_UNKNOWN_PROVIDER_ERROR, provider)
	}
}
