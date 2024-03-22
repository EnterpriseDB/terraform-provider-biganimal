package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
)

type TagClient struct {
	API
}

func NewTagClient(api API) *TagClient {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	api.HTTPClient = httpClient
	tc := TagClient{API: api}
	return &tc
}

func (tc TagClient) Create(ctx context.Context, tag api.TagRequest) (*string, error) {
	response := struct {
		Data struct {
			TagId string `json:"tagId"`
		} `json:"data"`
	}{}

	b, err := json.Marshal(tag)
	if err != nil {
		return nil, err
	}

	body, err := tc.doRequest(ctx, http.MethodPost, "tags", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	return &response.Data.TagId, err
}

func (c ClusterClient) GetTags(ctx context.Context) ([]api.TagResponse, error) {
	response := struct {
		Data []api.TagResponse `json:"data"`
	}{}

	url := fmt.Sprintf("api/v3/tags")
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return response.Data, err
	}

	err = json.Unmarshal(body, &response)

	return response.Data, err
}
