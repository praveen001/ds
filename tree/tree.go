package tree

import "github.com/praveen001/ds/list"

// Tree ..
type Tree interface {
	// Insert a given value into the tree
	Insert(value interface{})

	// Delete a node (using value) from the tree
	Delete(value interface{}) bool

	// Contains return true if value exists in tree, otherwise false
	Contains(value interface{}) bool

	// Height returns the height of the tree (node/level count)
	Height() int

	// Min returns the minimum value present in the tree
	Min() (interface{}, bool)

	// Max returns the maximum value present in the tree
	Max() (interface{}, bool)

	// Length returns the total number of nodes in tree
	Length() int

	// Clear removes all the nodes from tree
	Clear()

	// InOrder returns a list.List with all values in-order
	InOrder() list.List
}
