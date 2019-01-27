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

// Add a value into the tree
//
// Returns false is value already exists in tree, otherwise true
func (at *AvlTree) Add(value interface{}) bool {
	at.Lock()
	defer at.Unlock()

	return at.add(value)
}

// RAdd a given value into the tree recursively
func (at *AvlTree) RAdd(value interface{}) {
	at.Lock()
	defer at.Unlock()

	at.root = at.radd(at.root, value)
}

// Remove a value from the tree
//
// Returns true if value was removed, otherwise false.
func (at *AvlTree) Remove(value interface{}) bool {
	at.Lock()
	defer at.Unlock()

	return at.remove(value)
}

// Height returns the height of the tree (node/level count) in O(1) Time Complexity.
func (at *AvlTree) Height() int {
	at.RLock()
	defer at.RUnlock()

	return at.height()
}

// Min returns the minimum value present in the tree
//
// Returns false if tree is empty
func (at *AvlTree) Min() (interface{}, bool) {
	at.RLock()
	defer at.RUnlock()

	return at.min()
}

// Max returns the maximum value present in the tree
//
// Returns false if tree is empty
func (at *AvlTree) Max() (interface{}, bool) {
	at.RLock()
	defer at.RUnlock()

	return at.max()
}

// Contains return true if value exists in tree, otherwise false
func (at *AvlTree) Contains(value interface{}) bool {
	at.RLock()
	defer at.RUnlock()

	return at.contains(value)
}

// Length returns the total number of nodes in tree
func (at *AvlTree) Length() int {
	at.RLock()
	defer at.RUnlock()

	return at.length()
}

// Clear all the nodes from tree
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

// PreOrder returns a list.List with all values in pre order
func (at *AvlTree) PreOrder() list.List {
	at.RLock()
	defer at.RUnlock()

	return at.preOrder()
}

// PostOrder returns a list.List with all values in post order
func (at *AvlTree) PostOrder() list.List {
	at.RLock()
	defer at.RUnlock()

	return at.postOrder()
}
