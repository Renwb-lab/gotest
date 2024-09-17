package main

import (
	"fmt"
)

// https://www.bilibili.com/video/BV19J411J7AZ/?p=5&spm_id_from=pageDriver&vd_source=f420410cb9f59880bca270f2acc60a94
func targan(v, e int, nodes []string, edges [][]int) [][]string {
	graph := make([][]int, v)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if len(graph[x]) == 0 {
			graph[x] = make([]int, 0)
		}
		graph[x] = append(graph[x], y)
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	res := [][]string{}
	time, st := 0, []int{}                     // time 记录访问时间， st 记录堆栈信息
	dfn, low := make([]int, v), make([]int, v) // dfn 记录真实访问时间，low记录调整后的时间
	var dfs func(n int)
	dfs = func(x int) {
		st = append(st, x)
		time += 1
		dfn[x], low[x] = time, time // 记录真实访问时间
		for _, y := range graph[x] {
			if dfn[y] == 0 { // 如果没有访问过则进行递归访问
				dfs(y)
			}
			low[x] = min(low[x], low[y]) // 调整更新后的时间
		}
		if dfn[x] == low[x] { // 如果真实访问时间和调整后的时间相同的话，则认为没有子节点可以和上游形成连通分量了
			line := []string{} // 进行出栈操作
			for len(st) > 0 && st[len(st)-1] != x {
				line = append(line, nodes[st[len(st)-1]])
				st = st[:len(st)-1]
			}
			if len(st) > 0 && st[len(st)-1] == x {
				line = append(line, nodes[st[len(st)-1]])
				st = st[:len(st)-1]
			}
			res = append(res, line)
		}
	}
	for x := 0; x < v; x += 1 {
		if dfn[x] == 0 {
			dfs(x)
		}
	}

	return res
}

func main318() {

	// 					a
	// 				/.	 \`   \.
	// 			b		  \`	 f
	// 		  /.  \`	   \`   /.
	// 		c		d	      g
	//		  \.  /.
	//			e
	//  /. 表示向下的箭头
	//  /` 表示向上的箭头

	v, e := 7, 9
	nodes := []string{"a", "b", "c", "d", "e", "f", "g"}
	egdes := [][]int{
		{0, 1}, {1, 2}, {2, 3}, {2, 4}, {3, 4}, {3, 1}, {0, 5}, {5, 6}, {6, 0},
	}
	res := targan(v, e, nodes, egdes)
	for i := range res {
		for j := range res[i] {
			fmt.Printf("%s\t", res[i][j])
		}
		fmt.Printf("\n")
	}
}
