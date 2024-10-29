package api

// mapstructure annotations are used by faraway away replica only and should be removed once we migrate faraway replica resouce to terraform plugin framework
type Tag struct {
	Color   *string `json:"color,omitempty" tfsdk:"color" mapstructure:"color"`
	TagId   string  `json:"tagId" tfsdk:"tag_id" mapstructure:"tag_id"`
	TagName string  `json:"tagName" tfsdk:"tag_name" mapstructure:"tag_name"`
}
