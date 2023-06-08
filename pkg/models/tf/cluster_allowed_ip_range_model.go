package tf

type ClusterAllowedIpRange struct {
	CidrBlock   string `tfsdk:"cidr_block"`
	Description string `tfsdk:"description"`
}
