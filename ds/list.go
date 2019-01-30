package ds

// List interface is implemented by all list data structures.
type List interface {
	// Len returns the number of elements in list
	Len() int

	// Front returns the first element of list or nil if the list is empty
	Front() interface{}

	// Back returns the last element of the list or nil if the list is empty
	Back() interface{}

	// PushFront inserts a new element with value v at the front of the list
	PushFront(v interface{})

	// PushBack inserts a new element with value v at the back of the list
	PushBack(v interface{})

	// Insert inserts a new element with value v at the given index i.
	// if index i is out of bound, it returns false, otherwise true
	Insert(i int, v interface{}) (ok bool)

	// Remove the element at given index i. Returns true if element was removed otherwise false.
	Remove(i int) (v interface{}, ok bool)

	// Clear the list
	Clear()

	// PushBackList inserts a copy of an other list at the back of list l.
	// The lists l and other may be the same. They must not be nil.
	PushBackList(l List)

	// PushFrontList inserts a copy of an other list at the front of list l.
	// The lists l and other may be the same. They must not be nil.
	PushFrontList(l List)

	// Contains returns true if the given value exists in the list, otherwise false
	Contains(value interface{}) bool

	// IndexOf returns the index of the given value if it exists, otherwise it returns -1
	IndexOf(value interface{}) int

	// Values returns all the values in the list as a slice
	Values() []interface{}

	// Clone creates a shallow copy and returns the reference
	Clone() List

	// Swap two values at two given indexes
	Swap(a, b int) bool

	// At ..
	At(i int) (interface{}, bool)

	// String ..
	String() string
}
