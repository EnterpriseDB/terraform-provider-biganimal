package models

type Provider struct {
	CloudProviderId   string `json:"cloudProviderId"`
	CloudProviderName string `json:"cloudProviderName,omitempty"`
	Connected         *bool  `json:"connected,omitempty"`
}

func (p Provider) String() string {
	return p.CloudProviderId
}

type AWSConnection struct {
	ExternalID string `json:"externalId"`
	RoleArn    string `json:"roleArn"`
}

type AzureConnection struct {
	ClientId       string `json:"clientId"`
	ClientSecret   string `json:"clientSecret"`
	SubscriptionId string `json:"subscriptionId"`
	TenantId       string `json:"tenantId"`
}
