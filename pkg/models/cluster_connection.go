package models

type ClusterConnection struct {
	DatabaseName            string `json:"databaseName"`
	PgUri                   string `json:"pgUri"`
	Port                    string `json:"port"`
	ServiceName             string `json:"serviceName"`
	ReadOnlyPgUri           string `json:"readOnlyPgUri"`
	ReadOnlyPort            string `json:"readOnlyPort"`
	ReadOnlyServiceName     string `json:"readOnlyServiceName"`
	Username                string `json:"username"`
	PrivateLinkServiceAlias string `json:"privateLinkServiceAlias"`
	PrivateLinkServiceName  string `json:"privateLinkServiceName"`
}
