package crypt

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	fmt.Println("Base64Encode:", Base64Encode("minzhan.com"))
	fmt.Println("Base64Decode:", Base64Decode("bWluemhhbi5jb20="))
}
