package main

import (
	"fmt"
	"time"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xa8q4g/
// 基本计算器 II

func calculate(s string) int {
	ret, num := 0, 0
	preSign := '+'
	helper := []int{}
	for i, c := range s {
		isDigit := c >= '0' && c <= '9'
		if isDigit {
			num = num*10 + int(c-'0')
		}
		if (i == len(s)-1) || (!isDigit && c != ' ') {
			switch preSign {
			case '+':
				helper = append(helper, num)
			case '-':
				helper = append(helper, -num)
			case '*':
				helper[len(helper)-1] *= num
			default:
				helper[len(helper)-1] /= num
			}
			num, preSign = 0, c
		}
	}
	for _, v := range helper {
		ret += v
	}
	return ret
}

func main91() {
	s := "282-1*2*13-30-2*2*2/2-95/5*2+55+804+3024"
	fmt.Println(calculate(s))
	time.Microsecond.Seconds()
}
