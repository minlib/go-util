package jsonx

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	formatDateTime = "2006-01-02 15:04:05"
)

type DateTime struct {
	time.Time
}

func NewDateTime(t time.Time) *DateTime {
	return &DateTime{t}
}

// Scan implements the Scanner interface.
func (t *DateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = DateTime{Time: value}
		return nil
	} else {
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
}

// Value insert timestamp into mysql need this function.
func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.Unix() == zeroTime.Unix() {
		return nil, nil
	} else {
		return t.Time, nil
	}
}

// MarshalJSON implements the json.Marshaler interface.
// MarshalJSON on DateTime format Time field with %Y-%m-%d %H:%M:%S
func (t DateTime) MarshalJSON() ([]byte, error) {
	var zeroTime time.Time
	if t.Time.Unix() == zeroTime.Unix() {
		return []byte("null"), nil
	} else {
		formatted := fmt.Sprintf("\"%s\"", t.Format(formatDateTime))
		return []byte(formatted), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	(*t).Time, err = time.ParseInLocation(`"`+formatDateTime+`"`, string(data), time.Local)
	return err
}
