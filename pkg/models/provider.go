package models

type Provider struct {
	CloudProviderId   string `json:"cloudProviderId"`
	CloudProviderName string `json:"cloudProviderName,omitempty"`
	Connected         *bool  `json:"connected,omitempty"`
}

func (p Provider) String() string {
	return p.CloudProviderId
}
