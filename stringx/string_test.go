package stringx

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"testing"
)

func TestHasAnyPrefix(t *testing.T) {
	fmt.Println(HasAnyPrefix("abc.png"))
	fmt.Println(HasAnyPrefix("abc.png", "bb", "cc", "dd"))
	fmt.Println(HasAnyPrefix("abc.png", "bb", "cc", "dd", "abc"))
	fmt.Println(HasAnyPrefix("abc.png", "ab"))
}

func TestHasAnySuffix(t *testing.T) {
	fmt.Println(HasAnySuffix("1234.png"))
	fmt.Println(HasAnySuffix("1234.png", "jpg", "jpeg", "gif"))
	fmt.Println(HasAnySuffix("1234.png", "jpg", "jpeg", "gif", "png"))
	fmt.Println(HasAnySuffix("1234.png", "png"))
}

func TestContainsAnyString(t *testing.T) {
	fmt.Println(ContainsAnyString("1234.png1"))                              // false
	fmt.Println(ContainsAnyString("1234.png2", "jpg", "jpeg", "gif"))        // false
	fmt.Println(ContainsAnyString("1234.png3", "jpg", "jpeg", "gif", "png")) // true
	fmt.Println(ContainsAnyString("1234.png4", "PNG"))                       // false
}

func TestEqualAnyString(t *testing.T) {
	fmt.Println(EqualAnyString("1234.png1"))                                    // false
	fmt.Println(EqualAnyString("1234.png2", "1234.png2", "jpeg", "gif"))        // true
	fmt.Println(EqualAnyString("1234.png3", "1234.PNG3", "jpeg", "gif", "png")) // false
	fmt.Println(EqualAnyString("1234.png4", "1234.png4"))                       // true
}

func TestEqualAnyFold(t *testing.T) {
	fmt.Println(EqualAnyFold("png"))
	fmt.Println(EqualAnyFold("png", "jpg", "jpeg", "gif"))
	fmt.Println(EqualAnyFold("png", "jpg", "jpeg", "gif", "png"))
	fmt.Println(EqualAnyFold("png", "PNG"))
}

func TestUnderlineToUpperHump(t *testing.T) {
	s1 := "sys_user_role"
	fmt.Println(UnderlineToUpperHump(s1)) // SysUserToken

	s2 := "sys_user_role_2022"
	fmt.Println(UnderlineToUpperHump(s2)) // SysUserRole_2022

	s3 := "_sys_user_role"
	fmt.Println(UnderlineToUpperHump(s3)) // SysUserRole
}

