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
func Encrypt(bytes []byte, h hash.Hash) []byte {
	h.Write(bytes)
	return h.Sum(nil)
}

// EncryptString 加密
func EncryptString(s string, h hash.Hash) []byte {
	return Encrypt([]byte(s), h)
}

// Md5 MD5加密
func Md5(bytes []byte) string {
	return hex.EncodeToString(Encrypt(bytes, md5.New()))
}

// Md5String MD5加密
func Md5String(s string) string {
	return Md5([]byte(s))
}

// Sha1 SHA1加密
func Sha1(bytes []byte) string {
	return hex.EncodeToString(Encrypt(bytes, sha1.New()))
}

// Sha1String SHA1加密
func Sha1String(s string) string {
	return Sha1([]byte(s))
}

// Sha224 Sha224加密
func Sha224(bytes []byte) string {
	return hex.EncodeToString(Encrypt(bytes, sha256.New224()))
}

// Sha224String Sha224加密
func Sha224String(s string) string {
	return Sha224([]byte(s))
}

// Sha256 SHA256加密
func Sha256(bytes []byte) string {
	return hex.EncodeToString(Encrypt(bytes, sha256.New()))
}

// Sha256String SHA256加密
func Sha256String(s string) string {
	return Sha256([]byte(s))
}

// Sha384 SHA384加密
func Sha384(bytes []byte) string {
	return hex.EncodeToString(Encrypt(bytes, sha512.New384()))
}

// Sha384String SHA384加密
func Sha384String(s string) string {
	return Sha384([]byte(s))
}

// Sha512 Sha512加密
func Sha512(bytes []byte) string {
	return hex.EncodeToString(Encrypt(bytes, sha512.New()))
}

// Sha512String Sha512加密
func Sha512String(s string) string {
	return Sha512([]byte(s))
}

// HmacMd5 HmacMd5加密
func HmacMd5(bytes, secret []byte) string {
	return hex.EncodeToString(Encrypt(bytes, hmac.New(md5.New, secret)))
}

// HmacMd5String HmacMd5加密
func HmacMd5String(s, secret string) string {
	return HmacMd5([]byte(s), []byte(secret))
}

// HmacSha1 HmacSha1加密
func HmacSha1(bytes, secret []byte) string {
	return hex.EncodeToString(Encrypt(bytes, hmac.New(sha1.New, secret)))
}

// HmacSha1String HmacSha1加密
func HmacSha1String(s, secret string) string {
	return HmacSha1([]byte(s), []byte(secret))
}

// HmacSha224 HmacSha224加密
func HmacSha224(bytes, secret []byte) string {
	return hex.EncodeToString(Encrypt(bytes, hmac.New(sha256.New224, secret)))
}

// HmacSha224String HmacSha224加密
func HmacSha224String(s, secret string) string {
	return HmacSha224([]byte(s), []byte(secret))
}

// HmacSha256 HmacSha256加密
func HmacSha256(bytes, secret []byte) string {
	return hex.EncodeToString(Encrypt(bytes, hmac.New(sha256.New, secret)))
}

// HmacSha256String HmacSha256加密
func HmacSha256String(s, secret string) string {
	return HmacSha256([]byte(s), []byte(secret))
}

// HmacSha384 HmacSha384加密
func HmacSha384(bytes, secret []byte) string {
	return hex.EncodeToString(Encrypt(bytes, hmac.New(sha512.New384, secret)))
}

// HmacSha384String HmacSha384加密
func HmacSha384String(s, secret string) string {
	return HmacSha384([]byte(s), []byte(secret))
}

// HmacSha512 HmacSha512加密
func HmacSha512(bytes, secret []byte) string {
	return hex.EncodeToString(Encrypt(bytes, hmac.New(sha512.New, secret)))
}

// HmacSha512String HmacSha512加密
func HmacSha512String(s, secret string) string {
	return HmacSha512([]byte(s), []byte(secret))
}
