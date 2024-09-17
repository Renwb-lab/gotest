package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// https://zhuanlan.zhihu.com/p/471490292
// 启动 2个groutine 2秒后取消， 第一个协程1秒执行完，第二个协程3秒执行完。

func ExistWithSecond(ctx context.Context, n int) {
	for i := 0; i < n; i += 1 {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		default:
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("exist with %d second\n", n)
}

func main071021() {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() {
		ExistWithSecond(ctx, 1)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		ExistWithSecond(ctx, 3)
		wg.Done()
	}()
}
