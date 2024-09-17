package main

import "fmt"

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

func copyGraph(node *GNode) *GNode {
	mapping := map[*GNode]*GNode{}

	var dfs func(node *GNode) *GNode
	dfs = func(node *GNode) *GNode {
		if _, ok := mapping[node]; ok {
			return nil
		}
		cur := &GNode{val: node.val}
		mapping[node] = cur
		if len(node.nei) == 0 {
			return cur
		}
		l := len(node.nei)
		cur.nei = make([]*GNode, l, l)
		for i, c := range node.nei {
			var r *GNode
			if t, ok := mapping[c]; ok {
				r = t
			} else {
				r = dfs(c)
			}
			cur.nei[i] = r
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
func main071026() {
	node1 := &GNode{val: 1, nei: make([]*GNode, 0)}
	node2 := &GNode{val: 2, nei: make([]*GNode, 0)}
	node3 := &GNode{val: 3, nei: make([]*GNode, 0)}
	node4 := &GNode{val: 4, nei: make([]*GNode, 0)}
	node1.nei = append(node1.nei, node2, node4)
	node2.nei = append(node2.nei, node1, node3)
	node3.nei = append(node3.nei, node2, node4)
	node4.nei = append(node4.nei, node1, node3)
	res := copyGraph(node1)

	printGraph(res)
	fmt.Println()
	for i := range res.nei {
		printGraph(res.nei[i])
		fmt.Println()
	}
}
