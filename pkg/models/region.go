package models

type Region struct {
	RegionId   string `json:"regionId,omitempty"`
	RegionName string `json:"regionName,omitempty"`
}

func (r Region) String() string {
	return r.RegionId
}
