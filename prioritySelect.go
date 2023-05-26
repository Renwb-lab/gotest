package main

import (
	"fmt"
	"time"
)

// https://zhuanlan.zhihu.com/p/471490292
// 当select监控多个chan同时到达就绪态时，如何先执行某个任务？

func priority_select(ch1, ch2 <-chan string) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val2 := <-ch2:
			// priority:
			// 	for {
			// 		select {
			// 		case val := <-ch1:
			// 			fmt.Println(val)
			// 		default:
			// 			break priority
			// 		}
			// 	}
			fmt.Println(val2)
		}
	}
}

func main104() {
	for i := 0; i < 10; i += 1 {
		ch1 := make(chan string)
		ch2 := make(chan string)

		go priority_select(ch1, ch2)

		go func() {
			ch2 <- "ch2"
		}()
		go func() {
			ch1 <- "ch1"
		}()
		time.Sleep(3 * time.Second)

		fmt.Println("===================")
	}
}
