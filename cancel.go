package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func RunFuncWithTime(f func(ctx context.Context) int, runtime int) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(runtime))
	var wg sync.WaitGroup
	defer wg.Wait()
	defer cancel()

	retChan := make(chan int, 1) // 注意这里是带缓冲带channel, 防止启动的goroutine堵塞
	wg.Add(1)
	go func() {
		defer wg.Done()
		retChan <- f(ctx)
		close(retChan) // goroutine执行结束后关闭channel
		fmt.Println("goroutine finish")
	}()
	select {
	case <-ctx.Done():
		fmt.Println("ctx Done in RunFuncWithTime")
		return -2
	case ret := <-retChan:
		fmt.Println("call success")
		return ret
	}
}

func RunFuncWithTimeV2(f func(ctx context.Context) int, runtime int) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(runtime))
	defer cancel()

	retChan := make(chan int)
	go func() {
		retChan <- f(ctx)
		fmt.Println("goroutinue finish")
	}()
	select {
	case <-ctx.Done():
		fmt.Println("ctx Done in RunFuncWithTime")
		return -1
	case ret := <-retChan:
		fmt.Println("call success")
		return ret
	}
}

func testFunc(loop int) func(ctx context.Context) int {
	return func(ctx context.Context) int {
		for i := 0; i < loop; i += 1 {
			select {
			case <-ctx.Done():
				fmt.Println("ctx Done in testFunc")
				return -1
			default:
				time.Sleep(time.Second)
			}
		}
		fmt.Println("testFunc")
		return 1
	}
}

func main() {
	begin := time.Now()
	fmt.Println(RunFuncWithTime(testFunc(5), 2))
	fmt.Println(time.Now().Sub(begin).String())

	// time.Sleep(8 * time.Second)
}
