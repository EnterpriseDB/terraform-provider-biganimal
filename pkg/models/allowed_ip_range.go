package models

type AllowedIpRange struct {
	CidrBlock   string `json:"cidrBlock" mapstructure:"cidr_block" tfsdk:"cidr_block"`
	Description string `json:"description" mapstructure:"description" tfsdk:"description"`
}
