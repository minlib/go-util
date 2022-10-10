package json

import (
	"encoding/json"
	"errors"
)

// ToJsonString 对象转为Json字符串
func ToJsonString(v any) string {
	if buf, err := json.Marshal(v); err == nil {
		return string(buf)
	} else {
		panic(errors.New("json serialization error"))
	}
}

// Parse Json字符串转为对象
func Parse(data string, v any) error {
	return json.Unmarshal([]byte(data), v)
}
