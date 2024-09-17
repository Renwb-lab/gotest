package main

import "fmt"

type UnionFind struct {
	heights []int
	parents []int
	n       int
}

func NewUnionFind(n int) *UnionFind {
	heights, parents := make([]int, 0, n), make([]int, 0, n)
	for i := 0; i < n; i += 1 {
		heights = append(heights, 1)
		parents = append(parents, i)
	}
	return &UnionFind{
		n:       n,
		parents: parents,
		heights: heights,
	}
}

func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return
	}
	if uf.heights[px] == uf.heights[py] {
		uf.heights[px] += 1
		uf.parents[py] = px
	} else if uf.heights[px] < uf.heights[py] {
		uf.parents[px] = py
	} else {
		uf.parents[py] = px
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] == x {
		return x
	}
	res := uf.Find(uf.parents[x])
	uf.parents[x] = res
	return res
}

func (uf *UnionFind) Group(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func main326() {
	uf := NewUnionFind(10)
	uf.Union(1, 2)
	uf.Union(1, 3)
	fmt.Println(uf.Find(3))
	fmt.Println(uf.Group(1, 3))
	fmt.Println(uf.Group(2, 5))
	uf.Union(3, 7)
	fmt.Println(uf.Group(2, 7))
	fmt.Println(uf.Group(2, 8))
}
