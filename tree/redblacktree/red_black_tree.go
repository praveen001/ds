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
func newNode(value interface{}, parent *Node, color string) *Node {
	return &Node{
		value:  value,
		parent: parent,
		color:  color,
	}
}

// Add a value into the tree
//
// Returns false is value already exists in tree, otherwise true
func (rbt *RedBlackTree) Add(value interface{}) bool {
	rbt.lock()
	defer rbt.unlock()

	return rbt.add(value)
}

// Remove a value from the tree
//
// Returns true if value was removed, otherwise false.
func (rbt *RedBlackTree) Remove(value interface{}) bool {
	rbt.lock()
	defer rbt.unlock()

	return rbt.remove(value)
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

// Contains return true if value exists in tree, otherwise false
func (rbt *RedBlackTree) Contains(value interface{}) bool {
	rbt.rlock()
	defer rbt.runlock()

	return rbt.contains(value)
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
