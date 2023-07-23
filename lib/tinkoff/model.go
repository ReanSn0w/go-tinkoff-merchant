package tinkoff

import "fmt"

type ErrorResponse struct {
	Status       int64        `json:"status"`
	Timestamp    string       `json:"timestamp,omitempty"`
	ErrorMessage string       `json:"error,omitempty"`
	ErrorFiedls  []ErrorField `json:"errors,omitempty"`
	Message      string       `json:"message,omitempty"`
	Path         string       `json:"path,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprint("response error: ", e)
}

type ErrorField struct {
	Field          *string `json:"field,omitempty"`
	DefaultMessage *string `json:"defaultMessage,omitempty"`
	RejectedValue  *string `json:"rejectedValue,omitempty"`
	Code           *string `json:"code,omitempty"`
}
