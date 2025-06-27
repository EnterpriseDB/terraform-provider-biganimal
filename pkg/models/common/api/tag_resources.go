package api

type TagResource struct {
	ResourceType string `json:"resourceType"`
	ResourceId   string `json:"resourceId"`
	ResourceName string `json:"resourceName"`
	ProjectId    string `json:"projectId"`
}
