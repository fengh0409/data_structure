package main

import "fmt"

const size int = 10

// 顺序栈
type Stack struct {
	data [size]int
	top  int
}

func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return 0, fmt.Errorf("stack is empty")
	}
	result := s.data[s.top]
	s.top--
	return result, nil
}

func (s *Stack) Push(e int) error {
	if s.top == size-1 {
		return fmt.Errorf("stack overflow")
	}
	s.data[s.top+1] = e
	s.top += 1
	return nil
}

// 链栈，必须有个栈顶指针top
type StackNode struct {
	data int
	next *StackNode
}
type ListStack struct {
	top *StackNode
}

func (l *ListStack) Pop() (int, error) {
	if l.top == nil {
		return 0, fmt.Errorf("empty stack")
	}

	top := l.top
	l.top = l.top.next
	return top.data, nil
}

func (l *ListStack) Push(e int) error {
	node := &StackNode{
		data: e,
	}
	if l.top != nil {
		node.next = l.top
	}
	l.top = node
	return nil
}
