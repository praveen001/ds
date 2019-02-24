package ds

// Iterator allows to iterator through the data structures
type Iterator interface {
	// Previous moves the iterator to previous element
	Previous() bool

	// Next moves the iterator to next element
	Next() bool

	// Index returns current index
	Index() int

	// Value returns value at current index
	Value() interface{}
}
