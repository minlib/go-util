package core

import (
	"database/sql/driver"
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

type Integer struct {
	Int32 *int32
}

// NewInteger returns a new Integer
func NewInteger[E constraints.Integer](value E) Integer {
	v := int32(value)
	return Integer{Int32: &v}
}

// Int32Def 获取值
func (i Integer) Int32Def() int32 {
	if i.Int32 == nil {
		return 0
	} else {
		return *i.Int32
	}
}

// Scan implements the Scanner interface.
func (i *Integer) Scan(value interface{}) error {
	switch v := value.(type) {
	case int32:
		*i = NewInteger(v)
		return nil
	case int:
		*i = NewInteger(v)
		return nil
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return err
		}
		*i = NewInteger(val)
		return nil
	default:
		return nil
	}
}

// Value implements the driver Valuer interface.
func (i Integer) Value() (driver.Value, error) {
	if i.Int32 != nil {
		return *i.Int32, nil
	} else {
		return nil, nil
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (i Integer) MarshalJSON() ([]byte, error) {
	if i.Int32 != nil {
		return []byte(fmt.Sprintf(`%v`, *i.Int32)), nil
	} else {
		return []byte("null"), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Integer) UnmarshalJSON(data []byte) error {
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
	i.Int32 = &v
	return nil
}

func (i Integer) String() string {
	if i.Int32 == nil {
		return ""
	}
	return strconv.FormatInt(int64(*i.Int32), 10)
}
