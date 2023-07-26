package tinkoff

import "fmt"

type ErrorResponse struct {
	Status       int64        `json:"status"`
	Timestamp    string       `json:"timestamp,omitempty"`
	ErrorCode    string       `json:"ErrorCode,omitempty"`
	ErrorMessage string       `json:"error,omitempty"`
	ErrorFiedls  []ErrorField `json:"errors,omitempty"`
	Message      string       `json:"Message,omitempty"`
	Path         string       `json:"path,omitempty"`
	Details      string       `json:"Details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("response error \n status: %v \n message: %v, fields: %v, error: %v \n, details: %v\n", e.Status, e.Message, e.ErrorFiedls, e.ErrorMessage, e.Details)
}

type ErrorField struct {
	Field          *string `json:"field,omitempty"`
	DefaultMessage *string `json:"defaultMessage,omitempty"`
	RejectedValue  *string `json:"rejectedValue,omitempty"`
	Code           *string `json:"code,omitempty"`
}
