package bean

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

// Copy 将数据源复制到目标对象
// @source 数据源对象
// @dest 目标对象
func Copy(source, dest interface{}) error {
	sourceValue := reflect.ValueOf(source)
	//targetType := reflect.TypeOf(dest)
	//if !destValue.CanSet() {
	//	return errors.New("dest value can't be set")
	//}
	//fmt.Println("sourceType:", sourceType)   // []*bean.FruitA
	//fmt.Println("sourceValue:", sourceValue) // [0xc000058740 0xc000058780 ... 0xc000058980 0xc0000589c0]
	//fmt.Println("targetType:", targetType)   // []*bean.FruitB
	//fmt.Println("destValue:", destValue) // []
	//fmt.Println("TypeType:", reflect.TypeOf(targetType)) // *reflect.rtype
	switch sourceValue.Type().Kind() {
	case reflect.Array, reflect.Slice:
		destValue := reflect.ValueOf(dest)
		if destValue.Kind() != reflect.Ptr {
			return errors.New("dest value can't a pointer type")
		}
		for sourceValue.Kind() == reflect.Ptr {
			sourceValue = sourceValue.Elem()
		}
		if !sourceValue.IsValid() {
			return errors.New("source value invalid")
		}
		for destValue.Kind() == reflect.Ptr {
			destValue = destValue.Elem()
		}
		if !destValue.IsValid() {
			return errors.New("dest value invalid")
		}
		destItemType := destValue.Type().Elem()
		fmt.Println(destItemType) // *bean.FruitB
		ptrLevel := 0
		for destItemType.Kind() == reflect.Ptr {
			ptrLevel++
			destItemType = destItemType.Elem()
			fmt.Println(destItemType) // bean.FruitB
		}
		//targetTypeSlice := reflect.MakeSlice(targetType, 0, 0)
		//fmt.Println("targetTypeSlice:", reflect.TypeOf(targetTypeSlice))             // reflect.Value
		//fmt.Println("targetTypeSlice:", reflect.TypeOf(targetTypeSlice.Interface())) // []*bean.FruitB
		destValueSlice := make([]reflect.Value, 0)
		fmt.Println(reflect.TypeOf(destValueSlice)) // []reflect.Value
		for i := 0; i < sourceValue.Len(); i++ {
			sourceItemValue := sourceValue.Index(i)
			//if sourceItemValue.Kind() != reflect.Ptr {
			//	sourceItemValue = sourceItemValue.Addr()
			//}
			destItemValue := reflect.New(destItemType)
			copyObj(sourceItemValue.Interface(), destItemValue.Interface())
			if ptrLevel == 0 {
				destItemValue = destItemValue.Elem()
			}
			destValueSlice = append(destValueSlice, destItemValue)

			//fmt.Println(sourceItemValue)                              // &{1 名称1 81 0.5372820936538049 2022-05-14 23:08:55.776294 +0  800 CST m=+0.017055401}
			//fmt.Println(sourceItemValue.Interface())                  // &{1 名称1 81 0.5372820936538049 2022-05-14 23:08:55.776294 +080  0 CST m=+0.017055401}
			//fmt.Println(destItemValue)                                // &{1 名称1 81 0.5372820936538049 2022-05-14 23:08:55.776294 +0  800 CST m=+0.017055401 0001-01-01 00:00:00 +0000 UTC }
			//fmt.Println(destItemValue.Interface())                    // &{1 名称1 81 0.5372820936538049 2022-05-14 23:08:55.776294 +0  800 CST m=+0.017055401 0001-01-01 00:00:00 +0000 UTC }
			//fmt.Println(destItemValue.Elem())                         // {1 名称1 81 0.5372820936538049 2022-05-14 23:17:06.4145388 +0800 C  ST m=+0.017129201 0001-01-01 00:00:00 +0000 UTC }
			//fmt.Println(destItemValue.Elem().Interface())             // {1 名称1 81 0.5372820936538049 2022-05-14 23:17:06.4145388 +0800 C  ST m=+0.017129201 0001-01-01 00:00:00 +0000 UTC }
			//fmt.Println(reflect.TypeOf(sourceItemValue.Interface()))  // *bean.FruitA
			//fmt.Println(reflect.ValueOf(sourceItemValue.Interface())) // &{1 名称1 81 0.5372820936538049 2022-05-14 15:20:45.6038055 +0800 CST m=+0.003174001}
			//fmt.Println(reflect.TypeOf(destItemValue.Interface()))    // *bean.FruitB
			//fmt.Println(reflect.ValueOf(destItemValue.Interface()))   // &{1 名称1 81 0.5372820936538049 2022-05-14 15:19:40.1222062 +0800 CST m=+0.002609001 0001-01-01 00:00:00 +0000 UTC }
			//marshal, _ := json.Marshal(destItemValue.Interface())
			//fmt.Println(string(marshal))
		}
		destValueTemp := reflect.Append(destValue, destValueSlice...)
		destValue.Set(destValueTemp)
	case reflect.Struct:
		copyObj(source, dest)
	}
	return errors.New("source type invalid")
}

// copyObj 将数据源复制到目标对象
// @source 数据源对象
// @dest 目标对象
func copyObj(source, dest interface{}) error {
	sourceValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr {
		return errors.New("dest value can't a pointer type")
	}
	for sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}
	if !sourceValue.IsValid() {
		return errors.New("source value invalid")
	}
	for destValue.Kind() == reflect.Ptr {
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

func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}
