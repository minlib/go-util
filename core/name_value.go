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

// Scan implements the Scanner interface.
func (t *NameValueSlice) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &t)
}

// Value implements the driver Valuer interface.
func (t *NameValueSlice) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// GetNames get name list
func (t *NameValueSlice) GetNames() []string {
	var names []string
	if t != nil {
		for _, v := range *t {
			names = append(names, v.Name)
		}
	}
	return names
}

// GetValues get value list
func (t *NameValueSlice) GetValues() []string {
	var values []string
	if t != nil {
		for _, v := range *t {
			values = append(values, v.Value)
		}
	}
	return values
}

// GetValueByName 通过名称获取值
func (t *NameValueSlice) GetValueByName(name string) (string, bool) {
	if t == nil {
		return "", false
	}
	for _, nv := range *t {
		if nv.Name == name {
			return nv.Value, true
		}
	}
	return "", false
}
