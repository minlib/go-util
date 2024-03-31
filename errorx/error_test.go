package errorx

import (
	"fmt"
	"github.com/minlib/go-util/jsonx"
	"testing"
)

func TestNew(t *testing.T) {
	var err1 = New(500, "server busy")
	fmt.Println(err1)
	fmt.Println(jsonx.MarshalString(err1))
	fmt.Println(err1.Code)
	fmt.Println(err1.Message)
	fmt.Println(err1.Error())

	var err2 = New(500, "server busy")
	fmt.Println(err2)
	fmt.Println(jsonx.MarshalString(err2))
}

func TestNewParams(t *testing.T) {
	err := NewParams(501, "Hello %s", "Min Zhan")
	err2 := NewParams(501, "Hello")
	fmt.Println(err)
	fmt.Println(err2)
}

func TestFormat(t *testing.T) {
	fmt.Println(New(501, "Hello %s").Format("Zhang San"))
}
