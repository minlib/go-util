package json

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Long struct {
	Int64 int64
	Valid bool
}

// NewLong returns a new Long
func NewLong[E constraints.Integer](value E) *Long {
	return &Long{int64(value), true}
}

// NewLongSlice returns a new Long slice
func NewLongSlice[E constraints.Integer](values []E) []*Long {
	var ls []*Long
	for _, v := range values {
		l := NewLong(v)
		ls = append(ls, l)
	}
	return ls
}

// Scan implements the Scanner interface.
func (l *Long) Scan(value interface{}) error {
	if value == nil {
		l.Int64, l.Valid = 0, false
		return nil
	}
	switch v := value.(type) {
	case int64:
		l.Int64, l.Valid = v, true
		return nil
	case int:
		l.Int64, l.Valid = int64(v), true
		return nil
	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return err
		}
		l.Int64, l.Valid = i, true
		return nil
	default:
		// fmt.Printf("---- %v.%T\n", v, v)
		return nil
	}
}

// Value implements the driver Valuer interface.
func (l Long) Value() (driver.Value, error) {
	if !l.Valid {
		return nil, nil
	} else {
		return l.Int64, nil
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (l Long) MarshalJSON() ([]byte, error) {
	if !l.Valid {
		return nil, nil
	} else {
		return []byte(fmt.Sprintf(`"%v"`, l.Int64)), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *Long) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	str := string(data)
	str = strings.Trim(str, "\"")
	str = strings.Trim(str, " ")
	if str == "" || str == "null" {
		//l.Valid = false
		return nil
	}
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		// strconv.ParseInt: parsing "asdfasdfasd": invalid syntax
		// return fmt.Errorf("'%s' must be numeric", str)
		return err
	}
	l.Valid = true
	(*l).Int64 = value
	return nil
}
