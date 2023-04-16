package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xa6dkt/
// Excel表列序号

func titleToNumber(columnTitle string) int {
	ret := 0
	for _, c := range columnTitle {
		ret = ret*26 + int(c-'A'+1)
	}
	return ret
}

func main76() {
	fmt.Println(titleToNumber("AZ"))
}
