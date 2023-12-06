package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main160() {
	// ctx := context.WithValue(context.TODO(), "key", "value")
	// fmt.Println(ctx.Value("key"))
	// fmt.Println(ctx.Value("key1"))
	// context.WithCancel(context.TODO())

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go handle(wg, ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err(), time.Now())
	}
	wg.Wait()
}

func handle(wg *sync.WaitGroup, ctx context.Context, duration time.Duration) {
	defer func() {
		fmt.Println("process end")
		wg.Done()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err(), time.Now())
	case <-time.After(duration):
		fmt.Println("process request with", duration, time.Now())
	}
}
