package api

type Tag struct {
	Color   *string `json:"color,omitempty"`
	TagId   string  `json:"tagId"`
	TagName string  `json:"tagName"`
}
