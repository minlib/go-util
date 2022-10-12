package jsonx

import (
	"encoding/json"
	"errors"
)

// Marshal returns the JSON encoding string of v.
func Marshal(v any) string {
	if buf, err := json.Marshal(v); err == nil {
		return string(buf)
	} else {
		panic(errors.New("json serialization error"))
	}
}

// Unmarshal parses the JSON-encoded data and stores the result.
func Unmarshal(data string, v any) error {
	return json.Unmarshal([]byte(data), v)
}
