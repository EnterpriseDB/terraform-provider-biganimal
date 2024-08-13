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

type ProjectClient struct{ API }

func NewProjectClient(api API) *ProjectClient {
	httpClient := http.Client{
		Timeout: clientTimeoutSeconds * time.Second,
	}

	api.HTTPClient = httpClient
	c := ProjectClient{API: api}
	return &c
}

func (c ProjectClient) Create(ctx context.Context, projectName string) (string, error) {
	response := struct {
		Data struct {
			ProjectId string `json:"projectId"`
		} `json:"data"`
	}{}

	project := map[string]string{"projectName": projectName}

	b, err := json.Marshal(project)
	if err != nil {
		return "", err
	}

	url := "projects"
	body, err := c.doRequest(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ProjectId, err
}

func (c ProjectClient) Read(ctx context.Context, id string) (*models.Project, error) {
	response := struct {
		Data models.Project `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s", id)
	body, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &response.Data, err
	}

	err = json.Unmarshal(body, &response)

	return &response.Data, err
}

func (c ProjectClient) List(ctx context.Context, query string) ([]*models.Project, error) {
	response := struct {
		Data []*models.Project `json:"data"`
	}{}

	url := "projects"
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

func (c ProjectClient) Update(ctx context.Context, projectId, projectName string) (string, error) {
	response := struct {
		Data struct {
			ProjectId string `json:"projectId"`
		} `json:"data"`
	}{}

	project := map[string]string{"projectName": projectName}
	b, err := json.Marshal(project)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("projects/%s", projectId)
	body, err := c.doRequest(ctx, http.MethodPut, url, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ProjectId, err
}

func (c ProjectClient) Delete(ctx context.Context, projectId string) error {
	url := fmt.Sprintf("projects/%s", projectId)
	_, err := c.doRequest(ctx, http.MethodDelete, url, nil)
	return err
}
