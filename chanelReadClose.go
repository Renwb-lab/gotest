package main

import (
	"fmt"
	"time"
)

// 向一个关闭的ch进行对数据的话，返回结果为默认值。
func main95() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			c := <-ch
			fmt.Println(c)
			time.Sleep(1)
		}
	}()
	go func() {
		close(ch)
	}()
	time.Sleep(100000)
}
