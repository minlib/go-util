package strs

import (
	"fmt"
	"testing"
)

func TestUnderlineToUpperHump(t *testing.T) {
	s1 := "sys_user_role"
	fmt.Println(UnderlineToUpperHump(s1)) // SysUserToken

	s2 := "sys_user_role_2022"
	fmt.Println(UnderlineToUpperHump(s2)) // SysUserRole_2022

	s3 := "_sys_user_role"
	fmt.Println(UnderlineToUpperHump(s3)) // SysUserRole
}

func TestHumpToUnderline(t *testing.T) {
	s1 := HumpToUnderline("SysUserToken2022") // sys_user_token2022
	fmt.Println(s1)

	s2 := HumpToUnderline("SysUserToken") // sys_user_token
	fmt.Println(s2)

	s3 := HumpToUnderline("_SysUserToken") // sys_user_token
	fmt.Println(s3)

	s4 := HumpToUnderline("sysUserToken") // sys_user_token
	fmt.Println(s4)
}
