package crypt

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	fmt.Println("Base64Encode:", Base64EncodeString("minzhan.com"))
	fmt.Println("Base64Decode:", Base64DecodeString("bWluemhhbi5jb20="))
}