func TestUnderlineToLowerHump(t *testing.T) {
	s1 := "sys_user_role"
	fmt.Println(UnderlineToLowerHump(s1)) // sysUserRole

	s2 := "sys_user_role_2022"
	fmt.Println(UnderlineToLowerHump(s2)) // sysUserRole_2022

	s3 := "_sys_user_role"
	fmt.Println(UnderlineToLowerHump(s3)) // sysUserRole
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

func TestIsBlank(t *testing.T) {
	fmt.Println(IsBlank(""))
	fmt.Println(IsBlank(" "))
	fmt.Println(IsBlank("  "))
	fmt.Println(IsBlank("	"))
	fmt.Println(IsBlank("minzhan.com") == false)
}

func TestIsNotBlank(t *testing.T) {
	fmt.Println(IsNotBlank("") == false)
	fmt.Println(IsNotBlank(" ") == false)
	fmt.Println(IsNotBlank("  ") == false)
	fmt.Println(IsNotBlank("		") == false)
	fmt.Println(IsNotBlank("minzhan.com"))
}

func TestIsAnyEmpty(t *testing.T) {
	fmt.Println(IsAnyEmpty() == false)
	fmt.Println(IsAnyEmpty("") == true)
	fmt.Println(IsAnyEmpty("", " ") == true)
	fmt.Println(IsAnyEmpty("", "minzhan") == true)
	fmt.Println(IsAnyEmpty("min", "minzhan") == false)
}

func TestIsAnyBlank(t *testing.T) {
	fmt.Println(IsAnyBlank() == false)
	fmt.Println(IsAnyBlank("") == true)
	fmt.Println(IsAnyBlank("", " 		") == true)
	fmt.Println(IsAnyBlank("	", " 		") == true)
	fmt.Println(IsAnyBlank("", "minzhan") == true)
	fmt.Println(IsAnyBlank("min", "minzhan") == false)
}

func TestIsAnyNotEmpty(t *testing.T) {
	fmt.Println(IsAnyNotEmpty() == false)
	fmt.Println(IsAnyNotEmpty("") == false)
	fmt.Println(IsAnyNotEmpty("", " ") == true)
	fmt.Println(IsAnyNotEmpty("", "			") == true)
	fmt.Println(IsAnyNotEmpty("", "minzhan") == true)
	fmt.Println(IsAnyNotEmpty("min", "minzhan") == true)
}

func TestIsAnyNotBlank(t *testing.T) {
	fmt.Println(IsAnyNotBlank() == false)
	fmt.Println(IsAnyNotBlank("") == false)
	fmt.Println(IsAnyNotBlank("", " 		") == false)
	fmt.Println(IsAnyNotBlank("	", " 		") == false)
	fmt.Println(IsAnyNotBlank("", "minzhan") == true)
	fmt.Println(IsAnyNotBlank("min", "minzhan") == true)
}

func TestIsNoneEmpty(t *testing.T) {
	fmt.Println(IsNoneEmpty() == false)
	fmt.Println(IsNoneEmpty("") == false)
	fmt.Println(IsNoneEmpty("", " 		") == false)
	fmt.Println(IsNoneEmpty("	", " 		") == true)
	fmt.Println(IsNoneEmpty("", "minzhan") == false)
	fmt.Println(IsNoneEmpty("min", "minzhan") == true)
}

func TestIsNoneBlank(t *testing.T) {
	fmt.Println(IsNoneBlank() == false)
	fmt.Println(IsNoneBlank("") == false)
	fmt.Println(IsNoneBlank("", " 		") == false)
	fmt.Println(IsNoneBlank("	", " 		") == false)
	fmt.Println(IsNoneBlank("", "minzhan") == false)
	fmt.Println(IsNoneBlank("min", "minzhan") == true)
}

func TestSplitToIntegers(t *testing.T) {
	printf(SplitToIntegers[int16]("1111,222,333", ","))
	printf(SplitToIntegers[int32]("1111,222,333", ","))
	printf(SplitToIntegers[int64]("1111,222,333", ","))
}

func TestJoin(t *testing.T) {
	type args[E constraints.Integer] struct {
		elems []E
		sep   string
	}
	type testCase[E constraints.Integer] struct {
		name string
		args args[E]
		want string
	}
	intTests := []testCase[int]{{
		name: "test1",
		args: args[int]{
			elems: []int{1, 2, 3, 4, 5},
			sep:   ",",
		},
		want: "1,2,3,4,5",
	}, {
		name: "test2",
		args: args[int]{
			elems: []int{1, 2, 3, 4, 5},
			sep:   " & ",
		},
		want: "1 & 2 & 3 & 4 & 5",
	}}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.elems, tt.args.sep); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
	int64Tests := []testCase[int64]{{
		name: "test1",
		args: args[int64]{
			elems: []int64{1, 2, 3, 4, 5},
			sep:   ",",
		},
		want: "1,2,3,4,5",
	}, {
		name: "test2",
		args: args[int64]{
			elems: []int64{1, 2, 3, 4, 5},
			sep:   " & ",
		},
		want: "1 & 2 & 3 & 4 & 5",
	}}
	for _, tt := range int64Tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.elems, tt.args.sep); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}

	uintTests := []testCase[uint]{{
		name: "test1",
		args: args[uint]{
			elems: []uint{1, 2, 3, 4, 5},
			sep:   ",",
		},
		want: "1,2,3,4,5",
	}, {
		name: "test2",
		args: args[uint]{
			elems: []uint{1, 2, 3, 4, 5},
			sep:   " & ",
		},
		want: "1 & 2 & 3 & 4 & 5",
	}}
	for _, tt := range uintTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.elems, tt.args.sep); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func printf(s any, err error) {
	fmt.Printf("%T, %v, err: %v\n", s, s, err)
}
