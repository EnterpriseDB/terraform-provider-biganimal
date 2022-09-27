package client

type CreateResponse struct {
	Data struct {
		ClusterId string `json:"clusterId"`
	} `json:"data"`
}

// {
//   "data": {
//     "clusterId": "string"
//   }
// }
