package core

import (
	"database/sql/driver"
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

type Int struct {
	Int32 *int32
}

// NewInt returns a new Int
func NewInt[E constraints.Integer](value E) Int {
	v := int32(value)
	return Int{Int32: &v}
}

// Int32Def 获取值
func (l Int) Int32Def() int32 {
	if l.Int32 == nil {
		return 0
	} else {
		return *l.Int32
	}
}

// Scan implements the Scanner interface.
func (l *Int) Scan(value interface{}) error {
	// if value == nil {
	// 	*l = NewInt(0)
	// 	return nil
	// }
	switch v := value.(type) {
	case int32:
		*l = NewInt(v)
		return nil
	case int:
		*l = NewInt(v)
		return nil
	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return err
		}
		*l = NewInt(i)
		return nil
	default:
		// fmt.Printf("---- %v.%T\n", v, v)
		return nil
	}
}

// Value implements the driver Valuer interface.
func (l Int) Value() (driver.Value, error) {
	if l.Int32 != nil {
		return *l.Int32, nil
	} else {
		return nil, nil
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (l Int) MarshalJSON() ([]byte, error) {
	if l.Int32 != nil {
		return []byte(fmt.Sprintf(`%v`, *l.Int32)), nil
	} else {
		return []byte("null"), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *Int) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Trim(str, "\"")
	str = strings.Trim(str, " ")
	if str == "" || str == "null" {
		return nil
	}
	value, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		// strconv.ParseInt: parsing "asdfasdfasd": invalid syntax
		// return fmt.Errorf("'%s' must be numeric", str)
		return err
	}
	v := int32(value)
	l.Int32 = &v
	return nil
}
