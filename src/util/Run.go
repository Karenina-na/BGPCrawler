package util

import (
	"os/exec"
	"runtime"
)

// Run 执行命令
//
//	@Description: 执行命令
func Run(command string) (E error) {
	var cmd *exec.Cmd

	os := runtime.GOOS
	switch os {
	case "windows":
		cmd = exec.Command("cmd", "/C", command)
	case "linux":
		cmd = exec.Command("/bin/bash", "-c", command)
	}

	// 命令的输出直接扔掉
	_, err := cmd.Output()
	return err
}
