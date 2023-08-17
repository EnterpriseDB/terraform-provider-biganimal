package api

type WitnessGroupParamsBody struct {
	Provider *CloudProvider `json:"provider,omitempty"`
	Region   *Region        `json:"region,omitempty"`
}
