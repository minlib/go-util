package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Marshal(t *testing.T) {
	var list NameValueSlice
	list = append(list, NameValue{
		Name:  "张三",
		Value: "18",
	})
	list = append(list, NameValue{
		Name:  "李四",
		Value: "22",
	})

	bytes, _ := json.Marshal(&list)
	fmt.Println(string(bytes))

	var list2 NameValueSlice
	bytes2, _ := json.Marshal(list2)
	fmt.Println(string(bytes2))
}

func Test_Unmarshal(t *testing.T) {
	var list NameValueSlice
	list = append(list, NameValue{
		Name:  "张三",
		Value: "18",
	})
	list = append(list, NameValue{
		Name:  "李四",
		Value: "22",
	})

	bytes, _ := json.Marshal(&list)

	var list2 NameValueSlice
	_ = json.Unmarshal(bytes, &list2)

	var list3 NameValueSlice
	_ = json.Unmarshal(nil, &list3)
}

func Test_GetValueByName(t *testing.T) {
	var list NameValueSlice
	list = append(list, NameValue{
		Name:  "张三",
		Value: "18",
	})
	list = append(list, NameValue{
		Name:  "李四",
		Value: "22",
	})
	value, ok := list.GetValueByName("张三")
	fmt.Println(value, ok)
}

func Test_String(t *testing.T) {
	var list NameValueSlice
	list = append(list, NameValue{
		Name:  "张三",
		Value: "18",
	})
	list = append(list, NameValue{
		Name:  "李四",
		Value: "22",
	})
	fmt.Println(list)
	fmt.Println(list.String())

	var list2 NameValueSlice
	fmt.Println(list2)
	fmt.Println(list2.String())

	var list3 = &list
	fmt.Println(list3)
	fmt.Println(list3.String())

	var list4 *NameValueSlice
	fmt.Println(list4)
}
