package models

type CloudProvider struct {
	CloudProviderId   string `json:"cloudProviderId,omitempty" tfsdk:"cloud_provider_id"`
	CloudProviderName string `json:"cloudProviderName,omitempty" tfsdk:"cloud_provider_name"`
}

type Project struct {
	ProjectId      string           `json:"projectId,omitempty" tfsdk:"project_id"`
	ProjectName    string           `json:"projectName,omitempty" tfsdk:"project_name"`
	UserCount      int              `json:"userCount,omitempty" tfsdk:"user_count"`
	ClusterCount   int              `json:"clusterCount,omitempty" tfsdk:"cluster_count"`
	CloudProviders []CloudProvider `json:"cloudProviders" tfsdk:"cloud_providers"`
}

// Check the return value, if ProjectName is also needed
func (p Project) String() string {
	return p.ProjectId
}
