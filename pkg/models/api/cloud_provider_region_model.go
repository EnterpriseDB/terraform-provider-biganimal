package api

type CloudProviderRegion struct {
	Continent          *string `json:"continent,omitempty"`
	CustomerLogsUrl    *string `json:"customerLogsUrl,omitempty"`
	CustomerMetricsUrl *string `json:"customerMetricsUrl,omitempty"`
	RegionId           string  `json:"regionId"`
	RegionName         string  `json:"regionName"`
	Status             *string `json:"status,omitempty"`
}
