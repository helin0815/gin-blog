package model

type RestResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Errors []Error     `json:"errors,omitempty"`
	Total  int64       `json:"total,omitempty" `
}
type Error struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}
