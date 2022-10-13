package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/apiv2"
)

var (
	Error400     = errors.New("Bad Request")
	Error401     = errors.New("Not Authorized")
	Error403     = errors.New("Formbidden")
	Error404     = errors.New("Resource Not Found")
	Error409     = errors.New("Conflict")
	Error412     = errors.New("Precondition Failed")
	Error429     = errors.New("Too Many Requests")
	Error500     = errors.New("API Internal Error")
	ErrorUnknown = errors.New("Unknown API Error")
)

type ApiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
	URL        string
	Token      string
	HTTPClient *http.Client
	userAgent  string
}

func NewClient() (*ApiClient, error) {
	url := os.Getenv("BA_API_URI")
	token := os.Getenv("BA_BEARER_TOKEN")
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	c := ApiClient{
		URL:        url,
		Token:      token,
		HTTPClient: httpClient,
		userAgent:  "terraform-provider-biganimal",
	}

	return &c, nil
}

func (c ApiClient) GetClusterByID(ctx context.Context, id string) (*apiv2.ClusterDetail, error) {
	response := struct {
		Data apiv2.ClusterDetail `json:"data"`
	}{}

	url := fmt.Sprintf("%s/clusters/%s", c.URL, id)
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
	err = json.Unmarshal(body, &response)

	return &response.Data, err
}

func (c ApiClient) GetClusterByName(ctx context.Context, name string) (*apiv2.ClusterDetail, error) {
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

func (c ApiClient) DeleteClusterByID(ctx context.Context, id string) error {
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

func (c ApiClient) CreateCluster(ctx context.Context, cluster apiv2.ClustersBody) (*CreateResponse, error) {
	url := fmt.Sprintf("%s/clusters", c.URL)

	b, err := json.Marshal(cluster)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
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

	response := CreateResponse{}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)
	return &response, err
}

func (c ApiClient) UpdateClusterById(ctx context.Context, cluster apiv2.ClustersClusterIdBody, id string) (*CreateResponse, error) {
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

	response := CreateResponse{}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)
	return &response, err
}

func getStatusError(code int) error {
	switch code {
	case 200:
		return nil
	case 202:
		return nil
	case 204:
		return nil
	case 400:
		return Error400
	case 401:
		return Error401
	case 403:
		return Error403
	case 404:
		return Error404
	case 409:
		return Error409
	case 412:
		return Error412
	case 429:
		return Error429
	case 500:
		return Error500
	default:
		return Error500
	}
}

func (c *ApiClient) HasOkCondition(conditions []apiv2.ClusterDetailConditions) bool {
	for _, cond := range conditions {
		if *cond.Type_ == "biganimal.com/deployed" && *cond.ConditionStatus == "True" {
			return true
		}
	}
	return false
}
