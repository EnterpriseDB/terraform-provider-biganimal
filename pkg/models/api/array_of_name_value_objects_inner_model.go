package api

type ArrayOfNameValueObjectsInner struct {
	Name  *string `json:"name,omitempty" tfsdk:"name"`
	Value *string `json:"value,omitempty" tfsdk:"value"`
}
