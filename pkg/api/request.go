package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (api API) doRequest(ctx context.Context, httpMethod, path string, body io.Reader) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", api.BaseURL, path)
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, url)

	req.Header.Add("user-agent", api.UserAgent)
	req.Header.Add("content-type", "application/json")

	if api.AccessKey != "" {
		req.Header.Add("x-access-key", api.AccessKey)
	} else {
		req.Header.Add("authorization", "Bearer "+api.Token)
	}

	res, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			tflog.Error(ctx, "Error closing response body", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}()

	out, _ := io.ReadAll(res.Body)

	err = getStatusError(res.StatusCode, out)
	if err != nil {
		tflog.Debug(ctx, string(out))
	}

	return out, err
}
