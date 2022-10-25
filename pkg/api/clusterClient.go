package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/kr/pretty"
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
		Timeout: 10 * time.Second,
	}

	c := ClusterClient{
		URL:        url,
		Token:      token,
		HTTPClient: httpClient,
	}

	return &c
}

func (c ClusterClient) Create(ctx context.Context, model any) (string, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	cluster := model.(models.Cluster)

	url := fmt.Sprintf("%s/clusters", c.URL)
	b, err := json.Marshal(cluster)
	if err != nil {
		return "", err
	}
	tflog.Debug(ctx, pretty.Sprintf(string(b)))
	body, err := doRequest(ctx, *c.HTTPClient, http.MethodPost, url, c.Token, bytes.NewBuffer(b))

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ClusterId, err
}

func (c ClusterClient) Read(ctx context.Context, id string) (*models.Cluster, error) {
	response := struct {
		Data models.Cluster `json:"data"`
	}{}

	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)
	body, err := doRequest(ctx, *c.HTTPClient, http.MethodGet, url, c.Token, nil)
	if err != nil {
		return &response.Data, err
	}

	tflog.Debug(ctx, "fuckery! "+string(body))
	err = json.Unmarshal(body, &response)

	return &response.Data, err
}

func (c ClusterClient) ReadByName(ctx context.Context, name string) (*models.Cluster, error) {
	clusters := struct {
		Data []models.Cluster `json:"data"`
	}{}

	url := fmt.Sprintf("%s/clusters?name=%s", c.URL, name)
	body, err := doRequest(ctx, *c.HTTPClient, http.MethodGet, url, c.Token, nil)
	if err != nil {
		return &models.Cluster{}, err
	}

	if err := json.Unmarshal(body, &clusters); err != nil {
		return &models.Cluster{}, err
	}

	if len(clusters.Data) != 1 {
		return &models.Cluster{}, Error404
	}

	return &clusters.Data[0], err
}

func (c ClusterClient) Update(ctx context.Context, model any, id string) (*models.Cluster, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	cluster := model.(models.Cluster)
	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)

	b, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	body, err := doRequest(ctx, *c.HTTPClient, http.MethodGet, url, c.Token, bytes.NewBuffer(b))
	if err != nil {
		return &models.Cluster{}, err
	}

	err = json.Unmarshal(body, &response)
	return nil, err
}

func (c ClusterClient) Delete(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)
	_, err := doRequest(ctx, *c.HTTPClient, http.MethodDelete, url, c.Token, nil)
	return err
}

func (c *ClusterClient) HasOkCondition(conditions []models.Condition) bool {
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
