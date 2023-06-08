package api

type PgType struct {
	PgTypeId                        string    `json:"pgTypeId"`
	PgTypeName                      string    `json:"pgTypeName"`
	SupportedClusterArchitectureIds *[]string `json:"supportedClusterArchitectureIds,omitempty"`
}
