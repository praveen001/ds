package redblacktree

import (
	"sync"

	"github.com/praveen001/ds/list"
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
	sync.RWMutex
}

// Node represents a node in red-black tree
type Node struct {
	height int
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
	}
}

// NewNode returns a new red-black tree node with given value
func newNode(value interface{}, parent *Node, color string) *Node {
	return &Node{
		height: 1,
		value:  value,
		parent: parent,
		color:  color,
	}
}

// Add a value into the tree
//
// Returns false is value already exists in tree, otherwise true
func (rbt *RedBlackTree) Add(value interface{}) bool {
	rbt.Lock()
	defer rbt.Unlock()

	return rbt.add(value)
}

// Remove a value from the tree
//
// Returns true if value was removed, otherwise false.
func (rbt *RedBlackTree) Remove(value interface{}) bool {
	rbt.Lock()
	defer rbt.Unlock()

	return rbt.remove(value)
}

// Height returns the height of the tree (node/level count) in O(1) Time Complexity.
func (rbt *RedBlackTree) Height() int {
	rbt.RLock()
	defer rbt.RUnlock()

	return rbt.height()
}

// Min returns the minimum value present in the tree
//
// Returns false if tree is empty
func (rbt *RedBlackTree) Min() (interface{}, bool) {
	rbt.RLock()
	defer rbt.RUnlock()

	return rbt.min()
}

// Max returns the maximum value present in the tree
//
// Returns false if tree is empty
func (rbt *RedBlackTree) Max() (interface{}, bool) {
	rbt.RLock()
	defer rbt.RUnlock()

	return rbt.max()
}

// Contains return true if value exists in tree, otherwise false
func (rbt *RedBlackTree) Contains(value interface{}) bool {
	rbt.RLock()
	defer rbt.RUnlock()

	return rbt.contains(value)
}

// Length returns the total number of nodes in tree
func (rbt *RedBlackTree) Length() int {
	rbt.RLock()
	defer rbt.RUnlock()

	return rbt.length()
}

// Clear all the nodes from tree
func (rbt *RedBlackTree) Clear() {
	rbt.Lock()
	defer rbt.Unlock()

	rbt.clear()
}

// InOrder returns a list.List with all values in-order
func (rbt *RedBlackTree) InOrder() list.List {
	rbt.RLock()
	defer rbt.RUnlock()

	return rbt.inOrder()
}

// PreOrder returns a list.List with all values in pre order
func (rbt *RedBlackTree) PreOrder() list.List {
	rbt.RLock()
	defer rbt.RUnlock()

	return rbt.preOrder()
}

// PostOrder returns a list.List with all values in post order
func (rbt *RedBlackTree) PostOrder() list.List {
	rbt.RLock()
	defer rbt.RUnlock()

	return rbt.postOrder()
}
