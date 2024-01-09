package model

type (
	ResponseBody struct {
		Status   int         `json:"status"`
		Message  string      `json:"message,omitempty"`
		Endpoint string      `json:"endpoint,omitempty"`
		Request  interface{} `json:"request,omitempty"`
		Data     interface{} `json:"data,omitempty"`
	}
)
