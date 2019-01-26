package tree

import "github.com/praveen001/ds/list"

// Tree ..
type Tree interface {
	// Add a given value into the tree
	Add(value interface{}) bool

	// Remove a node (using value) from the tree
	Remove(value interface{}) bool

	// Height returns the height of the tree (node/level count)
	Height() int

	// Min returns the minimum value present in the tree
	Min() (interface{}, bool)

	// Max returns the maximum value present in the tree
	Max() (interface{}, bool)

	// Contains return true if value exists in tree, otherwise false
	Contains(value interface{}) bool

	// Length returns the total number of nodes in tree
	Length() int

	// Clear removes all the nodes from tree
	Clear()

	// InOrder returns a list.List with all values in-order
	InOrder() list.List

	// PreOrder returns a list.List with all values in pre order
	PreOrder() list.List

	// PostOrder returns a list.List with all values in post order
	PostOrder() list.List
}
