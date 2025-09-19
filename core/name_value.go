package core

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NameValueSlice []NameValue

// Scan implements the Scanner interface.
func (s *NameValueSlice) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &s)
}

// Value implements the driver Valuer interface.
func (s *NameValueSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// GetNames get name list
func (s *NameValueSlice) GetNames() []string {
	var names []string
	if s != nil {
		for _, v := range *s {
			names = append(names, v.Name)
		}
	}
	return names
}

// GetValues get value list
func (s *NameValueSlice) GetValues() []string {
	var values []string
	if s != nil {
		for _, v := range *s {
			values = append(values, v.Value)
		}
	}
	return values
}

// GetValueByName 通过名称获取值
func (s *NameValueSlice) GetValueByName(name string) (string, bool) {
	if s == nil {
		return "", false
	}
	for _, nv := range *s {
		if nv.Name == name {
			return nv.Value, true
		}
	}
	return "", false
}

func (s NameValueSlice) String() string {
	var values []string
	for _, v := range s {
		values = append(values, v.Name+":"+v.Value)
	}
	return "[" + strings.Join(values, ",") + "]"
}
