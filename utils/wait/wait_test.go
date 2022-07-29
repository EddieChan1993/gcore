package wait

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type testAa struct {
	waitGroup WaitGroupWrapper
}

func TestWaitGroupWrapper_Wrap(t *testing.T) {
	tw := testAa{
		waitGroup: WaitGroupWrapper{},
	}
	tw.waitGroup.Wrap(func() {
		fmt.Println("test")
	})
	ctxCh, _ := context.WithTimeout(context.Background(), 5*time.Second)
	tw.waitGroup.Wrap(func() {
		for {
			select {
			case <-ctxCh.Done():
				fmt.Println("退出")
				return
			}
		}
	})
	tw.waitGroup.Wait()
}
