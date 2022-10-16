package api

import (
	"io"
	"net/http"
)

func doRequest(c http.Client, httpMethod, url, token string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", userAgent)
	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("content-type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if err := getStatusError(res.StatusCode); err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	out, _ := io.ReadAll(res.Body)
	return out, nil
}
