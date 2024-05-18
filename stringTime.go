package main

import "fmt"

// (100分) 有一个表示内容的字符串 content , 和一个表示某个单词的字符串 word , 求 content 中 word 出现的次数 ,
// 注意 : word 中的字母可以任意排序
// 示例1 :
// content = "qweqqqew"   word = "qwe"
// 输出 : 3
// 示例2 :
// content = "abab"   word = "ab"
// 输出 : 3

func StringTimes(content, word string) int {
	need, visited := make(map[rune]int), make(map[rune]int)
	for _, c := range word {
		need[c] += 1
	}

	left, right, validNum := 0, 0, 0
	ret := 0
	for right < len(content) {
		c := rune(content[right])
		if _, ok := need[c]; ok {
			visited[c] += 1
			if visited[c] == need[c] {
				validNum += 1
			}
		}
		right += 1
		// 这里是连续的字符串，字符串字串
		for right-left == len(word) { // 特别重要
			if validNum == len(need) {
				ret += 1
			}
			c := rune(content[left])
			left += 1
			if _, ok := need[c]; ok {
				if _, ok := visited[c]; ok {
					visited[c] -= 1
					// 如果不满足条件的话，减少次数
					if visited[c] != need[c] {
						validNum -= 1
					}
				}
			}
		}
	}

	return ret
}

func main() {
	{
		content, word := "qweqwqew", "qwe"
		r := StringTimes(content, word)
		fmt.Println(r)
	}
	{
		content, word := "abab", "ab"
		r := StringTimes(content, word)
		fmt.Println(r)
	}
	{
		content, word := "abab", "c"
		r := StringTimes(content, word)
		fmt.Println(r)
	}
}
