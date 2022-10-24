package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/apiv2"
)

const (
	userAgent = "terraform-provider-biganimal"
)

type ClusterClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
	URL        string
	Token      string
	HTTPClient *http.Client
}

func NewClusterClient(url, token string) *ClusterClient {
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	c := ClusterClient{
		URL:        url,
		Token:      token,
		HTTPClient: httpClient,
	}

	return &c
}

type something struct {
	Data struct {
		ClusterId string `json:"clusterId"`
	} `json:"data"`
}

func (c ClusterClient) Create(ctx context.Context, model any) (string, error) {
	response := something{}

	cluster := model.(apiv2.ClustersBody)

	url := fmt.Sprintf("%s/clusters", c.URL)
	b, err := json.Marshal(cluster)
	if err != nil {
		return "", err
	}

	body, err := doRequest(ctx, *c.HTTPClient, http.MethodPost, url, c.Token, bytes.NewBuffer(b))

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ClusterId, err
}

func (c ClusterClient) Read(ctx context.Context, id string) (*apiv2.ClusterDetail, error) {
	response := struct {
		Data apiv2.ClusterDetail `json:"data"`
	}{}

	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)
	body, err := doRequest(ctx, *c.HTTPClient, http.MethodGet, url, c.Token, nil)
	if err != nil {
		return &response.Data, err
	}
	err = json.Unmarshal(body, &response)

	// connectionString

	return &response.Data, err
}

func (c ClusterClient) ReadByName(ctx context.Context, name string) (*apiv2.ClusterDetail, error) {
	clusters := struct {
		Data []apiv2.ClusterDetail `json:"data"`
	}{}

	url := fmt.Sprintf("%s/clusters?name=%s", c.URL, name)
	body, err := doRequest(ctx, *c.HTTPClient, http.MethodGet, url, c.Token, nil)
	if err != nil {
		return &apiv2.ClusterDetail{}, err
	}

	if json.Unmarshal(body, &clusters) != nil {
		return &apiv2.ClusterDetail{}, err
	}

	if len(clusters.Data) != 1 {
		return &apiv2.ClusterDetail{}, Error404
	}

	return &clusters.Data[0], err
}

func (c ClusterClient) ConnectionString(ctx context.Context, id string) (*ClusterConnection, error) {
	response := struct {
		Data ClusterConnection `json:"data"`
	}{}

	url := fmt.Sprintf("%s/clusters/%s/connection/", c.URL, id)
	body, err := doRequest(ctx, *c.HTTPClient, http.MethodGet, url, c.Token, nil)
	if err != nil {
		return &ClusterConnection{}, err
	}

	if json.Unmarshal(body, &response) != nil {
		return &ClusterConnection{}, err
	}
	return &response.Data, nil
}

func (c ClusterClient) Update(ctx context.Context, model any, id string) (*apiv2.ClusterDetail, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	cluster := model.(apiv2.ClustersClusterIdBody)
	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)

	b, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	body, err := doRequest(ctx, *c.HTTPClient, http.MethodPut, url, c.Token, bytes.NewBuffer(b))
	if err != nil {
		return &apiv2.ClusterDetail{}, err
	}

	err = json.Unmarshal(body, &response)
	return nil, err
}

func (c ClusterClient) Delete(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)
	_, err := doRequest(ctx, *c.HTTPClient, http.MethodDelete, url, c.Token, nil)
	return err
}

func (c *ClusterClient) HasOkCondition(conditions []apiv2.ClusterDetailConditions) bool {
	for _, cond := range conditions {
		if *cond.Type_ == "biganimal.com/deployed" && *cond.ConditionStatus == "True" {
			return true
		}
	}
	return false
}

type ClusterResponse struct {
	Data struct {
		ClusterId string `json:"clusterId"`
	} `json:"data"`
}

type ClusterConnection struct {
	DatabaseName        string `json:"databaseName"`
	PgUri               string `json:"pgUri"`
	Port                string `json:"port"`
	ServiceName         string `json:"serviceName"`
	ReadOnlyPgUri       string `json:"readOnlyPgUri"`
	ReadOnlyPort        string `json:"readOnlyPort"`
	ReadOnlyServiceName string `json:"readOnlyServiceName"`
	Username            string `json:"username"`
}
