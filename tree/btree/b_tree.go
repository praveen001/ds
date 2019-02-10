package btree

import (
	"sync"

	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/utils"
)

// BTree ..
type BTree struct {
	root    *Node
	size    int
	order   int
	compare utils.CompareFunc
	mtx     sync.RWMutex
	sync    bool
}

// Node represents a node in the tree
type Node struct {
	entries  []*entry
	children []*Node
}

type entry struct {
	key   interface{}
	value interface{}
}

// New creates a new instance of b-tree and returns it
func New(c utils.CompareFunc, order int) *BTree {
	return &BTree{
		compare: c,
		order:   order,
		sync:    true,
	}
}

// NewNode returns a new binary tree node with given value
func newNode() *Node {
	return &Node{}
}

// Set a value into the tree
//
// Returns false if value already exists in tree, otherwise true
func (bt *BTree) Set(key, value interface{}) bool {
	bt.lock()
	defer bt.unlock()

	return bt.set(key, value)
}

// Get a value by key from tree
//
// Returns value if key exists, otherwise it returns nil, false
func (bt *BTree) Get(key interface{}) (interface{}, bool) {
	bt.lock()
	defer bt.unlock()

	return bt.get(key)
}

// Remove a key from the tree
//
// Returns true if key was removed, otherwise false.
func (bt *BTree) Remove(key interface{}) bool {
	bt.lock()
	defer bt.unlock()

	return bt.remove(key)
}

// Min returns the minimum value present in the tree
//
// Returns false if tree is empty
func (bt *BTree) Min() (ds.Node, bool) {
	bt.rlock()
	defer bt.runlock()

	return bt.min()
}

// Max returns the maximum value present in the tree
//
// Returns false if tree is empty
func (bt *BTree) Max() (ds.Node, bool) {
	bt.rlock()
	defer bt.runlock()

	return bt.max()
}

// Contains return true if key exists in tree, otherwise false
func (bt *BTree) Contains(key interface{}) bool {
	bt.rlock()
	defer bt.runlock()

	return bt.contains(key)
}

// Len returns the total number of nodes in tree
func (bt *BTree) Len() int {
	bt.rlock()
	defer bt.runlock()

	return bt.len()
}

// Clear all the nodes from tree
func (bt *BTree) Clear() {
	bt.lock()
	defer bt.unlock()

	bt.clear()
}

// Keys returns a ds.List with all keys
func (bt *BTree) Keys() ds.List {
	bt.rlock()
	defer bt.runlock()

	return bt.keys()
}

// Values returns a ds.List with all values
func (bt *BTree) Values() ds.List {
	bt.rlock()
	defer bt.runlock()

	return bt.values()
}

// InOrder returns a ds.List with all values in-order
func (bt *BTree) InOrder() ds.List {
	bt.rlock()
	defer bt.runlock()

	return bt.inOrder()
}

// PreOrder returns a ds.List with all values in pre order
func (bt *BTree) PreOrder() ds.List {
	bt.rlock()
	defer bt.runlock()

	return bt.preOrder()
}

// PostOrder returns a ds.List with all values in post order
func (bt *BTree) PostOrder() ds.List {
	bt.rlock()
	defer bt.runlock()

	return bt.postOrder()
}

func (bt *BTree) lock() {
	if bt.sync {
		bt.mtx.Lock()
	}
}

func (bt *BTree) unlock() {
	if bt.sync {
		bt.mtx.Unlock()
	}
}

func (bt *BTree) rlock() {
	if bt.sync {
		bt.mtx.RLock()
	}
}

func (bt *BTree) runlock() {
	if bt.sync {
		bt.mtx.RUnlock()
	}
}
