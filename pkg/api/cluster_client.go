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
)

type ClusterClient struct{ API }

func NewClusterClient(api API) *ClusterClient {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	api.HTTPClient = httpClient
	c := ClusterClient{API: api}
	return &c
}

func (c ClusterClient) Create(ctx context.Context, projectId string, model any) (string, error) {
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

func (c ClusterClient) Read(ctx context.Context, projectId, id string) (*models.Cluster, error) {
	response := struct {
		Data models.Cluster `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/clusters/%s", projectId, id)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &response.Data, err
	}

	err = json.Unmarshal(body, &response)

	return &response.Data, err
}

func (c ClusterClient) ReadByName(ctx context.Context, projectId, name string, most_recent bool) (*models.Cluster, error) {
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
			return &models.Cluster{}, Error404
		}
	}

	return &clusters.Data[0], err
}

func (c ClusterClient) ConnectionString(ctx context.Context, projectId, id string) (*models.ClusterConnection, error) {
	response := struct {
		Data models.ClusterConnection `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/clusters/%s/connection/", projectId, id)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &models.ClusterConnection{}, err
	}

	if json.Unmarshal(body, &response) != nil {
		return &models.ClusterConnection{}, err
	}
	return &response.Data, nil
}

func (c ClusterClient) Update(ctx context.Context, cluster *models.Cluster, projectId, id string) (*models.Cluster, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/clusters/%s", projectId, id)

	b, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(ctx, http.MethodPut, url, bytes.NewBuffer(b))
	if err != nil {
		return &models.Cluster{}, err
	}

	err = json.Unmarshal(body, &response)
	return nil, err
}

func (c ClusterClient) Delete(ctx context.Context, projectId, id string) error {
	url := fmt.Sprintf("projects/%s/clusters/%s", projectId, id)
	_, err := c.doRequest(ctx, http.MethodDelete, url, nil)
	return err
}

type ClusterResponse struct {
	Data struct {
		ClusterId string `json:"clusterId"`
	} `json:"data"`
}
