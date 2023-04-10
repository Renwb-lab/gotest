package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xa9v8i/

// 单词拆分 II
// 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。
// 注意：词典中的同一个单词可能在分段中被重复使用多次。
//
// 示例 1：
// 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// 输出:["cats and dog","cat sand dog"]
// 示例 2：
// 输入:s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
// 输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
// 解释: 注意你可以重复使用字典中的单词。
// 示例 3：
// 输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// 输出:[]
//
// 提示：
// 1 <= s.length <= 20
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 10
// s 和 wordDict[i] 仅有小写英文字母组成
// wordDict 中所有字符串都 不同

func wordBreak2(s string, wordDict []string) []string {
	m := make(map[string]int)
	for _, word := range wordDict {
		if _, ok := m[word]; !ok {
			m[word] = 1
		}
	}
	size := len(s)
	ret, line := make([]string, 0), ""

	var dp func(idx int)
	dp = func(idx int) {
		if idx >= size {
			ret = append(ret, line[:len(line)-1])
			return
		}
		for i := idx + 1; i <= size; i += 1 {
			if _, ok := m[s[idx:i]]; ok {
				line += s[idx:i] + " "
				dp(i)
				line = line[:len(line)-i+idx-1]
			}
		}
	}
	dp(0)
	return ret
}

func main46() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	ret := wordBreak2(s, wordDict)
	for _, line := range ret {
		fmt.Println(line)
	}
}
