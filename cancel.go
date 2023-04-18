package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func RunFuncWithTime(f func(ctx context.Context) int, runtime int) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(runtime))
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	retChan := make(chan int, 1)
	go func() {
		retChan <- f(ctx)
		wg.Done()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("ctx Done")
		return -1
	case ret := <-retChan:
		fmt.Println("call success")
		return ret
	}
}

func testFunc(ctx context.Context) int {
	for i := 0; i < 10000000; i += 1 {
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}
	fmt.Println("testFunc")
	return 1
}

func main84() {
	begin := time.Now()
	fmt.Println(RunFuncWithTime(testFunc, 10))
	fmt.Println(time.Now().Sub(begin).String())
	time.Sleep(time.Second * time.Duration(10))
}
