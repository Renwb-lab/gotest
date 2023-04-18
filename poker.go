package main

import (
	"fmt"
	"math/rand"
)

// 实现扑克牌洗牌函数
// 题目：

// 扑克牌游戏中需要实现一个shuffle函数用来打散扑克牌，可以使用系统random库

// 函数头举例：
func ShufflePoker() [10]string {
	var poker = [10]string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	}
	size := len(poker)
	// rand.Seed(int64(time.Now().Unix()))
	// seed:
	// [1, 23 3, 4 ,4 5 6 6,7  8, 3 , 4 ]
	for i := size - 1; i > 0; i -= 1 {
		idx := rand.Intn(i + 1)
		poker[idx], poker[i] = poker[i], poker[idx]
	}

	// TODO
	return poker
}

func main81() {
	nums := 10000000
	m := make(map[string]int)
	for i := 0; i < nums; i += 1 {
		ret := ShufflePoker()
		m[ret[0]] += 1
	}
	for k, v := range m {
		fmt.Printf("%s %d\n", k, v)
	}
	fmt.Println()
}
