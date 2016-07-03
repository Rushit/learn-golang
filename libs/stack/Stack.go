package stack

import (
	ll "../../libs/linkedlist"
)

type Stack struct {
	top *ll.LinkedList
	Size int
}

func (s *Stack) Push(data int)  {
	if s.top == nil {
		s.top = &ll.LinkedList{nil, data}
	} else {
		newNode := &ll.LinkedList{s.top, data}
		s.top = newNode
	}
	s.Size++
}

func (s *Stack) Pop() (int, error)  {
	if(s.top == nil) {
		return nil, error{"Stack is empty"}
	}
	result := s.top.Data
	s.top = s.top.Next
	s.Size--
	return result, nil
}

func (s *Stack) Peek() (int, error) {
	if(s.top == nil) {
		return nil, error{"Stack is empty"}
	}
	return s.top.Data, nil
}


