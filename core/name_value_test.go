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
