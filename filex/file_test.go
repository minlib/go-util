package filex

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

func TestWriteFile(t *testing.T) {
	fileName := "C:\\Users\\Administrator\\Desktop\\temp\\temp.json"
	err := WriteFile(fileName, "this is test content.")
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
	}
}

func TestReadFile(t *testing.T) {
	content, err := ReadFile("C:\\Users\\Administrator\\Desktop\\temp\\temp.json")
	if err != nil {
		t.Errorf("ReadFile() error = %v", err)
	}
	fmt.Println(content)
}
