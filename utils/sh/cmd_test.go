package sh

import (
	"context"
	"fmt"
	"io"
	"log"
	"testing"
)

func TestCmd_AsyncExecSh(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	ins := InitCmd()
	ioR, err := ins.IoStdout(ctx, "echo 123")
	if err != nil {
		log.Fatal(err)
	}
	err = ins.AsyncExecSh(ctx, ioR)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case <-ctx.Done():
			return
		case outPut := <-ins.OutPut():
			if outPut == io.EOF.Error() {
				return
			}
			fmt.Print(outPut)
		}
	}
}
