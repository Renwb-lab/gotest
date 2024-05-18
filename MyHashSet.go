package main

import "container/list"

type MyHashSet struct {
	data []list.List
}

const base = 1 << 16

func Constructor() MyHashSet {
	return MyHashSet{data: make([]list.List, base)}
}

func (s *MyHashSet) hash(key int) int {
	return key & (base - 1)
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		idx := this.hash(key)
		this.data[idx].PushBack(key)
	}

}

func (this *MyHashSet) Remove(key int) {
	idx := this.hash(key)
	for it := this.data[idx].Front(); it != nil; it = e.Next() {
		if it.Value.(int) == key {
			this.data[idx].Remove(it)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	idx := this.hash(key)
	for it := this.data[idx].Front(); it != nil; it = e.Next() {
		if it.Value.(int) == key {
			return true
		}
	}
	return false
}
