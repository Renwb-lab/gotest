package main

// 实现 Trie (前缀树)

// 发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

// 请你实现 Trie 类：

// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

type Trie struct {
	IsEnd bool
	Next  []*Trie
}

func Constructor15() Trie {
	return Trie{
		IsEnd: false,
		Next:  make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	p := this
	for _, c := range word {
		if p.Next[c-'a'] == nil {
			node := Constructor15()
			p.Next[c-'a'] = &node
		}
		p = p.Next[c-'a']
	}
	p.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	p := this
	for _, c := range word {
		if p.Next[c-'a'] == nil {
			return false
		}
		p = p.Next[c-'a']
	}
	return p.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for _, c := range prefix {
		if p.Next[c-'a'] == nil {
			return false
		}
		p = p.Next[c-'a']
	}
	return true
}
