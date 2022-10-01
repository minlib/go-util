package crypt

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
func Encrypt(s string, h hash.Hash) []byte {
	h.Write([]byte(s))
	return h.Sum(nil)
}

// Md5 MD5加密
func Md5(s string) []byte {
	return Encrypt(s, md5.New())
}

// Md5Hex Md5加密的16进制字符串
func Md5Hex(s string) string {
	return hex.EncodeToString(Md5(s))
}

// Sha1 SHA1加密
func Sha1(s string) []byte {
	return Encrypt(s, sha1.New())
}

// Sha1Hex Sha1加密的16进制字符串
func Sha1Hex(s string) string {
	return hex.EncodeToString(Sha1(s))
}

// Sha224 Sha224加密
func Sha224(s string) []byte {
	return Encrypt(s, sha256.New224())
}

// Sha224Hex Sha224加密的16进制字符串
func Sha224Hex(s string) string {
	return hex.EncodeToString(Sha224(s))
}

// Sha256 SHA256加密
func Sha256(s string) []byte {
	return Encrypt(s, sha256.New())
}

// Sha256Hex Sha256加密的16进制字符串
func Sha256Hex(s string) string {
	return hex.EncodeToString(Sha256(s))
}

// Sha384 SHA384加密
func Sha384(s string) []byte {
	return Encrypt(s, sha512.New384())
}

// Sha384Hex Sha384加密的16进制字符串
func Sha384Hex(s string) string {
	return hex.EncodeToString(Sha384(s))
}

// Sha512 Sha512加密
func Sha512(s string) []byte {
	return Encrypt(s, sha512.New())
}

// Sha512Hex Sha512加密的16进制字符串
func Sha512Hex(s string) string {
	return hex.EncodeToString(Sha512(s))
}

// HmacMd5 HmacMd5加密
func HmacMd5(s, secret string) []byte {
	return Encrypt(s, hmac.New(md5.New, []byte(secret)))
}

// HmacMd5Hex HmacMd5加密的16进制字符串
func HmacMd5Hex(s, secret string) string {
	return hex.EncodeToString(HmacMd5(s, secret))
}

// HmacSha1 HmacSha1加密
func HmacSha1(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha1.New, []byte(secret)))
}

// HmacSha1Hex HmacSha1加密的16进制字符串
func HmacSha1Hex(s, secret string) string {
	return hex.EncodeToString(HmacSha1(s, secret))
}

// HmacSha224 HmacSha224加密
func HmacSha224(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha256.New224, []byte(secret)))
}

// HmacSha224Hex HmacSha224加密的16进制字符串
func HmacSha224Hex(s, secret string) string {
	return hex.EncodeToString(HmacSha224(s, secret))
}

// HmacSha256 HmacSha256加密
func HmacSha256(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha256.New, []byte(secret)))
}

// HmacSha256Hex HmacSha256加密的16进制字符串
func HmacSha256Hex(s, secret string) string {
	return hex.EncodeToString(HmacSha256(s, secret))
}

// HmacSha384 HmacSha384加密
func HmacSha384(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha512.New384, []byte(secret)))
}

// HmacSha384Hex HmacSha384加密的16进制字符串
func HmacSha384Hex(s, secret string) string {
	return hex.EncodeToString(HmacSha384(s, secret))
}

// HmacSha512 HmacSha512加密
func HmacSha512(s, secret string) []byte {
	return Encrypt(s, hmac.New(sha512.New, []byte(secret)))
}

// HmacSha512Hex HmacSha512加密的16进制字符串
func HmacSha512Hex(s, secret string) string {
	return hex.EncodeToString(HmacSha512(s, secret))
}
