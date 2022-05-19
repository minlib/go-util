package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// Encrypt 加密
func Encrypt(s string, h hash.Hash) string {
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5 MD5加密
func Md5(s string) string {
	return Encrypt(s, md5.New())
}

// Sha1 SHA1加密
func Sha1(s string) string {
	return Encrypt(s, sha1.New())
}

// Sha224 Sha224加密
func Sha224(s string) string {
	return Encrypt(s, sha256.New224())
}

// Sha256 SHA256加密
func Sha256(s string) string {
	return Encrypt(s, sha256.New())
}

// Sha384 SHA384加密
func Sha384(s string) string {
	return Encrypt(s, sha512.New384())
}

// Sha512 Sha512加密
func Sha512(s string) string {
	return Encrypt(s, sha512.New())
}

// HmacMd5 HmacMd5加密
func HmacMd5(s, secret string) string {
	return Encrypt(s, hmac.New(md5.New, []byte(secret)))
}

// HmacSha1 HmacSha1加密
func HmacSha1(s, secret string) string {
	return Encrypt(s, hmac.New(sha1.New, []byte(secret)))
}

// HmacSha224 HmacSha224加密
func HmacSha224(s, secret string) string {
	return Encrypt(s, hmac.New(sha256.New224, []byte(secret)))
}

// HmacSha256 HmacSha256加密
func HmacSha256(s, secret string) string {
	return Encrypt(s, hmac.New(sha256.New, []byte(secret)))
}

// HmacSha384 HmacSha384加密
func HmacSha384(s, secret string) string {
	return Encrypt(s, hmac.New(sha512.New384, []byte(secret)))
}

// HmacSha512 HmacSha512加密
func HmacSha512(s, secret string) string {
	return Encrypt(s, hmac.New(sha512.New, []byte(secret)))
}
