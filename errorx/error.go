package errorx

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

func (e *Error) Error() string {
	return e.Message
}
