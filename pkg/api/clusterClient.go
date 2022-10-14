package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	userAgent  string
}

func NewClusterClient(url, token string) *ClusterClient {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	c := ClusterClient{
		URL:        url,
		Token:      token,
		HTTPClient: httpClient,
		userAgent:  userAgent,
	}

	return &c
}

func (c ClusterClient) Create(ctx context.Context, model any) (string, error) {
	cluster := model.(apiv2.ClustersBody)

	url := fmt.Sprintf("%s/clusters", c.URL)

	b, err := json.Marshal(cluster)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	req.Header.Add("user-agent", c.userAgent)
	req.Header.Add("authorization", "Bearer "+c.Token)
	req.Header.Add("content-type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	if err := getStatusError(res.StatusCode); err != nil {
		return "", err
	}

	response := ClusterResponse{}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)
	return response.Data.ClusterId, err
}

func (c ClusterClient) Read(ctx context.Context, id string) (*apiv2.ClusterDetail, error) {
	response := struct {
		Data apiv2.ClusterDetail `json:"data"`
	}{}

	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", c.userAgent)
	req.Header.Add("authorization", "Bearer "+c.Token)
	req.Header.Add("content-type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err := getStatusError(res.StatusCode); err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)

	return &response.Data, err
}

func (c ClusterClient) ReadByName(ctx context.Context, name string) (*apiv2.ClusterDetail, error) {
	clusters := struct {
		Data []apiv2.ClusterDetail `json:"data"`
	}{}

	url := fmt.Sprintf("%s/clusters?name=%s", c.URL, name)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &apiv2.ClusterDetail{}, err
	}

	req.Header.Add("user-agent", c.userAgent)
	req.Header.Add("authorization", "Bearer "+c.Token)
	req.Header.Add("content-type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return &apiv2.ClusterDetail{}, err
	}

	if err := getStatusError(res.StatusCode); err != nil {
		return &apiv2.ClusterDetail{}, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &clusters)

	if len(clusters.Data) != 1 {
		return &apiv2.ClusterDetail{}, Error404
	}

	return &clusters.Data[0], err
}

func (c ClusterClient) Update(ctx context.Context, model any, id string) (*apiv2.ClusterDetail, error) {
	cluster := model.(apiv2.ClustersBody)
	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)

	b, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", c.userAgent)
	req.Header.Add("authorization", "Bearer "+c.Token)
	req.Header.Add("content-type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err := getStatusError(res.StatusCode); err != nil {
		return nil, err
	}

	response := ClusterResponse{}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)
	return nil, err
}

func (c ClusterClient) Delete(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("user-agent", c.userAgent)
	req.Header.Add("authorization", "Bearer "+c.Token)
	req.Header.Add("content-type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	return getStatusError(res.StatusCode)
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
