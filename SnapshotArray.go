package main

import (
	"fmt"
	"sort"
)

type SnapshotArray struct {
	snapCnt int
	data    [][][2]int
}

func Constructor3(length int) SnapshotArray {
	return SnapshotArray{
		snapCnt: 0,
		data:    make([][][2]int, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.data[index] = append(this.data[index], [2]int{this.snapCnt, val})
}

func (this *SnapshotArray) Snap() int {
	this.snapCnt++
	return this.snapCnt - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	x := sort.Search(len(this.data[index]), func(i int) bool {
		return this.data[index][i][0] > snap_id
	})
	if x == 0 {
		return 0
	}
	return this.data[index][x-1][1]
}

func main() {
	s := Constructor3(5)
	s.Set(0, 5)
	fmt.Println(s.Snap())
	s.Set(0, 6)
	fmt.Println(s.Get(0, 0))
}
