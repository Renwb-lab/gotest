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

	retChan := make(chan int, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		retChan <- f(ctx)
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

func testFunc(ctx context.Context) int {
	for i := 0; i < 1000; i += 1 {
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

func main() {
	begin := time.Now()
	fmt.Println(RunFuncWithTime(testFunc, 10))
	fmt.Println(time.Now().Sub(begin).String())
	// time.Sleep(time.Second * time.Duration(10))
}
