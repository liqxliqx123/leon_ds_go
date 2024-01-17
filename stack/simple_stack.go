package stack

import "errors"

// SimpleStack ...
type SimpleStack struct {
	elements []any
}

func (s *SimpleStack) Push(i any) {
	s.elements = append(s.elements, i)
}

func (s *SimpleStack) Pop() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("empty stack")

	}
	last := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return last, nil
}

func (s *SimpleStack) Peek() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("empty stack")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *SimpleStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *SimpleStack) Size() int {
	return len(s.elements)
}

func NewSimpleStack() *SimpleStack {
	return &SimpleStack{
		elements: make([]any, 0),
	}
}
