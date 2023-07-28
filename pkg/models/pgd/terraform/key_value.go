package terraform

type KeyValue struct {
	Name  string `tfsdk:"name"`
	Value string `tfsdk:"value"`
}
