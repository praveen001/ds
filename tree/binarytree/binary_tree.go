package binarytree

import (
	"sync"

	"github.com/praveen001/ds/list"

	"github.com/praveen001/ds/utils"
)

// BinaryTree represents a binary tree
type BinaryTree struct {
	root    *treeNode
	size    int
	compare utils.Comparator
	sync.RWMutex
}

// Node represents a node in a binary tree
type treeNode struct {
	value interface{}
	left  *treeNode
	right *treeNode
}

// New creates a new instance of binary tree and returns it
func New(c utils.Comparator) *BinaryTree {
	return &BinaryTree{
		compare: c,
	}
}

// NewNode returns a new binary tree node with given value
func newNode(value interface{}) *treeNode {
	return &treeNode{value, nil, nil}
}

// Insert given values into the tree
func (bt *BinaryTree) Insert(value interface{}) {
	bt.Lock()
	defer bt.Unlock()

	bt.insert(value)
}

// Delete a value from the tree
func (bt *BinaryTree) Delete(value interface{}) bool {
	bt.Lock()
	defer bt.Unlock()

	return bt.delete(value)
}

// Contains returns true if the given value exists in the tree, otherwise false
func (bt *BinaryTree) Contains(value interface{}) bool {
	bt.RLock()
	defer bt.RUnlock()

	return bt.contains(value)
}

// Height returns the height of the binary tree (using node/level counts)
func (bt *BinaryTree) Height() int {
	bt.RLock()
	defer bt.RUnlock()

	return bt.height()
}

// Min returns the minimum value from the tree
func (bt *BinaryTree) Min() interface{} {
	bt.RLock()
	defer bt.RUnlock()

	return bt.min()
}

// Max returns the maximum value from the tree
func (bt *BinaryTree) Max() interface{} {
	bt.RLock()
	defer bt.RUnlock()

	return bt.max()
}

// Count returns the total number of values in the tree
func (bt *BinaryTree) Count() int {
	bt.RLock()
	defer bt.RUnlock()

	return bt.count()
}

// Empty clears all the values in the tree
func (bt *BinaryTree) Empty() {
	bt.Lock()
	defer bt.Unlock()

	bt.empty()
}

// InOrder ..
func (bt *BinaryTree) InOrder() list.List {
	bt.RLock()
	defer bt.RUnlock()

	return bt.inOrder()
}
