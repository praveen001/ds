package ds

// Iterator allows to iterator through the data structures
type Iterator interface {
	// HasNext returns true if the iteration has more elements
	HasNext() bool

	// HasPrevious returns true if the iteration has a previous element
	HasPrevious() bool

	// Previous moves the iterator to previous element
	Previous() bool

	// Next moves the iterator to next element
	Next() bool

	// Index returns current index
	Index() int

	// Value returns value at current index
	Value() interface{}

	// Set value at current index
	Set(interface{}) bool

	// Add value at current index
	Add(interface{}) bool

	// Remove value at current index
	Remove() bool
}
