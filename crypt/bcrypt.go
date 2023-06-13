package crypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Bcrypt Bcrypt加密
func Bcrypt(bytes []byte) string {
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	if err != nil {
		panic(errors.New("加密失败," + err.Error()))
	}
	return string(hash)
}

// BcryptString Bcrypt加密字符串
func BcryptString(s string) string {
	return Bcrypt([]byte(s))
}

// BcryptMatches Bcrypt验证
func BcryptMatches(bytes, encodedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(encodedPassword, bytes)
	return err == nil
}

// BcryptMatchesString Bcrypt验证字符串
func BcryptMatchesString(s string, encodedPassword string) bool {
	return BcryptMatches([]byte(s), []byte(encodedPassword))
}
