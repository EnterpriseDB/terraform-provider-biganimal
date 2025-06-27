package api

type CSPTagResponse struct {
	Data []struct {
		CSPTagID    string `json:"cspTagId"`
		CSPTagKey   string `json:"cspTagKey"`
		CSPTagValue string `json:"cspTagValue"`
		Status      string `json:"status"`
	} `json:"data"`
}
