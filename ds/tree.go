package ds

// Tree ..
type Tree interface {
	// Set a value into the tree
	//
	// Returns false if key already exists in tree, otherwise true
	Set(key, value interface{}) bool

	// Get a value by key
	//
	// Returns value if key exists, otherwise it returns nil, false
	Get(key interface{}) (interface{}, bool)

	// Remove a value from the tree
	//
	// Returns true if value was removed, otherwise false.
	Remove(value interface{}) bool

	// Height returns the height of the tree (node/level count) in O(1) Time Complexity.
	Height() int

	// Min returns the minimum value present in the tree
	//
	// Returns false if tree is empty
	Min() (Node, bool)

	// Max returns the maximum value present in the tree
	//
	// Returns false if tree is empty
	Max() (Node, bool)

	// Contains return true if key exists in tree, otherwise false
	Contains(key interface{}) bool

	// Length returns the total number of nodes in tree
	Length() int

	// Clear all the nodes from tree
	Clear()

	// InOrder returns a ds.List with all values in-order
	InOrder() List

	// PreOrder returns a ds.List with all values in pre order
	PreOrder() List

	// PostOrder returns a ds.List with all values in post order
	PostOrder() List
}

// Node ..
type Node interface {
	Key() interface{}
	Value() interface{}
}
