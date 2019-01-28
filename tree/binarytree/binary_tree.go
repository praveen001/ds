package binarytree

import (
	"sync"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/tree"

	"github.com/praveen001/ds/utils"
)

// BinaryTree represents a binary tree
type BinaryTree struct {
	root    *Node
	size    int
	compare utils.CompareFunc
	sync.RWMutex
}

// Node represents a node in a binary tree
type Node struct {
	height int
	key    interface{}
	value  interface{}
	left   *Node
	right  *Node
}

// New creates a new instance of binary tree and returns it
func New(c utils.CompareFunc) *BinaryTree {
	return &BinaryTree{
		compare: c,
	}
}

// NewNode returns a new binary tree node with given value
func newNode(key, value interface{}) *Node {
	return &Node{1, key, value, nil, nil}
}

// Set a value into the tree
//
// Returns false is value already exists in tree, otherwise true
func (bt *BinaryTree) Set(key, value interface{}) bool {
	bt.Lock()
	defer bt.Unlock()

	return bt.set(key, value)
}

// Get a value by key from tree
//
// Returns value if key exists, otherwise it returns nil, false
func (bt *BinaryTree) Get(key interface{}) (interface{}, bool) {
	bt.Lock()
	defer bt.Unlock()

	return bt.get(key)
}

// Remove a value from the tree
//
// Returns true if value was removed, otherwise false.
func (bt *BinaryTree) Remove(value interface{}) bool {
	bt.Lock()
	defer bt.Unlock()

	return bt.remove(value)
}

// Height returns the height of the tree (node/level count) in O(1) Time Complexity.
func (bt *BinaryTree) Height() int {
	bt.RLock()
	defer bt.RUnlock()

	return bt.height()
}

// Min returns the minimum value present in the tree
//
// Returns false if tree is empty
func (bt *BinaryTree) Min() (tree.Node, bool) {
	bt.RLock()
	defer bt.RUnlock()

	return bt.min()
}

// Max returns the maximum value present in the tree
//
// Returns false if tree is empty
func (bt *BinaryTree) Max() (tree.Node, bool) {
	bt.RLock()
	defer bt.RUnlock()

	return bt.max()
}

// Contains return true if value exists in tree, otherwise false
func (bt *BinaryTree) Contains(value interface{}) bool {
	bt.RLock()
	defer bt.RUnlock()

	return bt.contains(value)
}

// Length returns the total number of nodes in tree
func (bt *BinaryTree) Length() int {
	bt.RLock()
	defer bt.RUnlock()

	return bt.length()
}

// Clear all the nodes from tree
func (bt *BinaryTree) Clear() {
	bt.Lock()
	defer bt.Unlock()

	bt.clear()
}

// InOrder returns a list.List with all values in-order
func (bt *BinaryTree) InOrder() list.List {
	bt.RLock()
	defer bt.RUnlock()

	return bt.inOrder()
}

// PreOrder returns a list.List with all values in pre order
func (bt *BinaryTree) PreOrder() list.List {
	bt.RLock()
	defer bt.RUnlock()

	return bt.preOrder()
}

// PostOrder returns a list.List with all values in post order
func (bt *BinaryTree) PostOrder() list.List {
	bt.RLock()
	defer bt.RUnlock()

	return bt.postOrder()
}
