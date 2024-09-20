package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

// func launch() {
// 	fmt.Println("launch")
// }

func main238() {
	abort := make(chan struct{}, 0)
	var wg sync.WaitGroup
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}

		select {
		case <-ctx.Done():
			fmt.Println("Finish")
			return
		}

	}()

	fmt.Println("Commencing countdown. Press return to abort")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}
