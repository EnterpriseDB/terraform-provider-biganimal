package api

type CloudProviderRegionInstanceType struct {
	Category         string  `json:"category"`
	Cpu              float64 `json:"cpu"`
	FamilyName       string  `json:"familyName"`
	InstanceTypeId   string  `json:"instanceTypeId"`
	InstanceTypeName string  `json:"instanceTypeName"`
	Ram              float64 `json:"ram"`
}
