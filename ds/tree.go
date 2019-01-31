package ds

// Tree ..
type Tree interface {
	// Set a value into the tree
	//
	// Returns false if key already exists in tree, otherwise true
	Set(k, v interface{}) (ok bool)

	// Get a value by key
	//
	// Returns value if key exists, otherwise it returns nil, false
	Get(k interface{}) (v interface{}, ok bool)

	// Remove a value from the tree
	//
	// Returns true if value was removed, otherwise false.
	Remove(k interface{}) (ok bool)

	// Height returns the height of the tree (node/level count) in O(1) Time Complexity.
	Height() (h int)

	// Min returns the minimum value present in the tree
	//
	// Returns false if tree is empty
	Min() (n Node, ok bool)

	// Max returns the maximum value present in the tree
	//
	// Returns false if tree is empty
	Max() (n Node, ok bool)

	// Contains return true if key exists in tree, otherwise false
	Contains(key interface{}) (ok bool)

	// Len returns the total number of nodes in tree
	Len() (l int)

	// Clear all the nodes from tree
	Clear()

	// InOrder returns a ds.List with all values in-order
	InOrder() (l List)

	// PreOrder returns a ds.List with all values in pre order
	PreOrder() (l List)

	// PostOrder returns a ds.List with all values in post order
	PostOrder() (l List)
}

// Node ..
type Node interface {
	Key() (k interface{})
	Value() (k interface{})
}
