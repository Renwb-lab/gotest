package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

// 1. 启动一个协程，每隔一秒打印"abcdefg"中的一个字符。主协程在三秒后退出。
func Worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	s := "abcdefg"
	for _, c := range s {
		select {
		case t := <-ctx.Done():
			fmt.Println(t)
			fmt.Println(unsafe.Sizeof(t))
			return
		default:
			time.Sleep(time.Second * time.Duration(1))
			fmt.Println(string(c))
		}
	}
}

func MainWorker() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(3))
	// 特别注意： defer cancel 的位置
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	defer func() {
		wg.Wait()
	}()

	go Worker(ctx, &wg)
}

func main89() {
	MainWorker()
	fmt.Println("main finished")
}
