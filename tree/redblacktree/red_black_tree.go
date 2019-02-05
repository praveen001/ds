package redblacktree

import (
	"sync"

	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/utils"
)

const (
	red   = "red"
	black = "black"
)

// RedBlackTree represents a red-black tree
type RedBlackTree struct {
	root    *Node
	size    int
	compare utils.CompareFunc
	mtx     sync.RWMutex
	sync    bool
}

// Node represents a node in red-black tree
type Node struct {
	key    interface{}
	value  interface{}
	parent *Node
	left   *Node
	right  *Node
	color  string
}

// New creates a new instance of red-black tree and returns it
func New(c utils.CompareFunc) *RedBlackTree {
	return &RedBlackTree{
		compare: c,
		sync:    true,
	}
}

// NewNode returns a new red-black tree node with given value
func newNode(key, value interface{}, parent *Node, color string) *Node {
	return &Node{
		key:    key,
		value:  value,
		parent: parent,
		color:  color,
	}
}

// Set a value into the tree
//
// Returns false is value already exists in tree, otherwise true
func (rbt *RedBlackTree) Set(key, value interface{}) bool {
	rbt.lock()
	defer rbt.unlock()

	return rbt.set(key, value)
}

// Get a value by key from tree
//
// Returns value if key exists, otherwise it returns nil, false
func (rbt *RedBlackTree) Get(key interface{}) (interface{}, bool) {
	rbt.lock()
	defer rbt.unlock()

	return rbt.get(key)
}

// Remove a key from the tree
//
// Returns true if key was removed, otherwise false.
func (rbt *RedBlackTree) Remove(key interface{}) bool {
	rbt.lock()
	defer rbt.unlock()

	return rbt.remove(key)
}

// Min returns the minimum value present in the tree
//
// Returns false if tree is empty
func (rbt *RedBlackTree) Min() (ds.Node, bool) {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.min()
}

// Max returns the maximum value present in the tree
//
// Returns false if tree is empty
func (rbt *RedBlackTree) Max() (ds.Node, bool) {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.max()
}

// Contains return true if key exists in tree, otherwise false
func (rbt *RedBlackTree) Contains(key interface{}) bool {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.contains(key)
}

// Len returns the total number of nodes in tree
func (rbt *RedBlackTree) Len() int {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.len()
}

// Clear all the nodes from tree
func (rbt *RedBlackTree) Clear() {
	rbt.lock()
	defer rbt.unlock()

	rbt.clear()
}

// Keys returns a ds.List with all keys
func (rbt *RedBlackTree) Keys() ds.List {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.keys()
}

// Values returns a ds.List with all values
func (rbt *RedBlackTree) Values() ds.List {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.values()
}

// InOrder returns a ds.List with all values in-order
func (rbt *RedBlackTree) InOrder() ds.List {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.inOrder()
}

// PreOrder returns a ds.List with all values in pre order
func (rbt *RedBlackTree) PreOrder() ds.List {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.preOrder()
}

// PostOrder returns a ds.List with all values in post order
func (rbt *RedBlackTree) PostOrder() ds.List {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.postOrder()
}

func (rbt *RedBlackTree) lock() {
	if rbt.sync {
		rbt.mtx.Lock()
	}
}

func (rbt *RedBlackTree) unlock() {
	if rbt.sync {
		rbt.mtx.Unlock()
	}
}

func (rbt *RedBlackTree) rlock() {
	if rbt.sync {
		rbt.mtx.RLock()
	}
}

func (rbt *RedBlackTree) runlock() {
	if rbt.sync {
		rbt.mtx.RUnlock()
	}
}
