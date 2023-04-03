package main

import (
	"fmt"
)

// 2. 有一个字符串 ，只有 ‘R’ 和 ‘B’ 组成 ，有如下规律 :
// 		n		s
// 		1		R
// 		2		BR
// 		3		RBBR
// 		4		BRRBRBBR
// 给出一个 n , 和 index , 求 ：第 n 个字符串(s) , 的第 index 个字母
// 示例1：
// n = 4    index = 6
// 输出 ：
// B

func RBChar(n, index int) byte {
	var dp func(n, index int) bool
	dp = func(n, index int) bool {
		if n == 1 && index == 0 {
			return true
		}
		// 计算出字符串的总长度
		len := 1 << (n - 1)
		if index >= len || index < 0 {
			return false
		}
		// 计算一半的长度
		half := len >> 1
		if index < half {
			// 如果小于一半的长度，就转化为计算n-1, index的字符串，然后翻转下。
			return !dp(n-1, index)
		} else {
			// 如果小于一半的长度，就转化为计算n-1, index-half的字符串，然后翻转下。
			return dp(n-1, index-half)
		}
	}
	if dp(n, index) {
		return 'R'
	}
	return 'B'
}

func main18() {
	fmt.Println(string(RBChar(4, 6)))
	fmt.Println(string(RBChar(4, 7)))
	fmt.Println(string(RBChar(4, 1)))
	fmt.Println(string(RBChar(4, 0)))
}
