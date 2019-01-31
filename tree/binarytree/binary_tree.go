package binarytree

import (
	"sync"

	"github.com/praveen001/ds/ds"

	"github.com/praveen001/ds/utils"
)

// BinaryTree represents a binary tree
type BinaryTree struct {
	root    *Node
	size    int
	compare utils.CompareFunc
	mtx     sync.RWMutex
	sync    bool
}

// Node represents a node in a binary tree
type Node struct {
	key   interface{}
	value interface{}
	left  *Node
	right *Node
}

// New creates a new instance of binary tree and returns it
func New(c utils.CompareFunc) *BinaryTree {
	return &BinaryTree{
		compare: c,
		sync:    true,
	}
}

// NewNode returns a new binary tree node with given value
func newNode(key, value interface{}) *Node {
	return &Node{key, value, nil, nil}
}

// Set a value into the tree
//
// Returns false is value already exists in tree, otherwise true
func (bt *BinaryTree) Set(key, value interface{}) bool {
	bt.lock()
	defer bt.unlock()

	return bt.set(key, value)
}

// Get a value by key from tree
//
// Returns value if key exists, otherwise it returns nil, false
func (bt *BinaryTree) Get(key interface{}) (interface{}, bool) {
	bt.lock()
	defer bt.unlock()

	return bt.get(key)
}

// Remove a value from the tree
//
// Returns true if value was removed, otherwise false.
func (bt *BinaryTree) Remove(value interface{}) bool {
	bt.lock()
	defer bt.unlock()

	return bt.remove(value)
}

// Min returns the minimum value present in the tree
//
// Returns false if tree is empty
func (bt *BinaryTree) Min() (ds.Node, bool) {
	bt.rlock()
	defer bt.runlock()

	return bt.min()
}

// Max returns the maximum value present in the tree
//
// Returns false if tree is empty
func (bt *BinaryTree) Max() (ds.Node, bool) {
	bt.rlock()
	defer bt.runlock()

	return bt.max()
}

// Contains return true if value exists in tree, otherwise false
func (bt *BinaryTree) Contains(value interface{}) bool {
	bt.rlock()
	defer bt.runlock()

	return bt.contains(value)
}

// Len returns the total number of nodes in tree
func (bt *BinaryTree) Len() int {
	bt.rlock()
	defer bt.runlock()

	return bt.len()
}

// Clear all the nodes from tree
func (bt *BinaryTree) Clear() {
	bt.lock()
	defer bt.unlock()

	bt.clear()
}

// InOrder returns a ds.List with all values in-order
func (bt *BinaryTree) InOrder() ds.List {
	bt.rlock()
	defer bt.runlock()

	return bt.inOrder()
}

// PreOrder returns a ds.List with all values in pre order
func (bt *BinaryTree) PreOrder() ds.List {
	bt.rlock()
	defer bt.runlock()

	return bt.preOrder()
}

// PostOrder returns a ds.List with all values in post order
func (bt *BinaryTree) PostOrder() ds.List {
	bt.rlock()
	defer bt.runlock()

	return bt.postOrder()
}

func (bt *BinaryTree) lock() {
	if bt.sync {
		bt.mtx.Lock()
	}
}

func (bt *BinaryTree) unlock() {
	if bt.sync {
		bt.mtx.Unlock()
	}
}

func (bt *BinaryTree) rlock() {
	if bt.sync {
		bt.mtx.RLock()
	}
}

func (bt *BinaryTree) runlock() {
	if bt.sync {
		bt.mtx.RUnlock()
	}
}
