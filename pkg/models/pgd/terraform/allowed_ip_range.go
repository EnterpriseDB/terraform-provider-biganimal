package terraform

type AllowedIpRange struct {
	CidrBlock   string `tfsdk:"cidr_block"`
	Description string `tfsdk:"description"`
}
