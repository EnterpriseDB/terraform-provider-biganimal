package api

import (
	"fmt"
	"net/http"
	"time"
)

const clientTimeoutSeconds = 60

type API struct {
	BaseURL    string
	Token      string
	AccessKey  string
	UserAgent  string
	HTTPClient http.Client
}

func NewAPI(ba_access_key, ba_bearer_token, ba_api_uri, userAgent string) *API {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	api := &API{
		BaseURL:    ba_api_uri,
		Token:      ba_bearer_token,
		AccessKey:  ba_access_key,
		UserAgent:  userAgent,
		HTTPClient: httpClient,
	}

	return api
}

func (api *API) ClusterClient() *ClusterClient {
	c := NewClusterClient(*api)
	return c
}

func (api *API) PGDClient() *PGDClient {
	c := NewPGDClient(*api)
	return c
}

func (api *API) RegionClient() *RegionClient {
	c := NewRegionClient(*api)
	return c
}

func (api *API) ProjectClient() *ProjectClient {
	c := NewProjectClient(*api)
	return c
}

func (api *API) CloudProviderClient() *CloudProviderClient {
	c := NewCloudProviderClient(*api)
	return c
}

func (api *API) TagClient() *TagClient {
	c := NewTagClient(*api)
	return c
}

func BuildAPI(meta any) *API {
	api, ok := meta.(*API)
	if !ok {
		panic(fmt.Sprintf("unknown api type.  %T != %T", api, meta))
	}
	return api
}
