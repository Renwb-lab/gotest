package main

import (
	"fmt"
	"sort"
)

// (200分) 输入一个字符串 , 只有数字和字母组成 , 并保持格式 {要重复的字符}{重复次数} , 如输入 a2b3 输出 aabbb , 以此类推 , 解密该字符串
// 需保持如下规律 :
// a. 重复次数少的 , 排在前面 , 如 a3b2 , 应输出 bbaaa
// b. 重复次数相同的 , 字典序靠前的 , 排在前面 , 如 b2a2 , 应输出 aabb
// 示例 :
// 输入 : a4b2bac3bad3abcd2
// 输出 : abcdabcdbbbacbacbacbadbadbadaaaa

func repeatString(inStr string) string {
	// a4b2bac3bad3abcd2
	// abcdabcdbbbacbacbacbadbadbadaaaa
	type Temp struct {
		num int
		str string
	}
	helper := make([]Temp, 0)
	idx, l := 0, len(inStr)
	for idx < l {
		s := ""
		for ; idx < l && inStr[idx] >= 'a' && inStr[idx] <= 'z'; idx += 1 {
			s = s + string(inStr[idx])
		}
		n := 0
		for ; idx < l && inStr[idx] >= '0' && inStr[idx] <= '9'; idx += 1 {
			n = n*10 + int(inStr[idx]-'0')
		}
		helper = append(helper, Temp{num: n, str: s})
	}
	sort.Slice(helper, func(i, j int) bool {
		return helper[i].num < helper[j].num || (helper[i].num == helper[j].num && helper[i].str < helper[j].str)
	})
	ret := ""
	for i := range helper {
		t := helper[i]
		for j := 0; j < t.num; j += 1 {
			ret += t.str
		}
	}
	return ret
}

func main16() {
	s := "a4b2bac3bad3abcd2"
	r := repeatString(s)
	fmt.Println(r)
}
