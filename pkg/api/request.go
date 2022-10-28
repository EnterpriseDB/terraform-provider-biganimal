package api

import (
	"context"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func doRequest(ctx context.Context, c http.Client, httpMethod, url, token string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, url)

	req.Header.Add("user-agent", userAgent)
	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("content-type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	out, _ := io.ReadAll(res.Body)
	err = getStatusError(res.StatusCode)
	if err != nil {
		tflog.Debug(ctx, string(out))
	}

	return out, err
}
