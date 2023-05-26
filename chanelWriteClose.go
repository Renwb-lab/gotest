package main

import "time"

// 向一个关闭的ch进行对数据的话，会panic.
func main94() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
			time.Sleep(1)
		}
	}()
	go func() {
		close(ch)
	}()
	time.Sleep(100000)
}
