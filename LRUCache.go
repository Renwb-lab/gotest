package main

// https://leetcode.cn/problems/lru-cache/

type LRUNode struct {
	Key  int
	Val  int
	Next *LRUNode
	Pre  *LRUNode
}

type LRUCache struct {
	Head *LRUNode
	Tail *LRUNode
	Hash map[int]*LRUNode
	cap  int
}

func Constructor1(capacity int) LRUCache {
	head := &LRUNode{Key: 0, Val: 0, Next: nil, Pre: nil}
	tail := &LRUNode{Key: 0, Val: 0, Next: nil, Pre: nil}
	head.Next = tail
	tail.Pre = head
	return LRUCache{Head: head, Tail: tail, cap: capacity, Hash: make(map[int]*LRUNode)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Hash[key]; ok {
		this.removeNode(node)
		this.addNodeInHead(node)
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) removeNode(node *LRUNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = nil
	node.Pre = nil
}
func (this *LRUCache) addNodeInHead(node *LRUNode) {
	node.Next = this.Head.Next
	this.Head.Next.Pre = node

	this.Head.Next = node
	node.Pre = this.Head
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Hash[key]; ok { // 存在的话
		this.Hash[key].Val = value
		node := this.Hash[key]
		this.removeNode(node)
		this.addNodeInHead(node)
	} else { // 不存在
		if len(this.Hash) == this.cap { // 如果满了的话，需要移除最后一个节点
			delete(this.Hash, this.Tail.Pre.Key)
			this.removeNode(this.Tail.Pre)
		}
		node := &LRUNode{Key: key, Val: value, Next: nil, Pre: nil}
		this.Hash[key] = node
		this.addNodeInHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
