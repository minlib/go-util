package json

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type LongSlice []int64

func (slice LongSlice) MarshalJSON() ([]byte, error) {
	values := make([]string, len(slice))
	for i, value := range []int64(slice) {
		values[i] = fmt.Sprintf(`"%v"`, value)
	}
	return []byte(fmt.Sprintf("[%v]", strings.Join(values, ","))), nil
}

func (slice *LongSlice) UnmarshalJSON(b []byte) error {
	// Try array of strings first.
	var values []string
	err := json.Unmarshal(b, &values)
	if err != nil {
		// Fall back to array of integers:
		var values []int64
		if err := json.Unmarshal(b, &values); err != nil {
			return err
		}
		*slice = values
		return nil
	}
	*slice = make([]int64, len(values))
	for i, value := range values {
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		(*slice)[i] = value
	}
	return nil
}
