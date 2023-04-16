package main

import "fmt"

// https://leetcode.cn/leetbook/read/top-interview-questions/x2a743/
// 课程表II
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for i := 0; i < numCourses; i += 1 {
		graph = append(graph, make([]int, 0))
	}
	for _, p := range prerequisites {
		// 先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
		first, second := p[1], p[0]
		graph[first] = append(graph[first], second)
		inDegree[second] += 1
	}
	// 寻找入读为0的节点
	queue := make([]int, 0)
	for i := 0; i < numCourses; i += 1 {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 广度遍历
	visited := make(map[int]bool)
	newQueue := make([]int, 0)
	ret := make([]int, 0)
	for len(queue) > 0 {
		for _, cur := range queue {
			visited[cur] = true
			ret = append(ret, cur)
			for _, next := range graph[cur] {
				inDegree[next] -= 1
			}
		}
		for i := 0; i < numCourses; i += 1 {
			if _, ok := visited[i]; ok {
				continue
			}
			if inDegree[i] == 0 {
				newQueue = append(newQueue, i)
			}
		}
		queue = append([]int{}, newQueue...)
		newQueue = make([]int, 0)
	}

	if len(visited) == numCourses {
		return ret
	}
	return []int{}
}

func main69() {
	numCourses := 2
	prerequisites := [][]int{
		{1, 0},
	}
	ret := findOrder(numCourses, prerequisites)
	for _, v := range ret {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}
