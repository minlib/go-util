package mapx

import (
	"fmt"
	"reflect"
	"strconv"
)

// ContainsKey returns the keys of the map m contains a key k.
func ContainsKey[M ~map[K]V, K comparable, V any](m M, k K) bool {
	if m == nil {
		return false
	}
	_, found := m[k]
	return found
}

// RemoveKeys 批量移除map中的指定key
// 参数:
//
//	m: 要操作的map，允许为nil（此时无操作）
//	keys: 要移除的key列表
//
// 注意:
//  1. 会直接修改原map，而非返回新map
//  2. 对于不存在的key会自动忽略
//  3. 支持任意可比较类型的key和任意值类型的map
func RemoveKeys[M ~map[K]V, K comparable, V any](m M, keys ...K) {
	if m == nil {
		return
	}
	for _, key := range keys {
		delete(m, key)
	}
}

// Get 通用获取函数，返回原始值和是否存在的标志
// 用于需要处理更复杂类型的场景
func Get(m map[string]interface{}, key string) (interface{}, bool) {
	if m == nil {
		return nil, false
	}
	val, exists := m[key]
	return val, exists
}

// GetString 从map中安全获取字符串值
// 参数:
//
//	m: 源map，允许为nil
//	key: 要获取的键
//	defaultValue: 当键不存在或转换失败时返回的默认值
//
// 返回:
//
//	成功获取的字符串或默认值
func GetString(m map[string]interface{}, key string, defaultValue string) string {
	// 处理nil map
	if m == nil {
		return defaultValue
	}

	// 获取原始值
	val, exists := m[key]
	if !exists {
		return defaultValue
	}

	// 空值处理
	if val == nil {
		return defaultValue
	}

	// 直接返回字符串类型
	if str, ok := val.(string); ok {
		return str
	}

	// 处理数字类型转换
	switch num := val.(type) {
	case int:
		return strconv.Itoa(num)
	case int8:
		return strconv.FormatInt(int64(num), 10)
	case int16:
		return strconv.FormatInt(int64(num), 10)
	case int32:
		return strconv.FormatInt(int64(num), 10)
	case int64:
		return strconv.FormatInt(num, 10)
	case uint:
		return strconv.FormatUint(uint64(num), 10)
	case uint8:
		return strconv.FormatUint(uint64(num), 10)
	case uint16:
		return strconv.FormatUint(uint64(num), 10)
	case uint32:
		return strconv.FormatUint(uint64(num), 10)
	case uint64:
		return strconv.FormatUint(num, 10)
	case float32:
		return strconv.FormatFloat(float64(num), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(num, 'f', -1, 64)
	}

	// 处理布尔类型
	if b, ok := val.(bool); ok {
		return strconv.FormatBool(b)
	}

	// 其他类型使用默认格式转换
	return fmt.Sprintf("%v", val)
}

// GetInt 从map中安全获取整数值
// 参数:
//
//	m: 源map，允许为nil
//	key: 要获取的键
//	defaultValue: 当键不存在或转换失败时返回的默认值
//
// 返回:
//
//	成功获取的整数或默认值
func GetInt(m map[string]interface{}, key string, defaultValue int) int {
	// 处理nil map
	if m == nil {
		return defaultValue
	}

	// 获取原始值
	val, exists := m[key]
	if !exists {
		return defaultValue
	}

	// 空值处理
	if val == nil {
		return defaultValue
	}

	// 处理整数类型
	switch num := val.(type) {
	case int:
		return num
	case int8:
		return int(num)
	case int16:
		return int(num)
	case int32:
		return int(num)
	case int64:
		return int(num)
	case uint:
		return int(num)
	case uint8:
		return int(num)
	case uint16:
		return int(num)
	case uint32:
		return int(num)
	case uint64:
		return int(num)
	}

	// 处理浮点类型（截断小数部分）
	switch num := val.(type) {
	case float32:
		return int(num)
	case float64:
		return int(num)
	}

	// 处理字符串类型（尝试解析为整数）
	if str, ok := val.(string); ok {
		num, err := strconv.Atoi(str)
		if err == nil {
			return num
		}
	}

	// 其他类型返回默认值
	return defaultValue
}

// GetFloat 从map中安全获取浮点数值
// 参数:
//
//	m: 源map，允许为nil
//	key: 要获取的键
//	defaultValue: 当键不存在或转换失败时返回的默认值
//
// 返回:
//
//	成功获取的浮点数或默认值
func GetFloat(m map[string]interface{}, key string, defaultValue float64) float64 {
	// 处理nil map
	if m == nil {
		return defaultValue
	}

	// 获取原始值
	val, exists := m[key]
	if !exists {
		return defaultValue
	}

	// 空值处理
	if val == nil {
		return defaultValue
	}

	// 处理整数类型
	switch num := val.(type) {
	case int:
		return float64(num)
	case int8:
		return float64(num)
	case int16:
		return float64(num)
	case int32:
		return float64(num)
	case int64:
		return float64(num)
	case uint:
		return float64(num)
	case uint8:
		return float64(num)
	case uint16:
		return float64(num)
	case uint32:
		return float64(num)
	case uint64:
		return float64(num)
	}

	// 处理浮点类型
	switch num := val.(type) {
	case float32:
		return float64(num)
	case float64:
		return num
	}

	// 处理字符串类型（尝试解析为浮点数）
	if str, ok := val.(string); ok {
		num, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return num
		}
	}

	// 其他类型返回默认值
	return defaultValue
}

// GetStringSlice 将map中的值转换为字符串切片
func GetStringSlice(m map[string]interface{}, key string, defaultValue []string) []string {
	if m == nil {
		return defaultValue
	}

	val, exists := m[key]
	if !exists {
		return defaultValue
	}

	// 尝试将值转换为切片
	slice, ok := val.([]interface{})
	if !ok {
		return defaultValue
	}

	// 转换每个元素为字符串
	result := make([]string, 0, len(slice))
	for _, item := range slice {
		if item == nil {
			result = append(result, "")
			continue
		}

		// 使用反射获取元素类型并转换为字符串
		switch reflect.TypeOf(item).Kind() {
		case reflect.String:
			result = append(result, item.(string))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			result = append(result, fmt.Sprintf("%v", item))
		case reflect.Bool:
			result = append(result, strconv.FormatBool(item.(bool)))
		default:
			result = append(result, fmt.Sprintf("%v", item))
		}
	}
	return result
}
