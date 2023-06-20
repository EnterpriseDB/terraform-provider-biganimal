package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
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
