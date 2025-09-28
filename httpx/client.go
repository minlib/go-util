package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/minlib/go-util/jsonx"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HttpClient 封装HTTP客户端，可配置超时等参数
type HttpClient struct {
	client *http.Client
}

// NewHttpClient 创建新的HTTP客户端实例，可指定超时时间
func NewHttpClient(timeout time.Duration) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Request 发送请求
func (c *HttpClient) Request(method, requestUrl string, headers map[string]string, body io.Reader) ([]byte, int, error) {
	req, err := http.NewRequest(method, requestUrl, body)
	if err != nil {
		return nil, 0, fmt.Errorf("创建请求失败: %w", err)
	}
	// 添加自定义Header
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	// 发送请求
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("读取响应体失败: %w", err)
	}
	return bytes, resp.StatusCode, nil
}

// Get 发送GET请求
func (c *HttpClient) Get(requestUrl string, headers map[string]string) ([]byte, int, error) {
	return c.Request(http.MethodGet, requestUrl, headers, nil)
}

// Post 发送POST请求
func (c *HttpClient) Post(requestUrl string, headers map[string]string, data interface{}) ([]byte, int, error) {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, 0, fmt.Errorf("JSON序列化失败: %w", err)
	}
	body := bytes.NewBuffer(jsonBody)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json; charset=utf-8"
	return c.Request(http.MethodPost, requestUrl, headers, body)
}

// PostForm 发送POST表单请求
func (c *HttpClient) PostForm(requestUrl string, headers map[string]string, data map[string]string) ([]byte, int, error) {
	values := url.Values{}
	for key, value := range data {
		values.Set(key, value)
	}
	body := strings.NewReader(values.Encode())
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return c.Request(http.MethodPost, requestUrl, headers, body)
}

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

func PostJson(url string, data interface{}) ([]byte, error) {
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
