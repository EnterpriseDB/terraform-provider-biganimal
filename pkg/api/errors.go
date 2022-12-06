package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	Error400     = errors.New("Bad Request")
	Error401     = errors.New("Not Authorized")
	Error403     = errors.New("Forbidden")
	Error404     = errors.New("Resource Not Found")
	Error409     = errors.New("Conflict")
	Error412     = errors.New("Precondition Failed")
	Error429     = errors.New("Too Many Requests")
	Error500     = errors.New("API Internal Error")
	ErrorUnknown = errors.New("Unknown API Error")
)

func getStatusError(code int, body []byte) error {
	if code <= 299 {
		return nil
	}

	apiErr := APIError{}

	err := json.Unmarshal(body, &apiErr)
	if err != nil {
		return err
	}
	return fmt.Errorf("%s %v", apiErr.Error.Status, apiErr.Error.Errors)

	// switch code {
	// case 200:
	// 	return nil
	// case 202:
	// 	return nil
	// case 204:
	// 	return nil
	// case 400:
	// 	return Error400
	// case 401:
	// 	return Error401
	// case 403:
	// 	return Error403
	// case 404:
	// 	return Error404
	// case 409:
	// 	return Error409
	// case 412:
	// 	return Error412
	// case 429:
	// 	return Error429
	// case 500:
	// 	return Error500
	// default:
	// 	return Error500
	// }
}

type APIError struct {
	Error struct {
		Status    int      `json:"status"`
		Message   string   `json:"message"`
		Errors    []string `json:"errors"`
		Reference string   `json:"reference"`
		Source    string   `json:"source"`
	} `json:"error"`
}
