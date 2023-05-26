package main

import (
	"fmt"
	"sync"
)

// https://zhuanlan.zhihu.com/p/471490292
// 两个协程交替打印10个字母和数字

var word = make(chan struct{})
var num = make(chan struct{})

func PrintlnNums() {
	for i := 0; i < 10; i += 1 {
		<-word
		fmt.Println(i)
		num <- struct{}{}
	}
}

func PrintlnWord() {
	for i := 0; i < 10; i += 1 {
		<-num
		fmt.Println("Hello World")
		if i < 9 {
			word <- struct{}{}
		}
	}
}

func main100() {
	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() {
		PrintlnNums()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		PrintlnWord()
		wg.Done()
	}()

	word <- struct{}{}
}
