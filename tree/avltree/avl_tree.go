package avltree

import (
	"sync"

	"github.com/praveen001/ds/list"

	"github.com/praveen001/ds/utils"
)

// AvlTree represents a binary tree
type AvlTree struct {
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
func New(c utils.Comparator) *AvlTree {
	return &AvlTree{
		compare: c,
	}
}

// NewNode returns a new binary tree node with given value
func newNode(value interface{}) *treeNode {
	return &treeNode{value, nil, nil}
}

// Insert a given value into the tree
func (at *AvlTree) Insert(value interface{}) {
	at.Lock()
	defer at.Unlock()

	at.insert(value)
}

// Delete a node (using value) from the tree
func (at *AvlTree) Delete(value interface{}) bool {
	at.Lock()
	defer at.Unlock()

	return at.delete(value)
}

// Contains return true if value exists in tree, otherwise false
func (at *AvlTree) Contains(value interface{}) bool {
	at.RLock()
	defer at.RUnlock()

	return at.contains(value)
}

// Height returns the height of the tree (node/level count)
func (at *AvlTree) Height() int {
	at.RLock()
	defer at.RUnlock()

	return at.height()
}

// Min returns the minimum value present in the tree
func (at *AvlTree) Min() (interface{}, bool) {
	at.RLock()
	defer at.RUnlock()

	return at.min()
}

// Max returns the maximum value present in the tree
func (at *AvlTree) Max() (interface{}, bool) {
	at.RLock()
	defer at.RUnlock()

	return at.max()
}

// Count returns the total number of nodes in tree
func (at *AvlTree) Count() int {
	at.RLock()
	defer at.RUnlock()

	return at.count()
}

// Empty removes all the nodes from tree
func (at *AvlTree) Empty() {
	at.Lock()
	defer at.Unlock()

	at.empty()
}

// InOrder returns a list.List with all values in-order
func (at *AvlTree) InOrder() list.List {
	at.RLock()
	defer at.RUnlock()

	return at.inOrder()
}
