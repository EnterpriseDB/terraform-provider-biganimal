package api

type CloudProvider struct {
	CloudProviderId   string `json:"cloudProviderId"`
	CloudProviderName string `json:"cloudProviderName"`
	Connected         *bool  `json:"connected,omitempty"`
}
