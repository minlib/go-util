package core

import (
	"database/sql/driver"
	"encoding/json"
)

type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NameValueSlice []NameValue

func (t *NameValueSlice) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *NameValueSlice) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &t)
}
