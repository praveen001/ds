package stack

import (
	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/list/linkedlist"
)

// Stack ..
type Stack struct {
	list ds.List
}

// Config config for stack
type Config struct {
	List ds.List
}

// New creates a new stack and returns it
func New() *Stack {
	return &Stack{
		list: linkedlist.New(),
	}
}

// NewWithConfig creates a new stack with given config
func NewWithConfig(c *Config) *Stack {
	return &Stack{
		list: c.List,
	}
}

// Push the value to the top of stack
func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

// Pop removes the element from the top of the stack, and returns it
func (s *Stack) Pop() (interface{}, bool) {
	return s.list.Remove(s.list.Len() - 1)
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (interface{}, bool) {
	return s.list.Get(s.list.Len() - 1)
}

// Len returns the total values in the stack
func (s *Stack) Len() int {
	return s.list.Len()
}

// Clear the stack
func (s *Stack) Clear() {
	s.list.Clear()
}

// Values returns all the values in the stack as a slice
func (s *Stack) Values() []interface{} {
	return s.list.Values()
}

// String returns the string representation of the stack
func (s *Stack) String() string {
	return s.list.String()
}
