package main

import "fmt"

type BiGraph struct {
	n, m           int
	peoples, tasks []int
	relations      [][]int
	assigned       []int
}

// 给一个节点添加匹配
func (bg *BiGraph) AssignOne(i int, visited []bool) bool {
	for j := 0; j < bg.m; j += 1 {
		// 如任务tj没有访问过，并且pi可以解决tj
		if !visited[j] && bg.relations[i][j] != 0 {
			visited[j] = true         // 任务tj已经被访问
			if bg.assigned[j] == -1 { // 如果任务tj没有分配给其他人
				bg.assigned[j] = i // 将任务tj分配给pi
				return true
			}
			x := bg.assigned[j]           // 如果任务tj没有分配给其他人x
			if bg.AssignOne(x, visited) { // 如果能够将x重新分配任务，则
				bg.assigned[j] = i // 将任务tj分配给pi
				return true
			}
		}
	}
	return false
}

// 计算二分图的最大匹配
func (bg *BiGraph) AssignAll() int {
	ans := 0
	for i := 0; i < bg.n; i += 1 {
		visited := make([]bool, bg.m)
		if bg.AssignOne(i, visited) {
			ans += 1
		}
	}
	return ans
}

// 计算任务的分配情况
func (bg *BiGraph) GetAssign() map[int]int {
	ans := make(map[int]int)
	for i, j := range bg.assigned {
		if j != -1 {
			ans[i] = j
		}
	}
	return ans
}

func main() {
	// 主函数为上图中二分图示例的情况
	n, m := 3, 4
	assigned := make([]int, m)
	for i := 0; i < m; i += 1 {
		assigned[i] = -1
	}
	bg := &BiGraph{
		n: n,
		m: m,
		relations: [][]int{
			{1, 1, 0, 1},
			{1, 0, 1, 0},
			{0, 1, 0, 1},
		},
		assigned: assigned,
	}
	fmt.Printf("最大匹配数量为: %d \n", bg.AssignAll())
	for i, j := range bg.GetAssign() {
		fmt.Printf("第 %d 人分配到第 %d 任务\n", i, j)
	}
}
