package avltree

import (
	"sync"

	"github.com/praveen001/ds/ds"

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
	key     interface{}
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
func newNode(key, value interface{}) *Node {
	return &Node{height: 1, key: key, value: value, left: nil, right: nil}
}

// Set a value into the tree
//
// Returns false is value already exists in tree, otherwise true
func (at *AvlTree) Set(key, value interface{}) bool {
	at.Lock()
	defer at.Unlock()

	return at.set(key, value)
}

// RSet a given value into the tree recursively
func (at *AvlTree) RSet(key, value interface{}) {
	at.Lock()
	defer at.Unlock()

	at.root = at.rset(at.root, key, value)
}

// Get a value by key from tree
//
// Returns value if key exists, otherwise it returns nil, false
func (at *AvlTree) Get(key interface{}) (interface{}, bool) {
	at.Lock()
	defer at.Unlock()

	return at.get(key)
}

// Remove a key from the tree
//
// Returns true if key was removed, otherwise false.
func (at *AvlTree) Remove(key interface{}) bool {
	at.Lock()
	defer at.Unlock()

	return at.remove(key)
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
func (at *AvlTree) Min() (ds.Node, bool) {
	at.RLock()
	defer at.RUnlock()

	return at.min()
}

// Max returns the maximum value present in the tree
//
// Returns false if tree is empty
func (at *AvlTree) Max() (ds.Node, bool) {
	at.RLock()
	defer at.RUnlock()

	return at.max()
}

// Contains return true if key exists in tree, otherwise false
func (at *AvlTree) Contains(key interface{}) bool {
	at.RLock()
	defer at.RUnlock()

	return at.contains(key)
}

// Len returns the total number of nodes in tree
func (at *AvlTree) Len() int {
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

// InOrder returns a ds.List with all values in-order
func (at *AvlTree) InOrder() ds.List {
	at.RLock()
	defer at.RUnlock()

	return at.inOrder()
}

// PreOrder returns a ds.List with all values in pre order
func (at *AvlTree) PreOrder() ds.List {
	at.RLock()
	defer at.RUnlock()

	return at.preOrder()
}

// PostOrder returns a ds.List with all values in post order
func (at *AvlTree) PostOrder() ds.List {
	at.RLock()
	defer at.RUnlock()

	return at.postOrder()
}
