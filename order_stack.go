package main

import "fmt"

type OrderedStack struct {
	stack []int
}

func NewOrderedStack() *OrderedStack {
	return &OrderedStack{
		stack: make([]int, 0),
	}
}

func (s *OrderedStack) Put(x int) int {
	for len(s.stack) > 0 && s.stack[len(s.stack)-1] < x {
		s.stack = s.stack[:len(s.stack)-1]
	}
	if len(s.stack) == 0 {
		s.stack = append(s.stack, x)
		return -1
	} else {
		r := s.stack[len(s.stack)-1]
		s.stack = append(s.stack, x)
		return r
	}
}

func main() {
	s := NewOrderedStack()
	arr := []int{3, 2, 1, 5, 6, 3}
	for _, x := range arr {
		fmt.Println(s.Put(x))
	}
}
