package crypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Bcrypt Bcrypt加密
func Bcrypt(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		panic(errors.New("加密失败," + err.Error()))
	}
	return string(hash)
}

// BcryptMatches Bcrypt验证
func BcryptMatches(s string, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(s))
	return err == nil
}
