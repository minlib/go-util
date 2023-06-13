package crypt

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	s := "123456"
	secret := "minzhan.com"
	fmt.Println("Md5:", Md5String(s))
	fmt.Println("Sha1:", Sha1String(s))
	fmt.Println("Sha224:", Sha224String(s))
	fmt.Println("Sha256:", Sha256String(s))
	fmt.Println("Sha384:", Sha384String(s))
	fmt.Println("Sha512:", Sha512String(s))
	fmt.Println("HmacMd5:", HmacMd5String(s, secret))
	fmt.Println("HmacSha1:", HmacSha1String(s, secret))
	fmt.Println("HmacSha224:", HmacSha224String(s, secret))
	fmt.Println("HmacSha256:", HmacSha256String(s, secret))
	fmt.Println("HmacSha384:", HmacSha384String(s, secret))
	fmt.Println("HmacSha512:", HmacSha512String(s, secret))
}
