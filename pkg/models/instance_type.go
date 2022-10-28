package models

type InstanceType struct {
	Category         string  `json:"category,omitempty"`
	Cpu              float64 `json:"cpu,omitempty"`
	FamilyName       string  `json:"familyName,omitempty"`
	InstanceTypeId   string  `json:"instanceTypeId"`
	InstanceTypeName string  `json:"instanceTypeName,omitempty"`
	Ram              float64 `json:"ram,omitempty"`
}

func (i InstanceType) String() string {
	return i.InstanceTypeId
}
