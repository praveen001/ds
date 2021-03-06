package binaryheap

import (
	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/utils"

	"github.com/praveen001/ds/list/linkedlist"
)

// BinaryHeap is a binary heap implementation using linked list internally.
type BinaryHeap struct {
	list    ds.List
	compare utils.CompareFunc
	variant int
}

// Config for binary heap.
type Config struct {
	List       ds.List
	Comparator utils.CompareFunc
	Variant    int
}

// New returns a empty binary heap.
func New(fn utils.CompareFunc) *BinaryHeap {
	return &BinaryHeap{
		list:    linkedlist.New(),
		compare: fn,
		variant: ds.MinHeap,
	}
}

// NewWithConfig returns a empty binary heap with config.
func NewWithConfig(c *Config) *BinaryHeap {
	return &BinaryHeap{
		list:    c.List,
		compare: c.Comparator,
		variant: c.Variant,
	}
}

// Push a given values into the tree.
//
// Returns true if push succeeds.
func (bh *BinaryHeap) Push(values ...interface{}) bool {
	return bh.push(values...)
}

// Pop removes first value from tree, and returns it.
//
// Second return value will be false if tree is empty.
func (bh *BinaryHeap) Pop() (interface{}, bool) {
	return bh.pop()
}

// Peek returns the first value in heap without removing it
//
// Second return value will be false if tree is empty.
func (bh *BinaryHeap) Peek() (interface{}, bool) {
	return bh.peek()
}

// Clear removes all the values from heap
func (bh *BinaryHeap) Clear() {
	bh.clear()
}

// Len gives the number of values in heap
func (bh *BinaryHeap) Len() int {
	return bh.len()
}

// String returns the string representation of the heap
func (bh *BinaryHeap) String() string {
	return bh.string()
}
