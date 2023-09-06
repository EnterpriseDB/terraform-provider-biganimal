package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type CloudProviderClient struct{ API }

func NewCloudProviderClient(api API) *CloudProviderClient {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	api.HTTPClient = httpClient
	c := CloudProviderClient{API: api}

	return &c
}

func (c CloudProviderClient) CreateAWSConnection(ctx context.Context, projectID string, conn models.AWSConnection) error {
	b, err := json.Marshal(conn)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("projects/%s/cloud-providers/aws/register", projectID)
	_, err = c.doRequest(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	return err
}

func (c CloudProviderClient) GetAWSConnection(ctx context.Context, projectID string) (*models.AWSConnection, error) {
	url := fmt.Sprintf("/projects/%s/cloud-providers/aws/register", projectID)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp := models.AWSConnection{}
	err = json.Unmarshal(body, &resp)

	return &resp, err
}

func (c CloudProviderClient) CreateAzureConnection(ctx context.Context, projectID string, conn models.AzureConnection) error {
	b, err := json.Marshal(conn)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("projects/%s/cloud-providers/azure/register", projectID)
	_, err = c.doRequest(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	return err
}

func (c CloudProviderClient) GetServiceAccountIds(ctx context.Context, projectID string, cspID string, regionID string, saIdsData models.ServiceAccountIds) error {
	b, err := json.Marshal(saIdsData)
	if err != nil {
		return err
	}

	req := fmt.Sprintf("projects/%s/cloud-providers/%s/regions/%s/service-account-ids", projectID, cspID, regionID)
	_, err = c.doRequest(ctx, http.MethodGet, req, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	return err
}

func (c CloudProviderClient) GetPeAllowedPrincipalIds(ctx context.Context, projectID string, cspID string, regionID string, cspAcIdsData models.PeAllowedPrincipalIds) error {
	b, err := json.Marshal(cspAcIdsData)
	if err != nil {
		return err
	}

	req := fmt.Sprintf("projects/%s/cloud-providers/%s/regions/%s/pe-allowed-principal-ids", projectID, cspID, regionID)
	_, err = c.doRequest(ctx, http.MethodGet, req, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	return err
}
