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
	mtx     sync.RWMutex
	sync    bool
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
		sync:    true,
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
	at.lock()
	defer at.unlock()

	return at.set(key, value)
}

// RSet a given value into the tree recursively
func (at *AvlTree) RSet(key, value interface{}) {
	at.lock()
	defer at.unlock()

	at.root = at.rset(at.root, key, value)
}

// Get a value by key from tree
//
// Returns value if key exists, otherwise it returns nil, false
func (at *AvlTree) Get(key interface{}) (interface{}, bool) {
	at.lock()
	defer at.unlock()

	return at.get(key)
}

// Remove a key from the tree
//
// Returns true if key was removed, otherwise false.
func (at *AvlTree) Remove(key interface{}) bool {
	at.lock()
	defer at.unlock()

	return at.remove(key)
}

// Min returns the minimum value present in the tree
//
// Returns false if tree is empty
func (at *AvlTree) Min() (ds.Node, bool) {
	at.rlock()
	defer at.runlock()

	return at.min()
}

// Max returns the maximum value present in the tree
//
// Returns false if tree is empty
func (at *AvlTree) Max() (ds.Node, bool) {
	at.rlock()
	defer at.runlock()

	return at.max()
}

// Contains return true if key exists in tree, otherwise false
func (at *AvlTree) Contains(key interface{}) bool {
	at.rlock()
	defer at.runlock()

	return at.contains(key)
}

// Len returns the total number of nodes in tree
func (at *AvlTree) Len() int {
	at.rlock()
	defer at.runlock()

	return at.len()
}

// Clear all the nodes from tree
func (at *AvlTree) Clear() {
	at.lock()
	defer at.unlock()

	at.clear()
}

// InOrder returns a ds.List with all values in-order
func (at *AvlTree) InOrder() ds.List {
	at.rlock()
	defer at.runlock()

	return at.inOrder()
}

// PreOrder returns a ds.List with all values in pre order
func (at *AvlTree) PreOrder() ds.List {
	at.rlock()
	defer at.runlock()

	return at.preOrder()
}

// PostOrder returns a ds.List with all values in post order
func (at *AvlTree) PostOrder() ds.List {
	at.rlock()
	defer at.runlock()

	return at.postOrder()
}

func (at *AvlTree) lock() {
	if at.sync {
		at.mtx.Lock()
	}
}

func (at *AvlTree) unlock() {
	if at.sync {
		at.mtx.Unlock()
	}
}

func (at *AvlTree) rlock() {
	if at.sync {
		at.mtx.RLock()
	}
}

func (at *AvlTree) runlock() {
	if at.sync {
		at.mtx.RUnlock()
	}
}
