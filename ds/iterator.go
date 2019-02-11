package ds

// Iterator allows to iterator through the data structures
type Iterator interface {
	// Previous moves the iterator to previous element
	Previous()

	// Next moves the iterator to next element
	Next()

	// HasNext returns true if the iteration has more elements
	HasNext() bool

	// HasPrevious returns true if the iteration has a previous element
	HasPrevious() bool

	// Index returns current index
	Index() int

	// Value returns value at current index
	Value() interface{}

	// First moves iterator to first element
	First()

	// Last moves iterator to last element
	Last()

	// MoveTo moves iterator to the given index i, returns false if index i is out of bound
	MoveTo(i int) bool
}
