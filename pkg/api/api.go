package api

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type API struct {
	Url   string
	Token string

	httpClient *http.Client
}

func NewAPI() *API {
	url := os.Getenv("BA_API_URI")
	token := os.Getenv("BA_BEARER_TOKEN")

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	api := &API{
		url,
		token,
		httpClient,
	}

	return api
}

func (api *API) ClusterClient() *ClusterClient {
	c := NewClusterClient(api.Url, api.Token)
	return c
}

func BuildAPI(meta any) *API {
	api, ok := meta.(*API)
	if !ok {
		panic(fmt.Sprintf("unknown api type.  %T != %T", api, meta))
	}
	return api
}
