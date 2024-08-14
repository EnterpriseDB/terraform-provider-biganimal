package api

type TagRequest struct {
	Color   *string `json:"color,omitempty"`
	TagName string  `json:"tagName"`
}
