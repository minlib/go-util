package cmd

import (
	"fmt"
	"github.com/minlib/go-util/convert"
	"os/exec"
)

func WinCmd(cmd string) (string, error) {
	c := exec.Command("cmd", "/c", cmd)
	outputByte, err := c.CombinedOutput()
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	output := convert.BytesToString(outputByte, convert.GB18030)
	return output, nil
}
