package main

import (
	"fmt"
	"sync"
	"time"
)

func main071016() {
	c := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		c <- 1
		fmt.Println("1 end")
		wg.Done()
	}()
	go func() {
		c <- 2
		fmt.Println("2 end")
		wg.Done()
	}()
	go func() {
		c <- 3
		fmt.Println("3 end")
		wg.Done()
	}()
	time.Sleep(time.Duration(1000))
	fmt.Println(<-c)
	wg.Wait()

}
