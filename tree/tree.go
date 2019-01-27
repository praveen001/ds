package tree

import "github.com/praveen001/ds/list"

// Tree ..
type Tree interface {
	// Add a value into the tree
	//
	// Returns false is value already exists in tree, otherwise true
	Add(value interface{}) bool

	// Remove a value from the tree
	//
	// Returns true if value was removed, otherwise false.
	Remove(value interface{}) bool

	// Height returns the height of the tree (node/level count) in O(1) Time Complexity.
	Height() int

	// Min returns the minimum value present in the tree
	//
	// Returns false if tree is empty
	Min() (interface{}, bool)

	// Max returns the maximum value present in the tree
	//
	// Returns false if tree is empty
	Max() (interface{}, bool)

	// Contains return true if value exists in tree, otherwise false
	Contains(value interface{}) bool

	// Length returns the total number of nodes in tree
	Length() int

	// Clear all the nodes from tree
	Clear()

	// InOrder returns a list.List with all values in-order
	InOrder() list.List

	// PreOrder returns a list.List with all values in pre order
	PreOrder() list.List

	// PostOrder returns a list.List with all values in post order
	PostOrder() list.List
}
