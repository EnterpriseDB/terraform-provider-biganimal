package models

import commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"

type Region struct {
	Id        string          `json:"regionId,omitempty" tfsdk:"region_id"`
	Name      string          `json:"regionName,omitempty" tfsdk:"name"`
	Status    string          `json:"status,omitempty" tfsdk:"status"`
	Continent string          `json:"continent,omitempty" tfsdk:"continent"`
	Tags      []commonApi.Tag `json:"tags,omitempty"`
}

func (r Region) String() string {
	return r.Id
}
