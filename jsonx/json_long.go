package jsonx

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Long struct {
	Int64 *int64
}

// NewLong returns a new Long
func NewLong[E constraints.Integer](value E) Long {
	v := int64(value)
	return Long{Int64: &v}
}

// Int64Def 获取值
func (l Long) Int64Def() int64 {
	if l.Int64 == nil {
		return 0
	} else {
		return *l.Int64
	}
}

// Scan implements the Scanner interface.
func (l *Long) Scan(value interface{}) error {
	// if value == nil {
	// 	*l = NewLong(0)
	// 	return nil
	// }
	switch v := value.(type) {
	case int64:
		*l = NewLong(v)
		return nil
	case int:
		*l = NewLong(v)
		return nil
	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return err
		}
		*l = NewLong(i)
		return nil
	default:
		// fmt.Printf("---- %v.%T\n", v, v)
		return nil
	}
}

// Value implements the driver Valuer interface.
func (l Long) Value() (driver.Value, error) {
	if l.Int64 != nil {
		return *l.Int64, nil
	} else {
		return nil, nil
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (l Long) MarshalJSON() ([]byte, error) {
	if l.Int64 != nil {
		return []byte(fmt.Sprintf(`"%v"`, *l.Int64)), nil
	} else {
		return []byte("null"), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *Long) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Trim(str, "\"")
	str = strings.Trim(str, " ")
	if str == "" || str == "null" {
		return nil
	}
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		// strconv.ParseInt: parsing "asdfasdfasd": invalid syntax
		// return fmt.Errorf("'%s' must be numeric", str)
		return err
	}
	l.Int64 = &value
	return nil
}
