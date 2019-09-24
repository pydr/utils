package utils

import "net/http"

type (
	Code struct {
		Status     bool   `json:"status"`
		Code       int    `json:"code"`
		Message    string `json:"message"`
		HttpStatus int    `json:"-"`
		Err        error  `json:"-"`
	}

	Response struct {
		*Code
		Data interface{} `json:"data,omitempty"`
	}
)

func MakeCode(httpCode, code int, status bool, msg string) *Code {
	return &Code{
		Status:     status,
		Code:       code,
		Message:    msg,
		HttpStatus: httpCode,
	}
}

func (c *Code) SetErrorMsg(err error) {
	c.Err = err
}

func MakeResp(code *Code, data interface{}) *Response {
	return &Response{
		Code: code,
		Data: data,
	}
}

var (
	Success = MakeCode(http.StatusOK, 0, true, "success")
)
