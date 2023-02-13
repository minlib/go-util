package errorx

import (
	"fmt"
	"testing"

	"github.com/minlib/go-util/json"
)

func TestNew(t *testing.T) {
	var err1 *Error = New(500, "server busy")
	fmt.Println(err1)
	fmt.Println(json.MarshalString(err1))
	fmt.Println(err1.Code)
	fmt.Println(err1.Message)
	fmt.Println(err1.Error())

	var err2 error = New(500, "server busy")
	fmt.Println(err2)
	fmt.Println(json.MarshalString(err2))
}
