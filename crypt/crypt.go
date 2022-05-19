package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

// Encrypt 加密
func Encrypt(s string, h hash.Hash) []byte {
	h.Write([]byte(s))
	return h.Sum(nil)
}

// Md5 MD5加密
func Md5(s string) []byte {
	return Encrypt(s, md5.New())
}

// Sha1 SHA1加密
func Sha1(s string) []byte {
	return Encrypt(s, sha1.New())
}

// Sha224 Sha224加密
func Sha224(s string) []byte {
	return Encrypt(s, sha256.New224())
}

// Sha256 SHA256加密
func Sha256(s string) []byte {
	return Encrypt(s, sha256.New())
}

// Sha384 SHA384加密
func Sha384(s string) []byte {
	return Encrypt(s, sha512.New384())
}

// Sha512 Sha512加密
func Sha512(s string) []byte {
	return Encrypt(s, sha512.New())
}

// HmacMd5 HmacMd5加密
func HmacMd5(s, secret string) []byte {
	return Encrypt(s, hmac.New(md5.New, []byte(secret)))
}

// HmacSha1 HmacSha1加密
func HmacSha1(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha1.New, []byte(secret)))
}

// HmacSha224 HmacSha224加密
func HmacSha224(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha256.New224, []byte(secret)))
}

// HmacSha224Hex 获取HmacSha224加密的字符串
func HmacSha224Hex(s, secret string) string {
	return string(HmacSha224(s, secret))
}

// HmacSha256 HmacSha256加密
func HmacSha256(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha256.New, []byte(secret)))
}

// HmacSha256Hex 获取HmacSha256加密的字符串
func HmacSha256Hex(s, secret string) string {
	return string(HmacSha256(s, secret))
}

// HmacSha384 获取HmacSha384加密的字节流
func HmacSha384(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha512.New384, []byte(secret)))
}

// HmacSha384Hex 获取HmacSha384加密的字符串
func HmacSha384Hex(s, secret string) string {
	return string(HmacSha384(s, secret))
}

// HmacSha512 获取HmacSha512加密的字节流
func HmacSha512(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha512.New, []byte(secret)))
}

// HmacSha512Hex 获取HmacSha512加密的字符串
func HmacSha512Hex(s, secret string) string {
	return string(HmacSha512(s, secret))
}
