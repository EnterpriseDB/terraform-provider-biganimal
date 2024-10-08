package api

type CSPTagRequest struct {
	AddTags    []AddTag  `json:"addTags"`
	DeleteTags []string  `json:"deleteTags"`
	EditTags   []EditTag `json:"editTags"`
}

type AddTag struct {
	CspTagKey   string `json:"cspTagKey"`
	CspTagValue string `json:"cspTagValue"`
}

type EditTag struct {
	CSPTagID    string `json:"cspTagId"`
	CSPTagKey   string `json:"cspTagKey"`
	CSPTagValue string `json:"cspTagValue"`
	Status      string `json:"status"`
}
