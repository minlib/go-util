package bean

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

// Copy 将数据源对象复制到目标对象
// @source 数据源对象
// @target 目标对象
func Copy(source, target interface{}) {
	//获取reflect.Type类型
	//sourceType := reflect.TypeOf(source)
	//sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	sourceType := sourceValue.Type()
	for i := 0; i < sourceValue.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := sourceType.Field(i).Name
		if ok := targetValue.FieldByName(name).IsValid(); ok {
			targetValue.FieldByName(name).Set(reflect.ValueOf(sourceValue.Field(i).Interface()))
		}
	}
}

func CopyObj(source, target interface{}) {
	//获取reflect.Type类型
	//sourceType := reflect.TypeOf(source)
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)
	//sourceType := sourceValue.Type()
	for i := 0; i < sourceValue.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := sourceType.Field(i).Name
		if ok := targetValue.FieldByName(name).IsValid(); ok {
			targetValue.FieldByName(name).Set(reflect.ValueOf(sourceValue.Field(i).Interface()))
		}
	}
}

func Copy2(source, target interface{}) {

	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)
	targetType := reflect.TypeOf(target)
	targetValue := reflect.ValueOf(target)

	fmt.Println(sourceValue.CanSet())
	fmt.Println(targetValue.CanSet())

	fmt.Println("sourceType:", sourceType)               // []*bean.fruitA
	fmt.Println("sourceValue:", sourceValue)             // [0xc000058740 0xc000058780 ... 0xc000058980 0xc0000589c0]
	fmt.Println("targetType:", targetType)               // []*bean.fruitB
	fmt.Println("targetValue:", targetValue)             // []
	fmt.Println("TypeType:", reflect.TypeOf(targetType)) // *reflect.rtype
	switch sourceType.Kind() {
	case reflect.Array, reflect.Slice:
		//targetTypeSlice := reflect.MakeSlice(targetType, 0, 0)
		//fmt.Println("targetTypeSlice:", reflect.TypeOf(targetTypeSlice))             // reflect.Value
		//fmt.Println("targetTypeSlice:", reflect.TypeOf(targetTypeSlice.Interface())) // []*bean.fruitB
		values := make([]reflect.Value, 0)
		fmt.Println("values:", reflect.TypeOf(values)) // []reflect.Value
		for i := 0; i < sourceValue.Len(); i++ {
			sour := sourceValue.Index(i)
			fmt.Println(sour)
			fmt.Println(sour.Interface())
			value := reflect.New(targetType.Elem().Elem())
			if value.Type().Kind() == reflect.Ptr {

			}
			fmt.Println(value)
			fmt.Println(value.Interface())

			Copy(sour.Interface(), value.Interface())

			fmt.Println(reflect.TypeOf(sour.Interface()))   // *bean.fruitA
			fmt.Println(reflect.ValueOf(sour.Interface()))  // &{1 名称1 81 0.5372820936538049 2022-05-14 15:20:45.6038055 +0800 CST m=+0.003174001}
			fmt.Println(reflect.TypeOf(value.Interface()))  // *bean.fruitB
			fmt.Println(reflect.ValueOf(value.Interface())) // &{1 名称1 81 0.5372820936538049 2022-05-14 15:19:40.1222062 +0800 CST m=+0.002609001 0001-01-01 00:00:00 +0000 UTC }

			values = append(values, value)
			//values = append(values, reflect.ValueOf(value.Interface()))
			//targetTypeSlice = append(targetTypeSlice, reflect.ValueOf(value.Interface()))

			//ptr := reflect.New(sliceOfT)
			//ptr.Elem().Set(reflect.MakeSlice(sliceOfT, 0, 0))
			//s := ptr.Interface()

			//typeOfT := reflect.TypeOf(t)
			//sliceOfT := reflect.SliceOf(typeOfT)
			//s := reflect.MakeSlice(sliceOfT, 0, 0).Interface()

			//e0 = append(e0, reflect.ValueOf(100))
			//e0 = append(e0, reflect.ValueOf(200))
			//e0 = append(e0, reflect.ValueOf(300))
			//e0 = append(e0, reflect.ValueOf(400))

			//reflect.Append(a0, e0...)

			marshal, _ := json.Marshal(value.Interface())
			fmt.Println(string(marshal))
		}

		fmt.Println("1111111111111111111")
		fmt.Println(values[0])

		c := reflect.Append(targetValue, values...)

		//fmt.Printf("变量的地址: %x\n", c)
		//fmt.Printf("变量的地址: %X\n", target)
		//fmt.Printf("变量的地址: %X\n", targetValue)

		fmt.Println(targetValue.CanSet())
		targetValue.Set(c)

		//targetValue = reflect.Append(targetValue, values...)

		fmt.Println(targetValue)

		//targetValue.Set(c)

		//SetUnexportedField(targetValue, values)

		fmt.Printf("变量的地址: %x\n", c)
		fmt.Printf("变量的地址: %x\n", target)

		fmt.Println(reflect.TypeOf(c))
		fmt.Println(reflect.ValueOf(c))

		fmt.Println(reflect.TypeOf(targetValue))
		fmt.Println(reflect.ValueOf(targetValue))

		fmt.Println(c.Kind())
		fmt.Println(c.IsNil())
		fmt.Println(c.CanAddr())
		fmt.Println(c.IsValid())

		fmt.Println(targetValue.Kind())
		fmt.Println(targetValue.IsNil())
		fmt.Println(targetValue.CanAddr())
		fmt.Println(targetValue.IsValid())

		marshalc, _ := json.Marshal(c.Interface())
		fmt.Println(string(marshalc))

		marshalb, _ := json.Marshal(targetValue.Interface())
		fmt.Println(string(marshalb))

		fmt.Println(c.IsNil())
		fmt.Println(targetValue.IsNil())

		l := []int{1, 2}
		typeOf := reflect.TypeOf(l)
		fmt.Println(typeOf)
		fmt.Println(typeOf.Name())

		//SetUnexportedField(targetTypeSlice, values)
		//targetValue.Set(c)
		//targetTypeSlice.Set(c)

		//targetTypeSlice2 := reflect.MakeSlice(targetType, 0, 0)
		//targetTypeSlice2.Set(c)

		fmt.Println(c)
		//
		//SetUnexportedField(targetTypeSlice, c)

		//fmt.Println(targetValue)

		//fmt.Println(targetValue[0])
		//
		//
		//fmt.Println("222222222222222222222")
		//targetValue.Set(cc)

	case reflect.Ptr:
		Copy2(sourceType.Elem(), targetValue)
	case reflect.Struct:
		for i := 0; i < sourceValue.NumField(); i++ {
			// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
			name := sourceType.Field(i).Name
			if ok := targetValue.FieldByName(name).IsValid(); ok {
				targetValue.FieldByName(name).Set(reflect.ValueOf(sourceValue.Field(i).Interface()))
			}
		}
	}

	////获取reflect.Type类型
	//targetValue := reflect.ValueOf(target).Elem()
	//sourceValue := reflect.ValueOf(source).Elem()
	////sourceType := sourceValue.Type()
	//for i := 0; i < sourceValue.NumField(); i++ {
	//	// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
	//	name := sourceType.Field(i).Name
	//	if ok := targetValue.FieldByName(name).IsValid(); ok {
	//		targetValue.FieldByName(name).Set(reflect.ValueOf(sourceValue.Field(i).Interface()))
	//	}
	//}
}

func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}
