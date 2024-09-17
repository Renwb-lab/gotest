package main

import "fmt"

type trieNode struct {
	children [26]*trieNode
	isEnd    bool // 标记是否结束
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: [26]*trieNode{},
		isEnd:    false,
	}
}

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{
		root: newTrieNode(),
	}
}

func (t *trie) Insert(words string) {
	p := t.root
	for _, w := range words {
		idx := w - 'a'
		if p.children[idx] == nil {
			p.children[idx] = newTrieNode()
		}
		p = p.children[idx]
	}
	p.isEnd = true
}

func (t *trie) Find(words string) bool {
	p := t.root
	for _, w := range words {
		idx := w - 'a'
		if p.children[idx] == nil {
			return false
		}
		p = p.children[idx]
	}
	return p != nil && p.isEnd
}

func main324() {
	t := newTrie()
	fmt.Println(t.Find("a"))
	t.Insert("abc")
	t.Insert("acd")
	fmt.Println(t.Find("ab"))
	fmt.Println(t.Find("abd"))
}
