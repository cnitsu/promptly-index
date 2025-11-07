// Package response base request & response struct
package response

type Response struct {
	Status  int    `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"msg"`
}

func Success(data any) *Response {
	return &Response{
		Status:  200,
		Data:    data,
		Message: "OK",
	}
}

func Error(msg string) *Response {
	return &Response{
		Status:  400,
		Data:    nil,
		Message: msg,
	}
}
