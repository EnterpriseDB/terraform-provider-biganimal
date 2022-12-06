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

type APIError struct {
	Error struct {
		Status    int      `json:"status"`
		Message   string   `json:"message"`
		Errors    []ErrorMessage `json:"errors"`
		Reference string   `json:"reference"`
		Source    string   `json:"source"`
	} `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func getStatusError(code int, body []byte) error {
	if code <= 299 {
		return nil
	}

	apiErr := APIError{}

	if json.Unmarshal(body, &apiErr) != nil {
		// FIXME: Is ErrorUnknown correct in that case?
		return ErrorUnknown
	}
	// FIXME: Not only the first message, but we need to print all the messages.
	errmessage := apiErr.Error.Errors[0].Message

	return fmt.Errorf("API Error %d: %s", apiErr.Error.Status, errmessage)

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