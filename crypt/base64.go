package crypt

import (
	"encoding/base64"
)

// Base64Encode Base64编码
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode Base64解码
func Base64Decode(s string) string {
	if bytes, err := base64.StdEncoding.DecodeString(s); err == nil {
		return string(bytes)
	}
	return ""
}
