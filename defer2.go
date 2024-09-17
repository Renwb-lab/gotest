package main

import (
	"fmt"
	"time"
)

func computeCost() func() {
	pre := time.Now()
	return func() {
		elapsed := time.Since(pre)
		fmt.Println("elapsed:", elapsed)
	}
}

func main071029() {
	defer computeCost()()
	time.Sleep(time.Second)
}
