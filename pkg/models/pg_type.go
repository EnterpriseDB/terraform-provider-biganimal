package models

type PgType struct {
	PgTypeId                        string    `json:"pgTypeId"`
	PgTypeName                      string    `json:"pgTypeName,omitempty"`
	SupportedClusterArchitectureIds *[]string `json:"supportedClusterArchitectureIds,omitempty"`
}

func (p PgType) String() string {
	return p.PgTypeId
}
