package crypt

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	s := "minzhan.com"
	s1 := Base64Encode(s)
	s2 := Base64Decode(s1)
	fmt.Println("Base64Encode:", s1)
	fmt.Println("Base64Decode:", s2)
}
