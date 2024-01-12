package compile

import (
	"DistroJudge/sandbox"
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestC_Compile_Run(t *testing.T) {
	c := c_Compile{}

	run, err := c.Run(context.Background(), "D:\\data\\108\\exe", 0, "1 2", 1024, 0)
	fmt.Println(run)
	if err != nil {
		t.Errorf("err: %v", err)
	}
}

func TestC_Compile_childRun(t *testing.T) {
	c := c_Compile{}

	cmd := exec.Command("D:\\data\\108\\exe")

	// 创建一个channel用于传递结果和错误
	resultCh := make(chan *Result, 1)
	errCh := make(chan error, 1)
	defer close(resultCh)
	defer close(errCh)

	go c.childRun(cmd, "1 2", resultCh, errCh)

	select {
	case result := <-resultCh:
		fmt.Println(result)
	case err := <-errCh:
		t.Errorf("err: %v", err)
	}
}

func TestC_Compile_childRun_WithTimeout(t *testing.T) {
	c := c_Compile{}
	tt := int64(1024)

	// 创建一个channel用于传递结果和错误
	resultCh := make(chan *Result, 1)
	errCh := make(chan error, 1)
	defer close(resultCh)
	defer close(errCh)

	ctx := context.Background()
	// 创建一个带有超时的context
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(tt)*time.Millisecond)
	defer cancel()
	cmd := sandbox.CommandContext(ctxWithTimeout, "D:\\data\\108\\exe")

	go c.childRun(cmd, "1 2", resultCh, errCh)

	select {
	case result := <-resultCh:
		fmt.Println(result)
	case err := <-errCh:
		t.Errorf("err: %v", err)
	case <-ctxWithTimeout.Done():
		// 超时时终止命令
		if cmd.Process != nil {
			err := cmd.Process.Kill()
			if err != nil {
				t.Errorf("err: %v", err)
			}
		}
		t.Errorf("判断超时")
	}
}
