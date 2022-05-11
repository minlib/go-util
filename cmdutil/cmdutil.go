package cmdutil

import (
	"fmt"
	"os/exec"
)

func WindowsCmd(cmd string) (string, error) {
	command := exec.Command("cmd", "/c", cmd)
	//command := exec.Command("tasklist")
	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("output err:", err)
		return "", err
	}
	fmt.Println(string(output))
	return string(output), nil
}
