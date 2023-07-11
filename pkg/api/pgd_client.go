package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd"
)

type PGDClient struct{ API }

func NewPGDClient(api API) *PGDClient {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	api.HTTPClient = httpClient
	c := PGDClient{API: api}
	return &c
}

func (c PGDClient) ReadByName(ctx context.Context, projectId, name string, most_recent bool) (*models.Cluster, error) {
	clusters := struct {
		Data []models.Cluster `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/clusters?name=%s", projectId, name)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &models.Cluster{}, err
	}

	if err := json.Unmarshal(body, &clusters); err != nil {
		return &models.Cluster{}, err
	}

	if len(clusters.Data) != 1 {
		if most_recent {
			sort.Slice(clusters.Data, func(i, j int) bool { return clusters.Data[i].CreatedAt.Seconds > clusters.Data[j].CreatedAt.Seconds })
		} else {
			return &models.Cluster{}, ErrorClustersSameName
		}
	}

	return &clusters.Data[0], err
}

func (c PGDClient) Create(ctx context.Context, projectId string, model any) (string, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	cluster := model.(models.Cluster)

	b, err := json.Marshal(cluster)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("projects/%s/clusters", projectId)
	body, err := c.doRequest(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ClusterId, err
}

func (c PGDClient) Read(ctx context.Context, projectId, clusterId string) (*models.Cluster, error) {
	response := struct {
		Data *models.Cluster `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/clusters/%s", projectId, clusterId)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return response.Data, err
	}

	err = json.Unmarshal(body, &response)

	return response.Data, err
}

func (c PGDClient) CalculateWitnessGroupParams(ctx context.Context, projectId string, WitnessGroupParamsBody pgd.WitnessGroupParamsBody) (*pgd.WitnessGroupParamsData, error) {
	var response pgd.WitnessGroupParamsResponse

	b, err := json.Marshal(WitnessGroupParamsBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("projects/%s/utils/calculate-witness-group-params", projectId)
	body, err := c.doRequest(ctx, http.MethodPut, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	return &response.Data, err
}
