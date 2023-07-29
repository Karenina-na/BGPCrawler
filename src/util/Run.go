package util

import "os/exec"

// Run 执行命令
//
//	@Description: 执行命令
func Run(command string) {
	// 识别是windows还是linux
	println(command)
	cmd := exec.Command("cmd", "/C", command)
	//cmd := exec.Command("/bin/bash", "-c", command)
	// 命令的输出直接扔掉
	_, err := cmd.Output()
	// 命令出错
	if err != nil {
		panic(err.Error())
	}
	// 命令启动和启动时出错
	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}
	// 等待结束
	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}
}
