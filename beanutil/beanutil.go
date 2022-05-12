package beanutil

import (
	"reflect"
)

// Copy 将数据源对象复制到目标对象
// @source 数据源对象
// @target 目标对象
func Copy(source, target interface{}) {
	//获取reflect.Type类型
	targetValue := reflect.ValueOf(target).Elem()
	sourceValue := reflect.ValueOf(source).Elem()
	sourceType := sourceValue.Type()
	for i := 0; i < sourceValue.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := sourceType.Field(i).Name
		if ok := targetValue.FieldByName(name).IsValid(); ok {
			targetValue.FieldByName(name).Set(reflect.ValueOf(sourceValue.Field(i).Interface()))
		}
	}
}
