package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

const (
	REGION_ACTIVE    = "ACTIVE"
	REGION_INACTIVE  = "INACTIVE"
	REGION_SUSPENDED = "SUSPENDED"
)

type RegionClient struct{ API }

func NewRegionClient(api API) *RegionClient {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	api.HTTPClient = httpClient
	c := RegionClient{API: api}

	return &c
}

func (c RegionClient) Create(ctx context.Context, model any) (string, error) {
	panic("Create not implemented")
}

func (c RegionClient) Read(ctx context.Context, project_id, csp_id, id string) (*models.Region, error) {
	regions, err := c.List(ctx, project_id, csp_id, id)
	if err != nil {
		return nil, err
	}
	for _, region := range regions {
		if region.Id == id {
			return region, nil
		}
	}

	return nil, errors.New("unable to find a unique region")
}

func (c RegionClient) List(ctx context.Context, project_id, csp_id, query string) ([]*models.Region, error) {
	response := struct {
		Data []*models.Region `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/cloud-providers/%s/regions", project_id, csp_id)
	if query != "" {
		url += fmt.Sprintf("?q=%s", query)
	}

	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return response.Data, err
	}

	err = json.Unmarshal(body, &response)

	return response.Data, err
}

func (c RegionClient) Update(ctx context.Context, action, project_id, csp_id, region_id string) error {
	url := fmt.Sprintf("projects/%s/cloud-providers/%s/regions/%s", project_id, csp_id, region_id)

	switch action {
	case REGION_ACTIVE:
		url = url + "/activate"
	case REGION_INACTIVE:
		url = url + "/delete"
	case REGION_SUSPENDED:
		url = url + "/suspend"
	default:
		return errors.New("unknown region action")
	}

	_, err := c.doRequest(ctx, http.MethodPost, url, nil)
	return err
}
