package main

import (
	"fmt"
	"sync"
)

// https://zhuanlan.zhihu.com/p/471490292
// 有三个函数，分别打印"cat", "fish", "dog"要求每一个函数都用一个goroutine，按照顺序打印100次。

var cat = make(chan struct{})
var fish = make(chan struct{})
var dog = make(chan struct{})

func Cat() {
	for i := 0; i < 100; i += 1 {
		<-dog
		fmt.Printf("%d %s\n", i, "cat")
		cat <- struct{}{}
	}
	fmt.Println("cat finish")
}

func Fish() {
	for i := 0; i < 100; i += 1 {
		<-cat
		fmt.Printf("%d %s\n", i, "fish")
		fish <- struct{}{}
	}
	fmt.Println("fish finish")
}

func Dog() {
	for i := 0; i < 100; i += 1 {
		<-fish
		fmt.Printf("%d %s\n", i, "dog")
		if i < 99 {
			dog <- struct{}{}
		}
	}
	fmt.Println("dog finish")
}

func main99() {
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		Cat()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		Fish()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		Dog()
		wg.Done()
	}()
	dog <- struct{}{}
}
