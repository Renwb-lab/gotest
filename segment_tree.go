package main

import "fmt"

type operator func(int, int) int

type segmentNode struct {
	leftVal, rightVal   int
	leftNode, rightNode *segmentNode
	val                 int
}

func newSegmentNode(leftVal, rightVal int, val int) *segmentNode {
	return &segmentNode{
		leftVal:  leftVal,
		rightVal: rightVal,
	}
}

func (p *segmentNode) Query(tree *segmentTree, begin, end int) int {
	if p.leftVal >= begin && p.rightVal <= end {
		return p.val
	}
	mid := p.leftVal + (p.rightVal-p.leftVal)/2
	if p.leftNode == nil {
		p.leftNode = newSegmentNode(p.leftVal, mid, tree.val)
	}
	if p.rightNode == nil {
		p.rightNode = newSegmentNode(mid+1, p.rightVal, tree.val)
	}
	leftResult, rightResult := tree.val, tree.val
	if begin <= mid {
		leftResult = p.leftNode.Query(tree, begin, end)
	}
	if end > mid {
		rightResult = p.rightNode.Query(tree, begin, end)
	}

	return tree.getOp(leftResult, rightResult)
}

func (p *segmentNode) Insert(tree *segmentTree, loc, val int) {
	if p.leftVal >= loc && p.rightVal <= loc {
		p.val = tree.setOp(p.val, val)
		return
	}
	mid := p.leftVal + (p.rightVal-p.leftVal)/2
	if p.leftNode == nil {
		p.leftNode = newSegmentNode(p.leftVal, mid, tree.val)
	}
	if p.rightNode == nil {
		p.rightNode = newSegmentNode(mid+1, p.rightVal, tree.val)
	}
	if loc <= mid {
		p.leftNode.Insert(tree, loc, val)
	}
	if loc > mid {
		p.rightNode.Insert(tree, loc, val)
	}
	p.val = tree.setOp(p.val, val)
}

type segmentTree struct {
	root  *segmentNode
	setOp operator
	getOp operator
	val   int
}

func newSegmentTree(left, right, val int, setOp, getOp operator) *segmentTree {
	return &segmentTree{
		root:  newSegmentNode(left, right, val),
		setOp: setOp,
		getOp: getOp,
		val:   val,
	}
}

func (tree *segmentTree) Insert(loc int, val int) {
	tree.root.Insert(tree, loc, val)
}

func (tree *segmentTree) Query(begin int, end int) int {
	return tree.root.Query(tree, begin, end)
}

func main() {
	{
		max := func(oldValue, newValue int) int {
			if oldValue > newValue {
				return oldValue
			}
			return newValue
		}
		tree := newSegmentTree(-1000, 1000, 0, max, max)
		tree.Insert(1, 2)
		tree.Insert(3, 3)
		tree.Insert(10, 10)
		tree.Insert(10, 12)
		tree.Insert(-100, 100)
		fmt.Println(tree.Query(0, 1000))
		fmt.Println(tree.Query(-1000, 1000))
	}

	{
		cnt := func(oldValue, newValue int) int {
			return oldValue + 1
		}
		sum := func(leftRes, rightRes int) int {
			return leftRes + rightRes
		}

		tree := newSegmentTree(-100, 100, 0, cnt, sum)
		tree.Insert(1, 2)
		tree.Insert(3, 3)
		tree.Insert(10, 10)
		tree.Insert(10, 12)
		tree.Insert(-100, 100)
		fmt.Println(tree.Query(0, 100))
		fmt.Println(tree.Query(-100, 100))
	}
}
