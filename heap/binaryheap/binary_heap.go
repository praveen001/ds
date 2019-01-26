package binaryheap

import (
	"github.com/praveen001/ds/utils"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/linkedlist"
)

// BinaryHeap is a binary heap implementation using linked list internally
type BinaryHeap struct {
	list    list.List
	compare utils.CompareFunc
}

// Config for binary heap
type Config struct {
	List list.List
}

// New returns a empty binary heap
func New(fn utils.CompareFunc) *BinaryHeap {
	return &BinaryHeap{
		list:    linkedlist.New(),
		compare: fn,
	}
}

// NewWithConfig returns a empty binary heap with config
func NewWithConfig(c *Config, fn utils.CompareFunc) *BinaryHeap {
	return &BinaryHeap{
		list:    c.List,
		compare: fn,
	}
}

// Push a given values into the tree
func (bh *BinaryHeap) Push(values ...interface{}) bool {
	return bh.push(values...)
}

// Pop a node (using value) from the tree
func (bh *BinaryHeap) Pop() (interface{}, bool) {
	return bh.pop()
}

// Peek returns the first value in heap without removing it
func (bh *BinaryHeap) Peek() (interface{}, bool) {
	return bh.peek()
}

// Clear removes all the values from heap
func (bh *BinaryHeap) Clear() {
	bh.clear()
}

// Length gives the number of values in heap
func (bh *BinaryHeap) Length() int {
	return bh.length()
}

// String returns the string representation of the binary heap
func (bh *BinaryHeap) String() string {
	return bh.list.String()
}
