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

// Insert a given value into the tree
func (bt *BinaryTree) Insert(value interface{}) {
	bt.Lock()
	defer bt.Unlock()

	bt.insert(value)
}

// Delete a node (using value) from the tree
func (bt *BinaryTree) Delete(value interface{}) bool {
	bt.Lock()
	defer bt.Unlock()

	return bt.delete(value)
}

// Contains return true if value exists in tree, otherwise false
func (bt *BinaryTree) Contains(value interface{}) bool {
	bt.RLock()
	defer bt.RUnlock()

	return bt.contains(value)
}

// Height returns the height of the tree (node/level count)
func (bt *BinaryTree) Height() int {
	bt.RLock()
	defer bt.RUnlock()

	return bt.height()
}

// Min returns the minimum value present in the tree
func (bt *BinaryTree) Min() (interface{}, bool) {
	bt.RLock()
	defer bt.RUnlock()

	return bt.min()
}

// Max returns the maximum value present in the tree
func (bt *BinaryTree) Max() (interface{}, bool) {
	bt.RLock()
	defer bt.RUnlock()

	return bt.max()
}

// Count returns the total number of nodes in tree
func (bt *BinaryTree) Count() int {
	bt.RLock()
	defer bt.RUnlock()

	return bt.count()
}

// Empty removes all the nodes from tree
func (bt *BinaryTree) Empty() {
	bt.Lock()
	defer bt.Unlock()

	bt.empty()
}

// InOrder returns a list.List with all values in-order
func (bt *BinaryTree) InOrder() list.List {
	bt.RLock()
	defer bt.RUnlock()

	return bt.inOrder()
}
