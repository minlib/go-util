package stringutil

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

func TestZeroFill(t *testing.T) {
	fmt.Println(ZeroFill(1, 5))       // 00001
	fmt.Println(ZeroFill(100, 6))     // 000100
	fmt.Println(ZeroFill(123456, 6))  // 123456
	fmt.Println(ZeroFill(1234567, 6)) // 1234567
}

func TestHideLeftLimit(t *testing.T) {
	fmt.Println(len("abcde123"))        // 8
	fmt.Println(len("我是程序员123"))        // 18
	fmt.Println(RuneLength("abcde123")) // 8
	fmt.Println(RuneLength("我是程序员123")) // 8

	fmt.Println(ReplaceOffset("123456789abcd", '*', 4, 4)) // 1234****9abcd

	fmt.Println(HideLeftLimit("123456789abcd", 4, 4))  // 1234****9abcd
	fmt.Println(HideLeftLimit("123456789abcd", 4, 8))  // 1234********d
	fmt.Println(HideLeftLimit("123456789abcd", 4, 9))  // 1234*********
	fmt.Println(HideLeftLimit("123456789abcd", 4, 10)) // 1234*********
	fmt.Println(HideLeftLimit("我是一个热爱编程的程序员", 2, 7))   // 我是*******程序员
	fmt.Println(HideLeftLimit("123456789abcd", 20, 4)) // 123456789abcd

	fmt.Println(HideRightLimit("123456789abcd", 4, 4))  // 12345****abcd
	fmt.Println(HideRightLimit("123456789abcd", 4, 8))  // 1********abcd
	fmt.Println(HideRightLimit("123456789abcd", 4, 9))  // *********abcd
	fmt.Println(HideRightLimit("123456789abcd", 4, 10)) // *********abcd
	fmt.Println(HideRightLimit("我是一个热爱编程的程序员", 3, 5))   // 我是一个*******程序员

	fmt.Println(HideLeft("李", 1))        // 李
	fmt.Println(HideLeft("张三", 1))       // 张*
	fmt.Println(HideLeft("王五五", 1))      // 王**
	fmt.Println(HideLeft("我是程序员123", 5)) // 我是程序员***

	fmt.Println(HideRight("李", 1))   // 李
	fmt.Println(HideRight("张三", 1))  // *三
	fmt.Println(HideRight("王五五", 1)) // **五
}
