package httpx

import (
	"fmt"
	"net/url"
	"testing"
)

func TestRawUrl(t *testing.T) {
	urlParams := url.Values{}
	urlParams.Add("b", "456")
	urlParams.Add("c", "abc")
	rawUrl, err := RawUrl("https://api.weixin.qq.com/wxa/revertcoderelease?a=123", urlParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rawUrl)
	}
}
