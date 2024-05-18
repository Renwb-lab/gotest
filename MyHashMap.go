package main

import (
	"container/list"
	"fmt"
)

const base = 10

type Node struct {
	key   int
	value int
}
type MyHashMap struct {
	m []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{
		m: make([]list.List, 1<<base),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	idx := this.hash(key)
	it := this.contain(idx, key)
	if it != nil {
		this.m[idx].Remove(it)
	}
	this.m[idx].PushBack(&Node{key, value})
	return
}

func (this *MyHashMap) Get(key int) int {
	idx := this.hash(key)
	it := this.contain(idx, key)
	if it == nil {
		return -1
	}
	return it.Value.(*Node).value
}

func (this *MyHashMap) Remove(key int) {
	idx := this.hash(key)
	it := this.contain(idx, key)
	if it != nil {
		this.m[idx].Remove(it)
	}
	return
}

func (this *MyHashMap) hash(key int) int {
	return key & ((1 << 8) - 1)
}

func (this *MyHashMap) contain(idx, key int) *list.Element {
	it := this.m[idx].Front()
	for it != nil {
		t := it.Value.(*Node)
		if t.key == key {
			return it
		}
		it = it.Next()
	}
	return nil
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
func main() {
	obj := Constructor()
	// key, value := 1, 2
	// obj.Put(key, value)
	// fmt.Println(obj.Get(key))
	// obj.Remove(key)
	// ["MyHashMap","remove","get","put","put","put","get","put","put","put","put"]
	// [[],          [14],    [4],[7,3],[11,1],[12,1],[7],[1,19],[0,3],[1,8],[2,6]]
	obj.Remove(14)
	fmt.Println(obj.Get(4))
	obj.Put(7, 3)
	obj.Put(11, 1)
	obj.Put(12, 1)
	fmt.Println(obj.Get(7))
	obj.Put(1, 19)
	obj.Put(0, 3)
	obj.Put(1, 8)
	obj.Put(2, 6)
}
