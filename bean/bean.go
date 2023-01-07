package bean

import (
	"errors"
	"reflect"
)

// CopyTo 源对象转为泛型指定的目标对象
// @source 源对象
// @fields 复制的字段，默认复制所有相同的字段
func CopyTo[E interface{}](source interface{}, fields ...string) (E, error) {
	var target E
	err := Copy(source, &target, fields...)
	return target, err
}

// Copy 源对象转为目标对象
// @source 源对象
// @target 目标对象
// @fields 复制的字段，默认复制所有相同的字段
func Copy(source, target interface{}, fields ...string) error {
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
		targetValue := reflect.ValueOf(target)
		if targetValue.Kind() != reflect.Ptr {
			return errors.New("target value can't a pointer type")
		}
		for targetValue.Kind() == reflect.Ptr {
			//if targetValue.IsNil() && targetValue.CanSet() {
			//	targetValue.Set(reflect.New(targetValue.Type().Elem()))
			//}
			targetValue = targetValue.Elem()
		}
		// 切片中项的类型
		targetItemType := targetValue.Type().Elem()
		isPointer := targetItemType.Kind() == reflect.Ptr
		if isPointer {
			targetItemType = targetItemType.Elem()
		}
		targetValueSlice := make([]reflect.Value, 0)
		for i := 0; i < sourceValue.Len(); i++ {
			sourceItemValue := sourceValue.Index(i)
			//if sourceItemValue.Kind() != reflect.Ptr {
			//	sourceItemValue = sourceItemValue.Addr()
			//}
			targetItemValue := reflect.New(targetItemType)
			copyObj(sourceItemValue.Interface(), targetItemValue.Interface(), fields...)
			if !isPointer {
				targetItemValue = targetItemValue.Elem()
			}
			targetValueSlice = append(targetValueSlice, targetItemValue)
		}
		targetValueTemp := reflect.Append(targetValue, targetValueSlice...)
		targetValue.Set(targetValueTemp)
	case reflect.Struct:
		copyObj(sourceValue.Interface(), target, fields...)
	default:
		return errors.New("source type invalid")
	}
	return nil
}

func copyObj(source, target interface{}, fields ...string) error {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)
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
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("target value can't a pointer type")
	}
	if targetValue.IsNil() {
		return errors.New("target value can't be nil")
	}
	for targetValue.Kind() == reflect.Ptr {
		if targetValue.IsNil() && targetValue.CanSet() {
			targetValue.Set(reflect.New(targetValue.Type().Elem()))
		}
		targetValue = targetValue.Elem()
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
			if ok := targetValue.FieldByName(fieldName).IsValid(); ok {
				targetValue.FieldByName(fieldName).Set(reflect.ValueOf(sourceValue.FieldByName(fieldName).Interface()))
			}
		}
	}
	return nil
}
