package api

type ClusterConnection struct {
	DatabaseName        string                              `json:"databaseName"`
	PgUri               *string                             `json:"pgUri,omitempty"`
	Port                string                              `json:"port"`
	ReadOnlyPgBouncer   *ClusterConnectionReadOnlyPgBouncer `json:"readOnlyPgBouncer,omitempty"`
	ReadOnlyPgUri       *string                             `json:"readOnlyPgUri,omitempty"`
	ReadOnlyPort        *string                             `json:"readOnlyPort,omitempty"`
	ReadOnlyServiceName *string                             `json:"readOnlyServiceName,omitempty"`
	ReadWritePgBouncer  *ClusterConnectionReadOnlyPgBouncer `json:"readWritePgBouncer,omitempty"`
	ServiceName         string                              `json:"serviceName"`
	Username            string                              `json:"username"`
}
