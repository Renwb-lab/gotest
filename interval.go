package main

import "fmt"

// 模拟栈的数据结构
// 包括入栈和出栈
// 有大小约束
// 类型：int

type CNode struct {
	val  int
	next *CNode
	pre  *CNode
}

type Stack struct {
	size int
	cap  int
	tail *CNode
}

func NewStack(cap int) *Stack {
	// head := &CNode{}
	tail := &CNode{}
	// head.next = tail
	// tail.pre = head

	// [...]->0
	return &Stack{
		cap:  cap,
		size: 0,
		tail: tail,
	}
}

func (s *Stack) Push(val int) error {
	if s.size == s.cap {
		return fmt.Errorf("full")
	}

	// [...] -> tail
	// [..pre n] -> tail

	n := &CNode{val: val, next: nil}
	pre := s.tail.pre // first Push, pre = nil
	if pre == nil {
		// pre -> n
		// pre <- n
		pre.next = n
		n.pre = pre
	}

	// n -> tail
	// n <- tail
	n.next = s.tail
	s.tail.pre = n

	s.size += 1
	return nil
}

// 1. 线程数量
// 		最小数量
// 		最大数量
// 2. 队列
//    先进先出
//    ...
// 3. 拒绝
//   [1, 2, 3] 4
//   {faststop -> strategy, }
//   {faststop -> strategy, }
// 	Register
//   {custom -> strategy, }
// 4. 停止
//    interface{}
// 5. 任务
//		正常逻辑
//      失败逻辑 callback
// 6. 周期性执行
// 7. 名称

// func sort(ctx contenxt, arr T[], cmp func(i, j int) bool, swap func (i, j int)) error

// A -> 5
// order [id, orderid, createtime]

// select count(1), day(createtime) from order
// where year(createime) >= lastyear() and year(creatime) < thisyear()
// group by day(createtime)
// order by day(createtime) desc

// func (s *Stack) push(n int) {
// 	// check full
// 	CNode* node = new CNode{n, s->head}
// 	s->head = node
// 	s->size += 1
// }

// func (s *Stack) pop() (int, error) {
// 	// check empty
// 	cur := s->head
// 	s->head = cur->next
// 	s->size -= 1
// 	return cur->val
// }
// 头插法

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, fmt.Errorf("empty")
	}
	// 0->1->2->3->0
	// 0 -> [...] -> n
	// [...] -> n

	// [1, 2, 3] -> tail
	// [1, 2] -> tail

	// pre cur tail
	cur := s.tail.pre
	pre := cur.pre // [n] -> tail
	ans := cur.val

	// pre -> tail
	// pre <- tail
	if pre != nil {
		pre.next = s.tail
	}
	s.tail.pre = pre

	// remove relation
	cur.next = nil
	cur.pre = nil
	s.size -= 1
	return ans, nil
}

func main256() {
	s := NewStack(2)
	err := s.Push(1)
	fmt.Println(err)
	err = s.Push(2)
	fmt.Println(err)
	err = s.Push(3)
	fmt.Println(err)

	v1, err := s.Pop()
	fmt.Println(v1, err)
	v2, err := s.Pop()
	fmt.Println(v2, err)
	v3, err := s.Pop()
	fmt.Println(v3, err)

	fmt.Println("Hello world.")
}
