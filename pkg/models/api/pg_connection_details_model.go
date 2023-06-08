package api

type PgConnectionDetails struct {
	DatabaseName string `json:"databaseName"`
	PgUri        string `json:"pgUri"`
	Port         string `json:"port"`
	ServiceName  string `json:"serviceName"`
	Username     string `json:"username"`
}
