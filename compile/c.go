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
	"sync"
	"time"
)

type c_Compile struct{}

func (c c_Compile) Compile(code string, language api.Task_Language, dir string) (execPath string, err error) {
	filePath := dir + "/main.cpp"
	err = file.Write(code, filePath)
	if err != nil {
		return
	}

	// gcc编译文件并输出指定地点
	execPath = dir + "/exe"
	cmd := sandbox.Command("g++", filePath, "-o", execPath)

	// 设置输出
	var out bytes.Buffer
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return
	}
	if out.String() != "" {
		err = errors.New(out.String())
		return
	}

	// gcc需要可执行权限
	cmd = sandbox.Command("chmod", "+x", execPath)
	err = cmd.Run()
	return
}

func (c c_Compile) Run(ctx context.Context, path string, language api.Task_Language, in string, t, memory uint64) (*Result, error) {
	cmd := sandbox.CommandContext(ctx, path)

	// 创建一个channel用于传递结果和错误
	resultCh := make(chan *Result, 1)
	errCh := make(chan error, 1)

	// 启动子线程执行命令
	go c.childRun(cmd, in, resultCh, errCh)

	// 监控线程，用于取消命令和监控时间
	go func() {
		select {
		case <-ctx.Done():
			// 如果父context取消，则取消子线程执行的命令
			_ = cmd.Process.Kill()
		case <-time.After(time.Duration(t) * time.Millisecond):
			// 如果时间超过限制，则取消子线程执行的命令
			_ = cmd.Process.Kill()
		}
	}()

	// 等待子线程执行完成或出错
	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return nil, err
	}
}

func (c c_Compile) childRun(cmd *exec.Cmd, in string, resultCh chan *Result, errCh chan error) {
	// 设置输入
	cmd.Stdin = strings.NewReader(in)
	// 设置输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 获取程序执行前的内存状态
	startMemory := getGoroutineMemConsume()
	startTime := time.Now()

	err := cmd.Run()

	elapsedTime := time.Since(startTime).Milliseconds()
	elapsedMemory := getGoroutineMemConsume() - startMemory

	result := &Result{
		OutPath: out.String(),
		Time:    uint64(elapsedTime),
		Memory:  elapsedMemory,
	}

	// 将结果或错误发送到相应的channel
	if err != nil {
		errCh <- err
	} else {
		resultCh <- result
	}
	close(resultCh)
	close(errCh)
}

// getGoroutineMemConsume return memory use with b
func getGoroutineMemConsume() uint64 {
	var c chan int
	var wg sync.WaitGroup
	const goroutineNum = 1e4 // 1 * 10^4

	memConsumed := func() uint64 {
		runtime.GC() //GC，排除对象影响
		var memStat runtime.MemStats
		runtime.ReadMemStats(&memStat)
		return memStat.Sys
	}

	noop := func() {
		wg.Done()
		<-c //防止goroutine退出，内存被释放
	}

	wg.Add(goroutineNum)
	before := memConsumed() //获取创建goroutine前内存
	for i := 0; i < goroutineNum; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed() //获取创建goroutine后内存

	return uint64(float64(after-before) / goroutineNum)
}
