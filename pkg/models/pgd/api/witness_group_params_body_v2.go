package api

type WitnessGroupParamsBodyV2 struct {
	Provider *CloudProvider `json:"provider,omitempty"`
	Region   *Region        `json:"region,omitempty"`
}
