package json

import (
	"encoding/json"
	"errors"
)

func ToJsonString(v any) string {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(errors.New("json serialization error"))
	}
	return string(buf)
}

func Parse(data string, v any) error {
	return json.Unmarshal([]byte(data), v)
}
