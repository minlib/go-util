package httpx

import (
	netUrl "net/url"
)

func RawUrl(url string, params map[string][]string) (string, error) {
	u, err := netUrl.Parse(url)
	if err != nil {
		return "", err
	}
	values := u.Query()
	for key, v := range params {
		for _, value := range v {
			values.Add(key, value)
		}
	}
	u.RawQuery = values.Encode()
	return u.String(), nil
}
