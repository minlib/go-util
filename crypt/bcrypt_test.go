package crypt

import (
	"fmt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	s := "123456"
	fmt.Println("Bcrypt:", BcryptString(s))
	fmt.Println("BcryptMatches:", BcryptMatchesString(s, "$2a$04$Mm1ezETBYxskau2pbSPTI.Dqhpj9SGd5yzGSDKilRV8.WQVKtRpMC"))
}
