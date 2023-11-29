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

func (c ClusterClient) ReadDeletedCluster(ctx context.Context, projectId, id string) (*models.Cluster, error) {
	response := struct {
		Data models.Cluster `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/deleted-clusters/%s", projectId, id)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &response.Data, err
	}

	err = json.Unmarshal(body, &response)

	return &response.Data, err
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
		return nil, err
	}

	if err := json.Unmarshal(body, &clusters); err != nil {
		return nil, err
	}

	if len(clusters.Data) != 1 {
		if most_recent {
			sort.Slice(clusters.Data, func(i, j int) bool { return clusters.Data[i].CreatedAt.Seconds > clusters.Data[j].CreatedAt.Seconds })
		} else {
			return nil, ErrorClustersSameName
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

func (c ClusterClient) GetServiceAccountIds(ctx context.Context, projectID string, cspID string, regionID string) (*models.ServiceAccountIds, error) {
	response := struct {
		Data models.ServiceAccountIds `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/cloud-providers/%s/regions/%s/service-account-ids", projectID, cspID, regionID)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &models.ServiceAccountIds{}, err
	}

	if json.Unmarshal(body, &response.Data) != nil {
		return &models.ServiceAccountIds{}, err
	}
	return &response.Data, nil
}

func (c ClusterClient) GetPeAllowedPrincipalIds(ctx context.Context, projectID string, cspID string, regionID string) (*models.PeAllowedPrincipalIds, error) {
	response := struct {
		Data models.PeAllowedPrincipalIds `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/cloud-providers/%s/regions/%s/pe-allowed-principal-ids", projectID, cspID, regionID)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &models.PeAllowedPrincipalIds{}, err
	}

	if json.Unmarshal(body, &response.Data) != nil {
		return &models.PeAllowedPrincipalIds{}, err
	}
	return &response.Data, nil
}

func (c ClusterClient) RestoreCluster(ctx context.Context, projectId, clusterId string, model models.RestoreCluster) (string, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	b, err := json.Marshal(model)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("projects/%s/clusters/%s/restore", projectId, clusterId)
	body, err := c.doRequest(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ClusterId, err
}

func (c ClusterClient) RestoreClusterFromDeleted(ctx context.Context, projectId, clusterId string, model models.RestoreCluster) (string, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	b, err := json.Marshal(model)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("projects/%s/deleted-clusters/%s/restore", projectId, clusterId)
	body, err := c.doRequest(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ClusterId, err
}
