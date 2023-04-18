package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xaqlgj/
// 逆波兰表达式求值

// 输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// 输出：22
// 解释：该算式转化为常见的中缀算术表达式为：
//   ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22

func evalRPN(tokens []string) int {
	helper := make([]int, 0)
	for i := range tokens {
		express := tokens[i]
		c, err := strconv.Atoi(tokens[i])
		if err != nil {
			size := len(helper)
			a, b := helper[size-2], helper[size-1]
			helper = helper[:size-2]
			if express == "-" {
				c = a - b
			} else if express == "+" {
				c = a + b
			} else if express == "*" {
				c = a * b
			} else {
				c = a / b
			}
			helper = append(helper, c)
		} else {
			helper = append(helper, c)
		}
	}
	return helper[0]
}

func main83() {
	tokens := []string{
		"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+",
	}
	fmt.Println(evalRPN(tokens))
}
