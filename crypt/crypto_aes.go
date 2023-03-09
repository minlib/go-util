package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AesEncrypt AES Encrypt CBC
func AesEncrypt(text []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// 填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	originData := append(text, padtext...)
	// 加密方式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AesDecrypt AES Decrypt CBC
func AesDecrypt(text string, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	crypted, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", nil
	}
	// 解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 输出到[]byte数组
	originData := make([]byte, len(crypted))
	blockMode.CryptBlocks(originData, crypted)
	// 去除填充,并返回
	length := len(originData)
	// 去掉最后一次的padding
	unpadding := int(originData[length-1])
	return string(originData[:(length - unpadding)]), nil
}
