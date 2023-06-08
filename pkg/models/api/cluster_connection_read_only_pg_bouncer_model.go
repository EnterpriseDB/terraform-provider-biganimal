package api

type ClusterConnectionReadOnlyPgBouncer struct {
	Data *PgConnectionDetails `json:"data,omitempty"`
}
