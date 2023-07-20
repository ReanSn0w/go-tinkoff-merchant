package tinkoff

import "fmt"

type ErrorResponse struct {
	Timestamp    string       `json:"timestamp"`
	Status       int64        `json:"status"`
	ErrorMessage string       `json:"error,omitempty"`
	ErrorFiedls  []ErrorField `json:"errors,omitempty"`
	Message      string       `json:"message,omitempty"`
	Path         string       `json:"path,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("response error: %v", e)
}

type ErrorField struct {
	Field          *string `json:"field,omitempty"`
	DefaultMessage *string `json:"defaultMessage,omitempty"`
	RejectedValue  *string `json:"rejectedValue,omitempty"`
	Code           *string `json:"code,omitempty"`
}
