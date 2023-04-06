package main

// https://leetcode.cn/problems/lru-cache/

type Node struct {
	Key  int
	Val  int
	Next *Node
	Pre  *Node
}

type LRUCache struct {
	Head *Node
	Tail *Node
	Hash map[int]*Node
	cap  int
}

func Constructor1(capacity int) LRUCache {
	head := &Node{Key: 0, Val: 0, Next: nil, Pre: nil}
	tail := &Node{Key: 0, Val: 0, Next: nil, Pre: nil}
	head.Next = tail
	tail.Pre = head
	return LRUCache{Head: head, Tail: tail, cap: capacity, Hash: make(map[int]*Node)}
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

func (this *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	node.Next = nil
	node.Pre = nil
}
func (this *LRUCache) addNodeInHead(node *Node) {
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
		node := &Node{Key: key, Val: value, Next: nil, Pre: nil}
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
