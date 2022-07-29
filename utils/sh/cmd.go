package sh

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type cmd struct {
	output chan string
}

type fnCall func(content, filename string) bool

func InitCmd() Cmder {
	output := make(chan string)
	return &cmd{
		output: output,
	}
}

func (this_ *cmd) OutPut() <-chan string {
	return this_.output
}

//IoStdout 同步执行shell
func (this_ *cmd) IoStdout(ctx context.Context, sh string) (io.ReadCloser, error) {
	c := exec.CommandContext(ctx, "bash", "-c", sh) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = c.Start()
	if err != nil {
		return nil, err
	}
	return stdout, nil
}

func (this_ *cmd) IoFile(filePath string) (io.Reader, error) {
	fn, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return fn, nil
}

//AsyncExecSh 异步执行shell
func (this_ *cmd) AsyncExecSh(ctx context.Context, ioR io.Reader) error {
	go func() {
		reader := bufio.NewReader(ioR)
		for {
			select {
			// 检测到ctx.Done()之后停止读取
			case <-ctx.Done():
				fmt.Println("处理超时，自动断开")
				return
			default:
				readString, err := reader.ReadString('\n')
				if err != nil || err == io.EOF {
					this_.output <- io.EOF.Error()
					return
				}
				this_.output <- readString
			}
		}
	}()
	return nil
}

func (this_ *cmd) AsyncExecFile(ctx context.Context, ioR io.Reader, fname string, fn fnCall) error {
	go func() {
		reader := bufio.NewReader(ioR)
		var rows int
		for {
			select {
			// 检测到ctx.Done()之后停止读取
			case <-ctx.Done():
				fmt.Println("处理超时，自动断开")
				return
			default:
				rows++
				readString, err := reader.ReadString('\n')
				if err != nil || err == io.EOF {
					this_.output <- io.EOF.Error()
					return
				}
				fn(readString, fname)
			}
		}
	}()
	return nil
}

type Cmder interface {
	OutPut() <-chan string
	IoStdout(ctx context.Context, sh string) (io.ReadCloser, error)
	IoFile(filePath string) (io.Reader, error)
	AsyncExecSh(ctx context.Context, ioR io.Reader) error
	AsyncExecFile(ctx context.Context, ioR io.Reader, fname string, fn fnCall) error
}
