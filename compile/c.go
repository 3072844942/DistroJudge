package compile

import (
	"DistroJudge/api"
	"DistroJudge/file"
	"DistroJudge/sandbox"
	"bytes"
	"context"
	"errors"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type c_Compile struct{}

func (c c_Compile) Compile(code string, language api.Task_Language, dir string) (execPath string, err error) {
	filePath := file.Path(dir + "/main.cpp")
	err = file.Write(code, filePath)
	if err != nil {
		return
	}

	// gcc编译文件并输出指定地点
	execPath = file.Path(dir + "/exe")
	cmd := sandbox.Command("g++", filePath, "-o", execPath)

	// 设置输出
	var out bytes.Buffer
	cmd.Stderr = &out
	if err = cmd.Run(); err != nil {
		err = errors.New(out.String())
		return
	}

	if runtime.GOOS == "linux" {
		// gcc需要可执行权限
		cmd = sandbox.Command("chmod", "+x", execPath)
		err = cmd.Run()
	}
	return
}

func (c c_Compile) Run(ctx context.Context, path string, language api.Task_Language, in string, t, memory uint64) (*Result, error) {
	// 创建一个带有超时的context
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(t)*time.Second)
	defer cancel()

	cmd := sandbox.CommandContext(ctxWithTimeout, path)

	// 创建一个channel用于传递结果和错误
	resultCh := make(chan *Result, 1)
	errCh := make(chan error, 1)

	// 启动子线程执行命令
	go c.childRun(cmd, in, resultCh, errCh)

	// 等待子线程执行完成或出错
	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return nil, err
	case <-ctxWithTimeout.Done():
		// 超时时终止命令
		if cmd.Process != nil {
			err := cmd.Process.Kill()
			if err != nil {
				return nil, err
			}
		}
		return nil, errors.New("判断超时")
	}
}

func (c c_Compile) childRun(cmd *exec.Cmd, in string, resultCh chan *Result, errCh chan error) {
	// 设置输入
	cmd.Stdin = strings.NewReader(in)
	// 设置输出
	var out bytes.Buffer
	cmd.Stdout = &out

	startTime := time.Now()
	err := cmd.Run()
	elapsedTime := time.Since(startTime).Milliseconds()

	result := &Result{
		OutPath: out.String(),
		Time:    uint64(elapsedTime),
		// todo 如何获取实时内存和高峰内存
		//Memory:  elapsedMemory,
	}

	if err != nil {
		errCh <- err
	} else {
		resultCh <- result
	}
}
