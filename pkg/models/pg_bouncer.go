package models

type PgBouncer struct {
	IsEnabled bool                 `json:"isEnabled"`
	Settings  *[]PgBouncerSettings `json:"settings,omitempty"`
}
