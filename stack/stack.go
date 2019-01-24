package stack

import (
	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/linkedlist"
)

// Stack ..
type Stack struct {
	list list.List
}

// New creates a new stack and returns it
func New() *Stack {
	return &Stack{
		list: linkedlist.New(),
	}
}

// Push the value to the top of stack
func (s *Stack) Push(value interface{}) *Stack {
	s.list.Append(value)
	return s
}

// Pop removes the element from the top of the stack, and returns it
func (s *Stack) Pop() interface{} {
	val, _ := s.list.Remove(s.list.Size() - 1)
	return val
}

// Peek returns the top element without removing it
func (s *Stack) Peek() interface{} {
	val, _ := s.list.Get(s.list.Size() - 1)
	return val
}

// Empty clears the stack
func (s *Stack) Empty() *Stack {
	s.list.Empty()
	return s
}

// Size returns the total values in the stack
func (s *Stack) Size() int {
	return s.list.Size()
}

// String returns the string representation of the stack
func (s *Stack) String() string {
	return s.list.String()
}
