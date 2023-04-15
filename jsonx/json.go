package jsonx

import (
	"encoding/json"
)

// Marshal returns the JSON encoding string of v.
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal parses the JSON-encoded data and stores the result.
func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// MarshalString returns the JSON encoding string of v.
func MarshalString(v any) string {
	if buf, err := json.Marshal(v); err == nil {
		return string(buf)
	} else {
		return ""
	}
}

// UnmarshalString parses the JSON-encoded data and stores the result.
func UnmarshalString(data string, v any) error {
	return json.Unmarshal([]byte(data), v)
}
