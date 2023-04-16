package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/x2vpk7/
// 分数到小数

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func fractionToDecimal(numerator int, denominator int) string {
	// 先考虑是否能够整除
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	// 不能整除的话，判断正负
	s := []byte{}
	if (numerator < 0) != (denominator < 0) {
		s = append(s, '-')
	}
	// 整数部分
	numerator, denominator = abs(numerator), abs(denominator)
	integerPart := numerator / denominator
	s = append(s, []byte(strconv.Itoa(integerPart))...)
	s = append(s, '.')
	// 小数部分
	indexMap := map[int]int{} // 用来存储开始循环后的连续字符
	remainder := numerator % denominator
	for remainder != 0 && indexMap[remainder] == 0 {
		indexMap[remainder] = len(s)
		remainder *= 10
		s = append(s, '0'+byte(remainder/denominator)) // 整数转byte
		remainder %= denominator
	}
	if remainder > 0 { // 有循环部分
		insertIndex := indexMap[remainder]
		s = append(s[:insertIndex], append([]byte{'('}, s[insertIndex:]...)...)
		s = append(s, ')')
	}
	return string(s)
}

func main72() {
	fmt.Println(fractionToDecimal(10, 9))
}
