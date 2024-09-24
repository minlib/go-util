package jsonx

import (
	"fmt"
	"testing"
)

func TestMarshalString(t *testing.T) {
	fmt.Println(MarshalString([]int64{12, 34, 567}))       // [12,34,567]
	fmt.Println(MarshalString([]string{"A", "BC", "DEF"})) // ["A","BC","DEF"]

	m := make(map[string]any, 2)
	m["age"] = 24
	m["name"] = "张三"
	fmt.Println(MarshalString(m)) // {"age":24,"name":"张三"}
}

func TestUnmarshalString(t *testing.T) {
	i := []int64{}
	err := UnmarshalString("[12,34,567]", &i)
	fmt.Println(i, err) // [12 34 567]

	s := []string{}
	err = UnmarshalString("[\"A\",\"BC\",\"DEF\"]", &s)
	fmt.Println(s, err) // [A BC DEF]

	m := make(map[string]any, 0)
	err = UnmarshalString("{\"age\":24,\"name\":\"张三\"}", &m)
	fmt.Println(m, err) // map[age:24 name:张三]
}
