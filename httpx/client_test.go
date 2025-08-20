package httpx

import (
	"fmt"
	"github.com/minlib/go-util/convert"
	"github.com/minlib/go-util/jsonx"
	"testing"
	"time"
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

func TestPostFormNew(t *testing.T) {
	client := NewHttpClient(3 * time.Second)
	requestUrl := "https://ym28.sydiaoke.cn/api/company/getCompany"
	data := map[string]string{
		"keyword": "腾讯",
		"page":    convert.IntToString(1),
	}
	headers := map[string]string{
		"token":        "2d92f38d-e4b6-46a3-980b-26f234f68d49",
		"Content-Type": "application/json; charset=utf-8",
	}
	bytes, code, err := client.PostForm(requestUrl, headers, data)
	fmt.Println(string(bytes), code, err)
}
