package main

import (
	"fmt"
)

// 1 - 2 - 3
//. - 4 -

// 1 - 2

// class GNode(val: Int, nei: List<GNode>)

// fun copyGraph(node: GNode): GNode

// 遍历
// 图，深度遍历，广度遍历
// 重复遍历的问题，visited map

type GNode struct {
	val int
	nei []*GNode
}

func copyGraphBfs(node *GNode) *GNode {
	if node == nil {
		return nil
	}
	// 保存映射关系
	mapping := make(map[*GNode]*GNode)
	queue := []*GNode{node}
	// 构建Node
	res := &GNode{node.val, nil}
	mapping[node] = res
	// 存放到queue中的节点已完成构造
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, nex := range cur.nei {
			// 构建当前节点的nei
			if len(mapping[cur].nei) == 0 {
				mapping[cur].nei = make([]*GNode, 0)
			}
			// 如果当前的节点已经出现在mapping中，则表明已经访问过，则不需要重复访问
			// 只要将关系链接起来就可以了
			if n, ok := mapping[nex]; ok {
				mapping[cur].nei = append(mapping[cur].nei, n)
			} else {
				// 构建nei节点，更新mapping, 关联关系, 进队列
				n := &GNode{nex.val, nil}
				mapping[nex] = n
				mapping[cur].nei = append(mapping[cur].nei, n)
				queue = append(queue, nex)
			}
		}
	}
	return res
}

func copyGraphDfs(node *GNode) *GNode {
	var dfs func(node *GNode) *GNode
	mapping := make(map[*GNode]*GNode)
	dfs = func(node *GNode) *GNode {
		if node == nil {
			return node
		}
		if res, ok := mapping[node]; ok {
			return res
		}
		// 复制当前节点
		cur := &GNode{val: node.val}
		// 保存映射关系
		mapping[node] = cur
		if len(node.nei) == 0 {
			return cur
		}
		cur.nei = make([]*GNode, len(node.nei))
		for i, nex := range node.nei {
			cur.nei[i] = dfs(nex)
		}
		return cur
	}
	return dfs(node)
}

func printGraph(node *GNode) {
	visited := map[*GNode]struct{}{}

	var dfs func(node *GNode)
	dfs = func(node *GNode) {
		if _, ok := visited[node]; ok {
			return
		}
		visited[node] = struct{}{}
		fmt.Println(node.val)

		for _, c := range node.nei {
			c := c
			if _, ok := visited[c]; !ok {
				dfs(c)
			}
		}
	}

	dfs(node)
}

func main() {
	node1 := &GNode{val: 1, nei: make([]*GNode, 0)}
	node2 := &GNode{val: 2, nei: make([]*GNode, 0)}
	node3 := &GNode{val: 3, nei: make([]*GNode, 0)}
	node4 := &GNode{val: 4, nei: make([]*GNode, 0)}
	node1.nei = append(node1.nei, node2, node4)
	node2.nei = append(node2.nei, node1, node3)
	node3.nei = append(node3.nei, node2, node4)
	node4.nei = append(node4.nei, node1, node3)
	res := copyGraphDfs(node1)

	printGraph(res)
	fmt.Println()
	for i := range res.nei {
		printGraph(res.nei[i])
		fmt.Println()
	}
}
