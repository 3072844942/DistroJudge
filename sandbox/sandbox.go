package sandbox

import (
	"os/exec"
	"syscall"
)

func Command(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// todo
		ProcessAttributes: &syscall.SecurityAttributes{},
		// 隔离 uts,ipc,pid,mount,user,network
		ThreadAttributes: &syscall.SecurityAttributes{},
	}
	return cmd
}
