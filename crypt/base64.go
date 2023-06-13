package crypt

import (
	"encoding/base64"
)

// Base64EncodeString Base64编码
func Base64EncodeString(s string) string {
	return Base64Encode([]byte(s))
}

// Base64Encode Base64编码
func Base64Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

// Base64DecodeString Base64解码
func Base64DecodeString(s string) string {
	if bytes := Base64Decode(s); bytes != nil {
		return string(bytes)
	} else {
		return ""
	}
}

// Base64Decode Base64解码为字节
func Base64Decode(s string) []byte {
	if bytes, err := base64.StdEncoding.DecodeString(s); err == nil {
		return bytes
	}
	return nil
}
