package main

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	s := "123456"
	secret := "minzhan.com"
	fmt.Println("Md5:", Md5Hex(s))
	fmt.Println("Sha1:", Sha1Hex(s))
	fmt.Println("Sha224:", Sha224Hex(s))
	fmt.Println("Sha256:", Sha256Hex(s))
	fmt.Println("Sha384:", Sha384Hex(s))
	fmt.Println("Sha512:", Sha512Hex(s))
	fmt.Println("HmacMd5:", HmacMd5Hex(s, secret))
	fmt.Println("HmacSha1:", HmacSha1Hex(s, secret))
	fmt.Println("HmacSha224:", HmacSha224Hex(s, secret))
	fmt.Println("HmacSha256:", HmacSha256Hex(s, secret))
	fmt.Println("HmacSha384:", HmacSha384Hex(s, secret))
	fmt.Println("HmacSha512:", HmacSha512Hex(s, secret))
}
