package models

type CloudProvider struct {
	CloudProviderId   string `json:"cloudProviderId,omitempty" mapstructure:"cloud_provider_id"`
	CloudProviderName string `json:"cloudProviderName,omitempty" mapstructure:"cloud_provider_name"`
}

type Project struct {
	ProjectId      string           `json:"projectId,omitempty" mapstructure:"project_id"`
	ProjectName    string           `json:"projectName,omitempty" mapstructure:"name"`
	UserCount      int              `json:"userCount,omitempty" mapstructure:"user_count"`
	ClusterCount   int              `json:"clusterCount,omitempty" mapstructure:"cluster_count"`
	CloudProviders *[]CloudProvider `json:"cloudProviders" mapstructure:"cloud_providers"`
}

// Check the return value, if ProjectName is also needed
func (p Project) String() string {
	return p.ProjectId
}
