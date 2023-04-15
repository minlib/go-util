package jsonx

import (
	"fmt"
	"testing"
)

// import (
// 	"fmt"
// 	"testing"
// )

// func TestMarshal(t *testing.T) {
// 	fmt.Println(Marshal([]int64{12, 34, 567}))       // [12,34,567]
// 	fmt.Println(Marshal([]string{"A", "BC", "DEF"})) // ["A","BC","DEF"]

// 	m := make(map[string]any, 2)
// 	m["age"] = 24
// 	m["name"] = "张三"
// 	fmt.Println(Marshal(m)) // {"age":24,"name":"张三"}
// }

// func TestUnmarshal(t *testing.T) {
// 	i := []int64{}
// 	Unmarshal("[12,34,567]", &i)
// 	fmt.Println(i) // [12 34 567]

// 	s := []string{}
// 	Unmarshal("[\"A\",\"BC\",\"DEF\"]", &s)
// 	fmt.Println(s) // [A BC DEF]

// 	m := make(map[string]any, 0)
// 	Unmarshal("{\"age\":24,\"name\":\"张三\"}", &m)
// 	fmt.Println(m) // map[age:24 name:张三]
// }

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
	UnmarshalString("[12,34,567]", &i)
	fmt.Println(i) // [12 34 567]

	s := []string{}
	UnmarshalString("[\"A\",\"BC\",\"DEF\"]", &s)
	fmt.Println(s) // [A BC DEF]

	m := make(map[string]any, 0)
	UnmarshalString("{\"age\":24,\"name\":\"张三\"}", &m)
	fmt.Println(m) // map[age:24 name:张三]
}
