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

type CSPTagClient struct {
	API
}

func NewCSPTagClient(api API) *CSPTagClient {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	api.HTTPClient = httpClient
	tc := CSPTagClient{API: api}
	return &tc
}

func (c CSPTagClient) Put(ctx context.Context, projectID, cloudProviderID string, cspTagReq api.CSPTagRequest) (bool, error) {
	b, err := json.Marshal(cspTagReq)
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf("projects/%s/cloud-providers/%s/tags", projectID, cloudProviderID)
	_, err = c.doRequest(ctx, http.MethodPut, url, bytes.NewBuffer(b))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c CSPTagClient) Get(ctx context.Context, projectID, cloudProviderID string) (*api.CSPTagResponse, error) {
	response := &api.CSPTagResponse{}

	url := fmt.Sprintf("projects/%s/cloud-providers/%s/tags", projectID, cloudProviderID)

	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

func (tc CSPTagClient) Delete(ctx context.Context, tagId string) error {
	url := fmt.Sprintf("tags/%s", tagId)

	_, err := tc.doRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	return nil
}
