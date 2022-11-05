package bean

import (
	"errors"
	"reflect"
)

// copyObj 将数据源复制到目标对象
// @source 数据源对象
// @dest 目标对象
// @fields 复制的字段名，默认复制全部相同的字段
func copyObj(source, dest interface{}, fields ...string) error {
	sourceValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(dest)
	if !sourceValue.IsValid() {
		return errors.New("source value invalid")
	}
	if sourceValue.Kind() == reflect.Ptr {
		if sourceValue.IsNil() {
			return errors.New("source value can't nil")
		}
		for sourceValue.Kind() == reflect.Ptr {
			sourceValue = sourceValue.Elem()
		}
	}
	if destValue.Kind() != reflect.Ptr {
		return errors.New("dest value can't a pointer type")
	}
	if destValue.IsNil() {
		return errors.New("dest value can't be nil")
	}
	for destValue.Kind() == reflect.Ptr {
		if destValue.IsNil() && destValue.CanSet() {
			destValue.Set(reflect.New(destValue.Type().Elem()))
		}
		destValue = destValue.Elem()
	}
	// 获取设置的属性列表
	hasFieldAll := len(fields) == 0 // 不设置默认全部字段
	fieldMap := make(map[string]struct{}, len(fields))
	for _, fieldName := range fields {
		fieldMap[fieldName] = struct{}{}
	}
	for i := 0; i < sourceValue.NumField(); i++ {
		fieldName := sourceValue.Type().Field(i).Name
		hasField := hasFieldAll
		if !hasFieldAll {
			if _, found := fieldMap[fieldName]; found {
				hasField = true
			}
		}
		if hasField {
			if ok := destValue.FieldByName(fieldName).IsValid(); ok {
				destValue.FieldByName(fieldName).Set(reflect.ValueOf(sourceValue.FieldByName(fieldName).Interface()))
			}
		}
	}
	return nil
}

// Copy 将数据源复制到目标对象
// @source 数据源对象
// @dest 目标对象
// @fields 复制的字段名，默认复制全部相同的字段
func Copy(source, dest interface{}, fields ...string) error {
	sourceValue := reflect.ValueOf(source)
	if !sourceValue.IsValid() {
		return errors.New("source value invalid")
	}
	if sourceValue.Kind() == reflect.Ptr {
		if sourceValue.IsNil() {
			return errors.New("source value can't nil")
		}
		for sourceValue.Kind() == reflect.Ptr {
			sourceValue = sourceValue.Elem()
		}
	}
	switch sourceValue.Type().Kind() {
	case reflect.Array, reflect.Slice:
		destValue := reflect.ValueOf(dest)
		if destValue.Kind() != reflect.Ptr {
			return errors.New("dest value can't a pointer type")
		}
		for destValue.Kind() == reflect.Ptr {
			//if destValue.IsNil() && destValue.CanSet() {
			//	destValue.Set(reflect.New(destValue.Type().Elem()))
			//}
			destValue = destValue.Elem()
		}
		// 切片中项的类型
		destItemType := destValue.Type().Elem()
		isPointer := destItemType.Kind() == reflect.Ptr
		if isPointer {
			destItemType = destItemType.Elem()
		}
		destValueSlice := make([]reflect.Value, 0)
		for i := 0; i < sourceValue.Len(); i++ {
			sourceItemValue := sourceValue.Index(i)
			//if sourceItemValue.Kind() != reflect.Ptr {
			//	sourceItemValue = sourceItemValue.Addr()
			//}
			destItemValue := reflect.New(destItemType)
			copyObj(sourceItemValue.Interface(), destItemValue.Interface(), fields...)
			if !isPointer {
				destItemValue = destItemValue.Elem()
			}
			destValueSlice = append(destValueSlice, destItemValue)
		}
		destValueTemp := reflect.Append(destValue, destValueSlice...)
		destValue.Set(destValueTemp)
	case reflect.Struct:
		copyObj(sourceValue.Interface(), dest, fields...)
	}
	return errors.New("source type invalid")
}
