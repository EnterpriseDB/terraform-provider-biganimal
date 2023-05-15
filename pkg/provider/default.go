package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type defaultString struct {
	Desc    string
	Default string
}

func DefaultString(desc string, Default string) *defaultString {
	return &defaultString{Desc: desc, Default: Default}
}

func (d defaultString) Description(_ context.Context) string {
	return d.Desc
}

func (d defaultString) MarkdownDescription(_ context.Context) string {
	return d.Desc
}

func (d defaultString) DefaultString(ctx context.Context, request defaults.StringRequest, response *defaults.StringResponse) {
	response.PlanValue = types.StringValue(d.Default)
}
