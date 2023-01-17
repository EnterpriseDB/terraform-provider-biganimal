package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	Error404              = errors.New("resource not found")
	ErrorUnknown          = errors.New("unknown API error")
	ErrorClustersSameName = NewMultipleClustersSameNameError()
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

func NewMultipleClustersSameNameError() error {
	apiErr := &APIError{}
	apiErr.Error.Status = 400
	apiErr.Error.Message = "Bad Request - multiple clusters with the same name"
	apiErr.Error.Errors = []ErrorMessage{{Message: "There are more than one clusters with the same name. Please use most_recent=true to get the latest one."}}

	return &BigAnimalError{
		APIError: *apiErr,
	}
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
	}
}
