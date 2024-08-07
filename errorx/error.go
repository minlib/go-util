package errorx

import (
	"fmt"
)

// Error is a trivial implementation of error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// NewParams returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func NewParams(code int, message string, params ...any) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(message, params...),
	}
}

// Error return error message
func (e *Error) Error() string {
	return e.Message
}

// Format return a formatted and new error object
func (e *Error) Format(params ...any) *Error {
	return &Error{
		Code:    e.Code,
		Message: fmt.Sprintf(e.Message, params...),
	}
}

// MessageOf return a message and new error object
func (e *Error) MessageOf(message string) *Error {
	return &Error{
		Code:    e.Code,
		Message: message,
	}
}
