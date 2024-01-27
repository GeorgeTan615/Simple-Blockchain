package errors

import "time"

type ErrorResp struct {
	Message   string    `json:message`
	Timestamp time.Time `json:timestamp`
}

func NewErrorResp(msg string) *ErrorResp {
	return &ErrorResp{
		Message:   msg,
		Timestamp: time.Now(),
	}
}
