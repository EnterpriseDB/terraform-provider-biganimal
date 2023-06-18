package api

type ClusterAllowedIpRange struct {
	CidrBlock   string `json:"cidrBlock" tfsdk:"cidr_block"`
	Description string `json:"description" tfsdk:"description"`
}
