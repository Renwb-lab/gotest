package main

import "fmt"

// (200分) 输入一个字符串 , 只有数字 , 字母 , 正负号 , 小数点 组成 , 求该字符串中出现过长度最长的合法数字( 包含正负号 , 小数点 ) , 长度相同 , 输出任意一个最长的即可
// 示例 :
// 输入 : 1234567890abcd9.+12345.678.9ed
// 输出 : +12345.678

func longestNumber(instr string) string {
	idx, l := 0, len(instr)
	isNumFunc := func(c byte) bool {
		return c >= '0' && c <= '9'
	}
	numFunc := func() {
		for ; idx < l; idx += 1 {
			if !isNumFunc(instr[idx]) {
				break
			}
		}
	}
	maxLen, start, len := 0, 0, -1
	for idx < l {
		// 跳过非数字的字符 只有+/-/[0-9]能够过去
		for ; idx < l; idx += 1 {
			c := instr[idx]
			if c == '+' || c == '-' || isNumFunc(c) {
				break
			}
		}
		if idx == l {
			break
		}
		// 数字的话，就持续计算下去
		b, c := idx, instr[idx]
		if c == '-' || c == '+' {
			idx += 1
		}
		numFunc()
		if idx+1 < l && instr[idx] == '.' && isNumFunc(instr[idx+1]) {
			idx += 1
			numFunc()
		}
		if idx-b > maxLen {
			maxLen = idx - b
			start, len = b, idx-b
		}
	}
	ret := ""
	for i := 0; i < len; i += 1 {
		ret += string(instr[start+i])
	}
	return ret
}

func main17() {
	{
		s := "1234567890abcd9.+12345.678.9ed"
		r := longestNumber(s)
		fmt.Println(r)
	}
	{
		s := "abadad"
		r := longestNumber(s)
		fmt.Println(r)
	}
	{
		s := "a.0b0.0001a0.0000dad"
		r := longestNumber(s)
		fmt.Println(r)
	}
}
