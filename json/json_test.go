package json

import (
	"fmt"
	"testing"
)

func TestToJsonString(t *testing.T) {
	fmt.Println(ToJsonString([]int64{12, 34, 567}))       // [12,34,567]
	fmt.Println(ToJsonString([]string{"A", "BC", "DEF"})) // ["A","BC","DEF"]

	m := make(map[string]any, 2)
	m["age"] = 24
	m["name"] = "张三"
	fmt.Println(ToJsonString(m)) // {"age":24,"name":"张三"}
}

func TestParse(t *testing.T) {
	i := []int64{}
	Parse("[12,34,567]", &i)
	fmt.Println(i) // [12 34 567]

	s := []string{}
	Parse("[\"A\",\"BC\",\"DEF\"]", &s)
	fmt.Println(s) // [A BC DEF]

	m := make(map[string]any, 0)
	Parse("{\"age\":24,\"name\":\"张三\"}", &m)
	fmt.Println(m) // map[age:24 name:张三]
}
