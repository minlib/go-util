package main

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	s := "123456"
	secret := "MZ0011234"
	fmt.Println("Md5:", Md5(s))
	fmt.Println("Sha1:", Sha1(s))
	fmt.Println("Sha224:", Sha224(s))
	fmt.Println("Sha256:", Sha256(s))
	fmt.Println("Sha384:", Sha384(s))
	fmt.Println("Sha512:", Sha512(s))
	fmt.Println("HmacMd5:", HmacMd5(s, secret))
	fmt.Println("HmacSha1:", HmacSha1(s, secret))
	fmt.Println("HmacSha224:", HmacSha224(s, secret))
	fmt.Println("HmacSha256:", HmacSha256(s, secret))
	fmt.Println("HmacSha384:", HmacSha384(s, secret))
	fmt.Println("HmacSha512:", HmacSha512(s, secret))
}
