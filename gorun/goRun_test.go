package gorun

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGoRun(t *testing.T) {
	InitGoRuntime()
	for i := 0; i < 100; i++ {
		tmp := i
		Go(func(ctx context.Context) {
			ticker := time.NewTicker(3 * time.Second)
			defer func() {
				ticker.Stop()
				fmt.Println("ticker退出", tmp)
			}()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("exit", tmp)
					return
				case <-ticker.C:
					fmt.Println("定时")
				}
			}
		})
	}
	CloseGoRuntime()
}
