package bean

import (
	"errors"
	"fmt"
	"reflect"
)

// copyObj 将数据源复制到目标对象
// @source 数据源对象
// @dest 目标对象
func copyObj(source, dest interface{}) error {
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
	// 设置结构体中相同属性的值
	for i := 0; i < sourceValue.NumField(); i++ {
		fieldName := sourceValue.Type().Field(i).Name
		if ok := destValue.FieldByName(fieldName).IsValid(); ok {
			destValue.FieldByName(fieldName).Set(reflect.ValueOf(sourceValue.Field(i).Interface()))
		}
	}
	return nil
}

// copyObj 将数据源复制到目标对象
// @source 数据源对象
// @dest 目标对象
func copyField(source, dest interface{}, fields ...string) error {
	for i, field := range fields {
		fmt.Println(i, field)
	}

	if len(fields) > 0 {

	}

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
	// 设置结构体中相同属性的值
	for i := 0; i < sourceValue.NumField(); i++ {
		fieldName := sourceValue.Type().Field(i).Name
		if ok := destValue.FieldByName(fieldName).IsValid(); ok {
			destValue.FieldByName(fieldName).Set(reflect.ValueOf(sourceValue.Field(i).Interface()))
		}
	}
	return nil
}

// Copy 将数据源复制到目标对象
// @source 数据源对象
// @dest 目标对象
func Copy(source, dest interface{}) error {
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
		fmt.Println(reflect.TypeOf(destValueSlice)) // []reflect.Value
		for i := 0; i < sourceValue.Len(); i++ {
			sourceItemValue := sourceValue.Index(i)
			//if sourceItemValue.Kind() != reflect.Ptr {
			//	sourceItemValue = sourceItemValue.Addr()
			//}
			destItemValue := reflect.New(destItemType)
			copyObj(sourceItemValue.Interface(), destItemValue.Interface())
			if !isPointer {
				destItemValue = destItemValue.Elem()
			}
			destValueSlice = append(destValueSlice, destItemValue)
		}
		destValueTemp := reflect.Append(destValue, destValueSlice...)
		destValue.Set(destValueTemp)
	case reflect.Struct:
		copyObj(sourceValue.Interface(), dest)
	}
	return errors.New("source type invalid")
}

func valuePrint(value reflect.Value) {
	fmt.Println("value:", value)
	fmt.Println("type:", value.Type())
	fmt.Println("type:", value.IsNil())
	fmt.Println("canSet:", value.CanSet())
	fmt.Printf("%p\n", value.Interface())
}

//func GetUnexportedField(field reflect.Value) interface{} {
//	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
//}
//
//func SetUnexportedField(field reflect.Value, value interface{}) {
//	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
//}
