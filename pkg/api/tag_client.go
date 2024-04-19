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

func (tc TagClient) Create(ctx context.Context, tagReq api.TagRequest) (*string, error) {
	response := struct {
		Data struct {
			TagId string `json:"tagId"`
		} `json:"data"`
	}{}

	b, err := json.Marshal(tagReq)
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

func (tc TagClient) Get(ctx context.Context, tagId string) (api.TagResponse, error) {
	response := struct {
		Data api.TagResponse `json:"data"`
	}{}

	url := fmt.Sprintf("tags/%s", tagId)

	body, err := tc.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return response.Data, err
	}

	err = json.Unmarshal(body, &response)

	return response.Data, err
}

func (tc TagClient) List(ctx context.Context) ([]api.TagResponse, error) {
	response := struct {
		Data []api.TagResponse `json:"data"`
	}{}

	body, err := tc.doRequest(ctx, http.MethodGet, "tags", nil)
	if err != nil {
		return response.Data, err
	}

	err = json.Unmarshal(body, &response)

	return response.Data, err
}
