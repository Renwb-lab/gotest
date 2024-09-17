package main

import "fmt"

func countCompleteSubstrings(word string, k int) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	check := func(arr []int) bool {
		for i := 0; i < 26; i += 1 {
			if arr[i] != 0 && arr[i] != k {
				return false
			}
		}
		return true
	}
	compute := func(b, e int) int {
		// 计算 abab 中 s 中每个字符 恰好 出现 k 次的数量
		// 字符均为小写字母，因此只有a-z 26种情况。 即出现单个字符，出现两个字符 ...
		// 即问题变成了abab中出现单个字符k次的数量，
		// 可以理解为长度为m*k的访问内，只出现了m个字符，且每个出现次数均为k -> 滑动串口问题
		ans := 0
		for m := 1; m <= 26; m += 1 {
			length := m * k
			if e-b < length {
				break
			}
			arr := make([]int, 26)
			for i := b; i < b+length && i < e; i += 1 {
				arr[int(word[i])-int('a')] += 1
			}
			if check(arr) {
				ans += 1
			}
			for i := b + length; i < e; i += 1 {
				arr[int(word[i])-int('a')] += 1
				arr[int(word[i-length])-int('a')] -= 1
				if check(arr) {
					ans += 1
				}
			}
		}
		return ans
	}
	l := len(word)
	res := 0
	for i := 0; i < l; {
		b := i
		// 先满足 相邻字符在字母表中的顺序 至多 相差 2，将问题简化
		for i += 1; i < l && abs(int(word[i])-int(word[i-1])) <= 2; i += 1 {
		}
		res += compute(b, i)
		// fmt.Println(word[b:i], res)
	}
	return res
}

func main071023() {
	fmt.Println(countCompleteSubstrings("gvgvvgv", 2))
}
