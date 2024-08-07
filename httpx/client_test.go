package httpx

import (
	"fmt"
	"testing"

	"github.com/minlib/go-util/jsonx"
)

func TestGet(t *testing.T) {
	if resp, err := Get("https://minzhan.com"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestPost(t *testing.T) {
	data := map[string]string{
		"prompt": "这里是内容",
	}
	if resp, err := Post("http://localhost:8080/test", "application/json", jsonx.MarshalString(data)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func TestPostForm(t *testing.T) {
	data := map[string][]string{
		"prompt": {"this is key"},
		"value":  {"this is value"},
	}
	if resp, err := PostForm("http://localhost:8080/test2", data); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
