package models

type Region struct {
	Id        string `json:"regionId,omitempty" tfsdk:"region_id"`
	Name      string `json:"regionName,omitempty" tfsdk:"name"`
	Status    string `json:"status,omitempty" tfsdk:"status"`
	Continent string `json:"continent,omitempty" tfsdk:"continent"`
}

func (r Region) String() string {
	return r.Id
}
