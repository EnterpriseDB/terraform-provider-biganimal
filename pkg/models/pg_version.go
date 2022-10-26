package models

type PgVersion struct {
	PgVersionId   string `json:"pgVersionId"`
	PgVersionName string `json:"pgVersionName,omitempty"`
}

func (p PgVersion) String() string {
	return p.PgVersionId
}
