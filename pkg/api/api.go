package api

import (
	"fmt"
	"net/http"
	"time"
)

type API struct {
	Url   string
	Token string

	httpClient *http.Client
}

func NewAPI(ba_bearer_token string, ba_api_uri string) *API {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	api := &API{
		ba_api_uri,
		ba_bearer_token,
		httpClient,
	}

	return api
}

func (api *API) ClusterClient() *ClusterClient {
	c := NewClusterClient(api.Url, api.Token)
	return c
}

func (api *API) RegionClient() *RegionClient {
	c := NewRegionClient(api.Url, api.Token)
	return c
}

func BuildAPI(meta any) *API {
	api, ok := meta.(*API)
	if !ok {
		panic(fmt.Sprintf("unknown api type.  %T != %T", api, meta))
	}
	return api
}
