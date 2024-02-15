package models

type PgBouncerSettings struct {
	Name      *string `json:"name,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Value     *string `json:"value,omitempty"`
}
