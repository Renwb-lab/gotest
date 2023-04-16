package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/x2rkbs/
// 单词接龙

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 广度搜索
	queue, visited := make([]string, 0), make(map[string]bool)
	queue = append(queue, beginWord)
	visited[beginWord] = true

	isAdj := func(s, w string) bool {
		if len(s) != len(w) {
			return false
		}
		t := 0
		for i := 0; i < len(s); i += 1 {
			if s[i] != w[i] {
				t += 1
			}
		}
		return t == 1
	}

	getAdj := func(s string) []string {
		ret := make([]string, 0)
		for _, w := range wordList {
			if _, ok := visited[w]; !ok {
				if isAdj(s, w) {
					ret = append(ret, w)
				}
			}
		}
		return ret
	}

	ret := 0
	for len(queue) > 0 {
		ret += 1
		newQueue := make([]string, 0)
		for _, v := range queue {
			if v == endWord {
				return ret
			}
			adjs := getAdj(v)
			for _, w := range adjs {
				visited[w] = true
				if w == endWord {
					return ret + 1
				}
			}
			newQueue = append(newQueue, adjs...)
		}
		queue = append([]string{}, newQueue...)
	}
	return 0
}

func main67() {
	beginWord, endWord := "hit", "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
