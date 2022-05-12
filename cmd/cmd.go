package cmd

import (
	"fmt"
	"os/exec"
)

func WinCmd(cmd string) (string, error) {
	command := exec.Command("cmd", "/c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("output err:", err)
		return "", err
	}
	fmt.Println(string(output))
	return string(output), nil
}
