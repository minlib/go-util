package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AESEncrypt(text string, key []byte) (string, error) {
	return AESCBCEncryptStd(text, key, key)
}

func AESDecrypt(text string, key []byte) (string, error) {
	return AESCBCDecryptStd(text, key, key)
}

func AESEncryptRawURL(text string, key []byte) (string, error) {
	return AESCBCEncryptRawURL(text, key, key)
}

func AESDecryptRawURL(text string, key []byte) (string, error) {
	return AESCBCDecryptRawURL(text, key, key)
}

func AESCBCEncryptRawURL(text string, key []byte, iv []byte) (string, error) {
	encrypted, err := AESCBCEncrypt([]byte(text), key, iv)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(encrypted), nil
}

func AESCBCDecryptRawURL(text string, key []byte, iv []byte) (string, error) {
	encrypted, err := base64.RawURLEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	decrypted, err := AESCBCDecrypt(encrypted, key, iv)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func AESCBCEncryptStd(text string, key []byte, iv []byte) (string, error) {
	encrypted, err := AESCBCEncrypt([]byte(text), key, iv)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func AESCBCDecryptStd(text string, key []byte, iv []byte) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	decrypted, err := AESCBCDecrypt(encrypted, key, iv)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

// AESCBCEncrypt encrypts data with AES algorithm in CBC mode
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256
// Note that AES block size is 16 bytes
func AESCBCEncrypt(text []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := PKCS7Padding(text, blockSize)
	// 加密方式
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize])
	encrypted := make([]byte, len(originData))
	blockMode.CryptBlocks(encrypted, originData)
	return encrypted, nil
}

// AESCBCDecrypt decrypts cipher text with AES algorithm in CBC mode
// Note that key length must be 16, 24 or 32 bytes to select AES-128, AES-192, or AES-256
// Note that AES block size is 16 bytes
func AESCBCDecrypt(encrypted []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	originData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(originData, encrypted)
	return PKCS7UnPadding(originData), nil
}

// PKCS7Padding fills plaintext as an integral multiple of the block length
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS7UnPadding removes padding data from the tail of plaintext
func PKCS7UnPadding(originData []byte) []byte {
	length := len(originData)
	unPadding := int(originData[length-1])
	return originData[:(length - unPadding)]
}
