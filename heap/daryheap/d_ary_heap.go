package daryheap

import (
	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/doublylinkedlist"
	"github.com/praveen001/ds/utils"
)

// DAryHeap is a d-ary heap implementation using linked list internally
type DAryHeap struct {
	list    list.List
	compare utils.CompareFunc
	d       int
	variant int
}

// Config for d-ary heap
type Config struct {
	List       list.List
	Comparator utils.CompareFunc
	D          int
	Variant    int
}

// New returns a empty d-ary heap
func New(fn utils.CompareFunc) *DAryHeap {
	return &DAryHeap{
		list:    doublylinkedlist.New(),
		compare: fn,
		d:       3,
		variant: 1,
	}
}

// NewWithConfig returns a empty d-ary heap with config
func NewWithConfig(c *Config) *DAryHeap {
	return &DAryHeap{
		list:    c.List,
		compare: c.Comparator,
		d:       c.D,
		variant: c.Variant,
	}
}

// Push a given values into the tree.
//
// Returns true if push succeeds.
func (dh *DAryHeap) Push(values ...interface{}) bool {
	return dh.push(values...)
}

// Pop removes first value from tree, and returns it.
//
// Second return value will be false if tree is empty.
func (dh *DAryHeap) Pop() (interface{}, bool) {
	return dh.pop()
}

// Peek returns the first value in heap without removing it
//
// Second return value will be false if tree is empty.
func (dh *DAryHeap) Peek() (interface{}, bool) {
	return dh.peek()
}

// Clear removes all the values from heap
func (dh *DAryHeap) Clear() {
	dh.clear()
}

// Len gives the number of values in heap
func (dh *DAryHeap) Len() int {
	return dh.len()
}

// String returns the string representation of the heap
func (dh *DAryHeap) String() string {
	return dh.list.String()
}
