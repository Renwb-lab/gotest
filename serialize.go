package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/leetbook/read/top-interview-questions/xasfi3/
// 二叉树的序列化与反序列化

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type Codec struct {
}

func Constructor5() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	ret := strconv.Itoa(root.Val)
	ret += "_" + this.serialize(root.Left) // 注意不能跳过空节点
	ret += "_" + this.serialize(root.Right)
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var dp func() *TreeNode
	idx, size := 0, len(data)
	dp = func() *TreeNode {
		if idx > size {
			return nil
		}
		if data[idx] == '#' {
			idx += 2 // 后移下标
			return nil
		}
		pre := idx
		for idx < size && data[idx] != '_' {
			idx += 1
		}
		n, _ := strconv.Atoi(string([]byte(data)[pre:idx]))
		idx += 1 // 注意即使为空节点，也要更新idx的数值
		root := &TreeNode{Val: n}
		root.Left = dp()
		root.Right = dp()
		return root
	}
	return dp()
}

func main303() {
	str := "1_2_#_#_3_4_#_#_5_#_#"
	t := Constructor5()
	root := t.deserialize(str)
	fmt.Println(t.serialize(root))
}
