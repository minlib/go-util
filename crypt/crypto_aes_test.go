package crypt

import (
	"fmt"
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	var key = "F78D51171B9186B7639B70D619090EEC"
	text := "https://minzhan.com/AESEncryptStd"
	encrypted, err := AESEncrypt(text, key)
	decrypted, err := AESDecrypt(encrypted, key)
	fmt.Println(encrypted, err)
	fmt.Println(decrypted, err)
	if text != decrypted {
		t.Errorf("AESEncrypt() got = %v, want %v", decrypted, text)
	}
}

func TestAESEncryptRawURL(t *testing.T) {
	var key = "F78D51171B9186B7639B70D619090EEC"
	text := "https://minzhan.com/AESEncryptRawURL"
	encrypted, err := AESEncryptRawURL(text, key)
	decrypted, err := AESDecryptRawURL(encrypted, key)
	fmt.Println(encrypted, err)
	fmt.Println(decrypted, err)
	if text != decrypted {
		t.Errorf("AESEncrypt() got = %v, want %v", decrypted, text)
	}
}

func TestAESCBCEncryptStd(t *testing.T) {
	var key = "F78D51171B9186B7639B70D619090EEC"
	var iv = "IV_ABCDEFGHIJKLM"
	text := "https://minzhan.com/AESEncryptRawURL"
	encrypted, err := AESCBCEncryptStd(text, key, iv)
	decrypted, err := AESCBCDecryptStd(encrypted, key, iv)
	fmt.Println(encrypted, err)
	fmt.Println(decrypted, err)
	if text != decrypted {
		t.Errorf("AESEncrypt() got = %v, want %v", decrypted, text)
	}
}
