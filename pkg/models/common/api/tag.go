package api

// mapstructure annotations are used by faraway away replica only and should be removed once we migrate faraway replica resouce to terraform plugin framework
type Tag struct {
	Color   *string `json:"color,omitempty" mapstructure:"color"`
	TagId   string  `json:"tagId" mapstructure:"tag_id"`
	TagName string  `json:"tagName" mapstructure:"tag_name"`
}
