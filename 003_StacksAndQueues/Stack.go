package ch03

import "errors"

type StackNode struct {
	data int
	next *StackNode
}

func NewStackNode(n int) *StackNode {
	l := new(StackNode)
	l.data = n
	return l
}

type Stack struct {
	top *StackNode
}

func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return 0, errors.New("Cannot pop a nil stack")
	}
	v := s.top
	s.top = v.next
	return v.data, nil
}

func (s *Stack) Push(n int) {
	newNode := NewStackNode(n)
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) Peek() (int, error) {
	if s.top == nil {
		return 0, errors.New("Cannot peek a nil stack")
	}
	v := s.top
	return v.data, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}