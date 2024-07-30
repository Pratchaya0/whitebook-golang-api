package responses

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	// Data   interface{} `json:"data,omitempty"`
}
