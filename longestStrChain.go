package main

import "fmt"

//  https://leetcode.cn/problems/longest-string-chain/?company_slug=ubiquanthr

// 1048. 最长字符串链
// 给出一个单词数组 words ，其中每个单词都由小写英文字母组成。

// 如果我们可以 不改变其他字符的顺序 ，在 wordA 的任何地方添加 恰好一个 字母使其变成 wordB ，那么我们认为 wordA 是 wordB 的 前身 。

// 例如，"abc" 是 "abac" 的 前身 ，而 "cba" 不是 "bcad" 的 前身
// 词链是单词 [word_1, word_2, ..., word_k] 组成的序列，k >= 1，其中 word1 是 word2 的前身，word2 是 word3 的前身，依此类推。一个单词通常是 k == 1 的 单词链 。

// 从给定单词列表 words 中选择单词组成词链，返回 词链的 最长可能长度 。

// 输入：words = ["a","b","ba","bca","bda","bdca"]
// 输出：4
// 解释：最长单词链之一为 ["a","ba","bda","bdca"]

// 输入：words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
// 输出：5
// 解释：所有的单词都可以放入单词链 ["xb", "xbc", "cxbc", "pcxbc", "pcxbcf"].

func longestStrChain(words []string) int {
	// 基本思路使用树的层次遍历
	shortLen, longLen := 1000, 0
	m := make(map[int][]string, 0)
	for i := range words {
		w := words[i]
		wl := len(w)
		if wl < shortLen {
			shortLen = wl
		}
		if wl > longLen {
			longLen = wl
		}

		if _, ok := m[wl]; !ok {
			m[wl] = make([]string, 0)
		}
		m[wl] = append(m[wl], w)
	}
	if longLen == shortLen {
		return 1
	}

	isBefore := func(beforeWord, afterword string) bool {
		idx := 0
		for ; idx < len(beforeWord); idx += 1 {
			if beforeWord[idx] != afterword[idx] {
				break
			}
		}
		if idx == len(beforeWord) {
			return true
		}
		bIdx, aIdx := idx, idx+1
		for ; aIdx < len(afterword); aIdx += 1 {
			if beforeWord[bIdx] != afterword[aIdx] {
				break
			}
			bIdx += 1
		}
		return aIdx == len(afterword)
	}
	getLongerWord := func(beforeWord string) []string {
		wl := len(beforeWord)
		wordList, ok := m[wl+1]
		if !ok {
			return []string{}
		}
		ret, newWordList := make([]string, 0), make([]string, 0)
		for _, afterword := range wordList {
			if isBefore(beforeWord, afterword) {
				ret = append(ret, afterword)
			} else {
				newWordList = append(newWordList, afterword)
			}
		}
		m[wl+1] = newWordList
		return ret
	}

	compute := func(length int) int {
		ret := 0
		queue := append([]string{}, m[length]...)
		for len(queue) > 0 {
			ret += 1
			newQueue := make([]string, 0)
			for i := range queue {
				w := queue[i]
				newQueue = append(newQueue, getLongerWord(w)...)
			}
			queue = newQueue
		}
		return ret
	}

	ret := 0
	// 这里需要注意：有些可能不是从最短的字符串开始的。
	for length := shortLen; length+ret <= longLen; length += 1 {
		temp := compute(length)
		if temp > ret {
			ret = temp
		}
	}

	return ret
}

func main37() {
	words := []string{
		// "a", "b", "ba", "bca", "bda", "bdca",
		// "xbc", "pcxbcf", "xb", "cxbc", "pcxbc",
		"ksqvsyq", "ks", "kss", "czvh", "zczpzvdhx", "zczpzvh",
		"zczpzvhx", "zcpzvh", "zczvh", "gr", "grukmj", "ksqvsq",
		"gruj", "kssq", "ksqsq", "grukkmj", "grukj", "zczpzfvdhx",
		"gru",
	}
	fmt.Println(longestStrChain(words))
}
