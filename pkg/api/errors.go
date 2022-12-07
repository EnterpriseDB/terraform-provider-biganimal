package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

var (
	Error404     = errors.New("resource not found")
	ErrorUnknown = errors.New("unknown API error")
)

type APIError struct {
	Error struct {
		Status    int            `json:"status"`
		Message   string         `json:"message"`
		Errors    []ErrorMessage `json:"errors"`
		Reference string         `json:"reference"`
		Source    string         `json:"source"`
	} `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type BigAnimalError struct {
	APIError APIError
	Err      error
}

func (baerr *BigAnimalError) Error() string {
	return fmt.Sprintf("status: %d - %v", baerr.APIError.Error.Status, baerr.APIError.Error.Message)
}

func (baerr *BigAnimalError) getDetails() string {
	var details string
	for _, err := range baerr.APIError.Error.Errors {
		details += fmt.Sprintln(err.Message)
	}
	return details
}

func FromErr(err error) diag.Diagnostics {
	if err == nil {
		return nil
	}

	baerr, ok := err.(*BigAnimalError)
	if ok {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprint(baerr.Err),
				Detail:   baerr.getDetails(),
			},
		}
	}
	return diag.FromErr(err)
}

func getStatusError(code int, body []byte) error {
	if code <= 299 {
		return nil
	}

	apiErr := APIError{}

	if json.Unmarshal(body, &apiErr) != nil {
		return ErrorUnknown
	}

	return &BigAnimalError{
		APIError: apiErr,
		Err:      fmt.Errorf("status: %d - %v", apiErr.Error.Status, apiErr.Error.Message),
	}
}
