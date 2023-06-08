package api

type ClusterAllowedIpRange struct {
	CidrBlock   string `json:"cidrBlock"`
	Description string `json:"description"`
}
