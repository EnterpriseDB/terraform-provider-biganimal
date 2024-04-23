package api

import (
	"net/http"
	"time"
)

type BeaconAnalyticsClient struct {
	API
}

func NewBeaconAnalyticsClient(api API) *BeaconAnalyticsClient {
	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	api.HTTPClient = httpClient
	c := BeaconAnalyticsClient{API: api}

	return &c
}
