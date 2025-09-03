package mapx

import (
	"reflect"
	"testing"
)

func TestContainsKey(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		k    string
		want bool
	}{
		{
			name: "nil map",
			m:    nil,
			k:    "a",
			want: false,
		},
		{
			name: "key exists",
			m:    map[string]int{"a": 1},
			k:    "a",
			want: true,
		},
		{
			name: "key not exists",
			m:    map[string]int{"a": 1},
			k:    "b",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsKey(tt.m, tt.k); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	m := map[string]interface{}{"name": "test", "age": 20}

	val, exists := Get(m, "name")
	if !exists || val != "test" {
		t.Error("Get() failed for existing key")
	}

	val2, exists2 := Get(m, "nonExist")
	if exists2 || val2 != nil {
		t.Error("Get() failed for non-existing key")
	}

	val3, exists3 := Get(nil, "name")
	if exists3 || val3 != nil {
		t.Error("Get() failed for nil map")
	}
}

func TestGetString(t *testing.T) {
	tests := []struct {
		name         string
		inputMap     map[string]interface{}
		key          string
		defaultValue string
		expected     string
	}{
		{
			name:         "get existing string",
			inputMap:     map[string]interface{}{"name": "test"},
			key:          "name",
			defaultValue: "default",
			expected:     "test",
		},
		{
			name:         "get int as string",
			inputMap:     map[string]interface{}{"age": 20},
			key:          "age",
			defaultValue: "0",
			expected:     "20",
		},
		{
			name:         "get int64 as string",
			inputMap:     map[string]interface{}{"id": int64(123456)},
			key:          "id",
			defaultValue: "0",
			expected:     "123456",
		},
		{
			name:         "get float64 as string",
			inputMap:     map[string]interface{}{"price": 99.99},
			key:          "price",
			defaultValue: "0",
			expected:     "99.99",
		},
		{
			name:         "get bool as string",
			inputMap:     map[string]interface{}{"active": true},
			key:          "active",
			defaultValue: "false",
			expected:     "true",
		},
		{
			name:         "get non-existing key",
			inputMap:     map[string]interface{}{"name": "test"},
			key:          "nonexist",
			defaultValue: "default",
			expected:     "default",
		},
		{
			name:         "handle nil map",
			inputMap:     nil,
			key:          "name",
			defaultValue: "default",
			expected:     "default",
		},
		{
			name:         "handle nil value",
			inputMap:     map[string]interface{}{"data": nil},
			key:          "data",
			defaultValue: "default",
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetString(tt.inputMap, tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name         string
		inputMap     map[string]interface{}
		key          string
		defaultValue int
		expected     int
	}{
		{
			name:         "get int",
			inputMap:     map[string]interface{}{"age": 20},
			key:          "age",
			defaultValue: 0,
			expected:     20,
		},
		{
			name:         "get int64",
			inputMap:     map[string]interface{}{"count": int64(100)},
			key:          "count",
			defaultValue: 0,
			expected:     100,
		},
		{
			name:         "get uint",
			inputMap:     map[string]interface{}{"size": uint(50)},
			key:          "size",
			defaultValue: 0,
			expected:     50,
		},
		{
			name:         "get float64 as int",
			inputMap:     map[string]interface{}{"score": 95.8},
			key:          "score",
			defaultValue: 0,
			expected:     95,
		},
		{
			name:         "get string number",
			inputMap:     map[string]interface{}{"page": "5"},
			key:          "page",
			defaultValue: 0,
			expected:     5,
		},
		{
			name:         "invalid string to int",
			inputMap:     map[string]interface{}{"name": "test"},
			key:          "name",
			defaultValue: -1,
			expected:     -1,
		},
		{
			name:         "non-existing key",
			inputMap:     map[string]interface{}{"age": 20},
			key:          "nonexist",
			defaultValue: -1,
			expected:     -1,
		},
		{
			name:         "handle nil map",
			inputMap:     nil,
			key:          "age",
			defaultValue: 0,
			expected:     0,
		},
		{
			name:         "handle nil value",
			inputMap:     map[string]interface{}{"data": nil},
			key:          "data",
			defaultValue: 0,
			expected:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetInt(tt.inputMap, tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetInt() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetFloat(t *testing.T) {
	// 辅助函数：容差比较浮点数
	almostEqual := func(a, b float64) bool {
		const epsilon = 0.000001
		diff := a - b
		if diff < 0 {
			diff = -diff
		}
		return diff < epsilon
	}

	tests := []struct {
		name         string
		m            map[string]interface{}
		key          string
		defaultValue float64
		want         float64
	}{
		{
			name:         "nil map",
			m:            nil,
			key:          "a",
			defaultValue: 0.0,
			want:         0.0,
		},
		{
			name:         "key not exists",
			m:            map[string]interface{}{"b": 3.14},
			key:          "a",
			defaultValue: 2.718,
			want:         2.718,
		},
		{
			name:         "value is nil",
			m:            map[string]interface{}{"a": nil},
			key:          "a",
			defaultValue: 1.0,
			want:         1.0,
		},
		{
			name:         "float32 type",
			m:            map[string]interface{}{"g": float32(12.34)},
			key:          "g",
			defaultValue: 0.0,
			want:         12.34,
		},
		{
			name:         "float64 type",
			m:            map[string]interface{}{"a": 3.1415926},
			key:          "a",
			defaultValue: 0.0,
			want:         3.1415926,
		},
		{
			name:         "int type (convert to float)",
			m:            map[string]interface{}{"a": 100},
			key:          "a",
			defaultValue: 0.0,
			want:         100.0,
		},
		{
			name:         "string parseable to float",
			m:            map[string]interface{}{"a": "123.45"},
			key:          "a",
			defaultValue: 0.0,
			want:         123.45,
		},
		{
			name:         "string unparseable to float",
			m:            map[string]interface{}{"a": "not_a_float"},
			key:          "a",
			defaultValue: 5.5,
			want:         5.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFloat(tt.m, tt.key, tt.defaultValue)
			if !almostEqual(got, tt.want) {
				t.Errorf("GetFloat() = %v, want %v (difference: %v)", got, tt.want, got-tt.want)
			}
		})
	}
}

func TestGetStringSlice(t *testing.T) {
	tests := []struct {
		name         string
		inputMap     map[string]interface{}
		key          string
		defaultValue []string
		expected     []string
	}{
		{
			name: "get string slice",
			inputMap: map[string]interface{}{
				"tags": []interface{}{"go", "json", "util"},
			},
			key:          "tags",
			defaultValue: []string{},
			expected:     []string{"go", "json", "util"},
		},
		{
			name: "get mixed type slice",
			inputMap: map[string]interface{}{
				"values": []interface{}{10, "20", 30.5, true},
			},
			key:          "values",
			defaultValue: []string{},
			expected:     []string{"10", "20", "30.5", "true"},
		},
		{
			name:         "non-existing key",
			inputMap:     map[string]interface{}{"name": "test"},
			key:          "tags",
			defaultValue: []string{"default"},
			expected:     []string{"default"},
		},
		{
			name:         "not a slice type",
			inputMap:     map[string]interface{}{"tags": "not a slice"},
			key:          "tags",
			defaultValue: []string{"default"},
			expected:     []string{"default"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetStringSlice(tt.inputMap, tt.key, tt.defaultValue)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GetStringSlice() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 辅助函数：判断浮点数是否在容差范围内相等
func almostEqual(a, b float64) bool {
	const epsilon = 1e-9
	if a == b {
		return true
	}
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < epsilon
}
