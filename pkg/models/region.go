package models

type Region struct {
	Id        string `json:"regionId,omitempty" mapstructure:"id"`
	Name      string `json:"regionName,omitempty" mapstructure:"name,omitempty"`
	Status    string `json:"status,omitempty" mapstructure:"status,omitempty"`
	Continent string `json:"continent,omitempty" mapstructure:"continent,omitempty"`
}

func (r Region) String() string {
	return r.Id
}
