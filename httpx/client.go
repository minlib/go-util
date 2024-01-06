package httpx

import (
	"github.com/minlib/go-util/jsonx"
	"io"
	"net/http"
	netUrl "net/url"
	"strings"
)

func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Post(url, contentType, data string) (string, error) {
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PostJson(url string, data any) (string, error) {
	return Post(url, "application/json", jsonx.MarshalString(data))
}

func PostForm(url string, data map[string][]string) (string, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func RawUrl(url string, urlParams netUrl.Values) (string, error) {
	u, err := netUrl.Parse(url)
	if err != nil {
		return "", err
	}
	values := u.Query()
	for key, v := range urlParams {
		for _, value := range v {
			values.Add(key, value)
		}
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}
