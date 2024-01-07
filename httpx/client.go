package httpx

import (
	"io"
	"net/http"
	"strings"

	"github.com/minlib/go-util/jsonx"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Post(url, contentType, data string) ([]byte, error) {
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func PostJson(url string, data any) ([]byte, error) {
	return Post(url, "application/json", jsonx.MarshalString(data))
}

func PostForm(url string, data map[string][]string) ([]byte, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
