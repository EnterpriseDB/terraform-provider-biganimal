package models

type KeyValue struct {
	Name  string `json:"name" mapstructure:"name" tfsdk:"name"`
	Value string `json:"value" mapstructure:"value" tfsdk:"value"`
}
