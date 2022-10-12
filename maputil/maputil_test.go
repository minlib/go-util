package maputil

import (
	"fmt"
	"testing"
)

func TestContainsKey(t *testing.T) {
	map1 := make(map[string]interface{})
	map1["a"] = 1
	fmt.Println(ContainsKey(map1, "a")) // true
	fmt.Println(ContainsKey(map1, "b")) // false

	map2 := make(map[int]bool)
	map2[1] = true
	fmt.Println(ContainsKey(map2, 1)) // true
	fmt.Println(ContainsKey(map2, 2)) // false
}
