package api

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (baerr *BigAnimalError) GetDetails() string {
	var details string
	for _, err := range baerr.APIError.Error.Errors {
		details += fmt.Sprintln(err.Message)
	}
	return details
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
