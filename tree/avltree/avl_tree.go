package avltree

import (
	"sync"

	"github.com/praveen001/ds/list"

	"github.com/praveen001/ds/utils"
)

// AvlTree represents a binary tree
type AvlTree struct {
	root    *Node
	size    int
	compare utils.CompareFunc
	sync.RWMutex
}

// Node represents a node in a binary tree
type Node struct {
	height  int
	bFactor int
	value   interface{}
	left    *Node
	right   *Node
}

// New creates a new instance of binary tree and returns it
func New(c utils.CompareFunc) *AvlTree {
	return &AvlTree{
		compare: c,
	}
}

// NewNode returns a new binary tree node with given value
func newNode(value interface{}) *Node {
	return &Node{height: 1, value: value, left: nil, right: nil}
}

// Insert a given value into the tree
func (at *AvlTree) Insert(value interface{}) bool {
	at.Lock()
	defer at.Unlock()

	return at.insert(value)
}

// RInsert a given value into the tree recursively
func (at *AvlTree) RInsert(value interface{}) {
	at.Lock()
	defer at.Unlock()

	at.root = at.rinsert(at.root, value)
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

// Length returns the total number of nodes in tree
func (at *AvlTree) Length() int {
	at.RLock()
	defer at.RUnlock()

	return at.length()
}

// Clear removes all the nodes from tree
func (at *AvlTree) Clear() {
	at.Lock()
	defer at.Unlock()

	at.clear()
}

// InOrder returns a list.List with all values in-order
func (at *AvlTree) InOrder() list.List {
	at.RLock()
	defer at.RUnlock()

	return at.inOrder()
}
