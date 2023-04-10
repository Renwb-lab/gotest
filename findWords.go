package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/xaorig/
// 单词搜索 II

// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。

// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

func findWords(board [][]byte, words []string) []string {
	l, r := len(board), len(board[0])
	forwards := [][]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}
	ret := make([]string, 0)
	var dp func(x, y int, word string, idx int) bool
	dp = func(x, y int, word string, idx int) bool {
		if idx == len(word)-1 {
			return true
		}
		board[x][y] = '#'
		flag := false
		for _, f := range forwards {
			xx, yy := x+f[0], y+f[1]
			if xx >= 0 && xx < l && yy >= 0 && yy < r && board[xx][yy] == word[idx+1] {
				if dp(xx, yy, word, idx+1) {
					flag = true
					break
				}
			}
		}
		board[x][y] = word[idx]
		return flag
	}
	isOccur := func(word string) bool {
		for i := 0; i < l; i += 1 {
			for j := 0; j < r; j += 1 {
				if board[i][j] == word[0] {
					if dp(i, j, word, 0) {
						return true
					}
				}
			}
		}
		return false
	}
	for _, w := range words {
		if isOccur(w) {
			ret = append(ret, w)
		}
	}
	return ret
}

type TTrie struct {
	Word string
	Next []*TTrie
}

func NewTTrie() TTrie {
	return TTrie{
		Word: "",
		Next: make([]*TTrie, 26),
	}
}

func (t *TTrie) Insert(w string) {
	p := t
	for _, c := range w {
		if p.Next[c-'a'] == nil {
			node := NewTTrie()
			p.Next[c-'a'] = &node
		}
		p = p.Next[c-'a']
	}
	p.Word = w
}

func findWords2(board [][]byte, words []string) []string {
	l, r := len(board), len(board[0])
	forwards := [][]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	t := NewTTrie()
	for _, w := range words {
		t.Insert(w)
	}

	seen := make(map[string]bool)
	var dp func(node *TTrie, x, y int)
	dp = func(node *TTrie, x, y int) {
		ch := board[x][y]
		node = node.Next[ch-'a']
		if node == nil {
			return
		}
		if node.Word != "" {
			seen[node.Word] = true
		}
		board[x][y] = '#'
		for _, f := range forwards {
			xx, yy := x+f[0], y+f[1]
			if xx >= 0 && xx < l && yy >= 0 && yy < r && board[xx][yy] != '#' {
				dp(node, xx, yy)
			}
		}
		board[x][y] = ch
	}
	for i := 0; i < l; i += 1 {
		for j := 0; j < r; j += 1 {
			dp(&t, i, j)
		}
	}

	ret := make([]string, 0, len(seen))
	for w := range seen {
		ret = append(ret, w)
	}

	return ret
}

func main47() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	ret := findWords2(board, words)
	for _, line := range ret {
		fmt.Println(line)
	}
}
