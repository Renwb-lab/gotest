package main

import "fmt"

// 2938. 区分黑球与白球
func minimumSteps(s string) int64 {
	ans := int64(0)
	sum := 0
	for i := 0; i < len(s); i += 1 {
		if s[i] == '1' {
			sum += 1 // 左边1的数量
		} else {
			// 将0移动到最右边
			ans += int64(sum)
		}
	}
	return ans
}

func main274() {
	fmt.Println(minimumSteps("111000100101"))
}
