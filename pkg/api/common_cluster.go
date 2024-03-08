package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CommonCluster struct {
	API
}

func (c CommonCluster) ClusterPause(ctx context.Context, projectId, clusterId string) (string, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/clusters/%s/pause", projectId, clusterId)
	body, err := c.doRequest(ctx, http.MethodPost, url, nil)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ClusterId, err
}

func (c CommonCluster) ClusterResume(ctx context.Context, projectId, clusterId string) (string, error) {
	response := struct {
		Data struct {
			ClusterId string `json:"clusterId"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("projects/%s/clusters/%s/resume", projectId, clusterId)
	body, err := c.doRequest(ctx, http.MethodPost, url, nil)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	return response.Data.ClusterId, err
}
