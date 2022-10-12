package fileutil

import (
	"fmt"
	"testing"
)

var path1 = "d:\\temp\\a\\"
var path2 = "d:\\temp\\b\\f.txt"
var path3 = "d:\\temp\\c\\d\\f.txt"
var path4 = "d:\\temp\\e"

func TestDir(t *testing.T) {
	fmt.Println(Dir(path1)) // d:\temp\a
	fmt.Println(Dir(path2)) // d:\temp\b
	fmt.Println(Dir(path3)) // d:\temp\c\d
	fmt.Println(Dir(path4)) // d:\temp
}

func TestIsDir(t *testing.T) {
	fmt.Println(IsDir(path1)) // true
	fmt.Println(IsDir(path2)) // false, The system cannot find the file specified.
	fmt.Println(IsDir(path3)) // false, The system cannot find the file specified.
	fmt.Println(IsDir(path4)) // false, The system cannot find the file specified.
}

func TestExist(t *testing.T) {
	fmt.Println(Exist(path1)) //
	fmt.Println(Exist(path2)) //
	fmt.Println(Exist(path3)) //
	fmt.Println(Exist(path4)) // Windows create file e command: echo > e
}

func TestMkdirAll(t *testing.T) {
	err := MkdirAll(path1, path2, path3, path4)
	fmt.Println(err)
}
