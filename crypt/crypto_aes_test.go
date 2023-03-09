package crypt

import (
	"fmt"
	"testing"
)

func TestAesCBCEncrypt(t *testing.T) {
	var key = []byte("KEY_ABCDEFGHIJKL")
	var iv = []byte("IV_ABCDEFGHIJKLM")
	c, _ := AesEncrypt([]byte("https://minzhan.com/"), key, iv)
	d, _ := AesDecrypt("y5X0slKLtZbY6f3CkJXxali8JhyXk19alm5AfFFSWxA=", key, iv)
	fmt.Println(c)
	fmt.Println(d)
}
