package ziputil

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	err := Zip("C:\\Users\\Admin\\Desktop\\bean.tar.gz", "C:\\Users\\Admin\\Desktop\\bean", "C:\\Users\\Admin\\Desktop\\json.txt")
	fmt.Println("err is", err)
}

func TestUnzip(t *testing.T) {
	err := Unzip("C:\\Users\\Admin\\Desktop\\bean.tar.gz", "C:\\Users\\Admin\\Desktop\\bean3")
	fmt.Println("err is", err)
}
