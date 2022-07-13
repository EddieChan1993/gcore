package sh

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

/**
  可执行非x权限的shell
*/
func RunSubShell(cmdLine []byte) (cmd *exec.Cmd) {
	defaultShell := os.Getenv("SHELL")
	if defaultShell == "" {
		defaultShell = "/bin/bash"
	}
	cmd = exec.Command(defaultShell)
	// 环境变量注入
	cmd.Env = append(
		cmd.Env,
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
		fmt.Sprintf("PATH=/home/dhcd/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:%s", os.Getenv("PATH")),
	)
	cmd.Stdin = bytes.NewReader(cmdLine)
	return
}
