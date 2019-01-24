package list

// List ..
type List interface {
	// Append new values to the ending of the list
	Append(values ...interface{})

	// Prepend adds new values to the beginning of the list
	Prepend(values ...interface{})

	// Get returns the value at the given index, if index is not in range, it returns IndexOutOfBound error
	Get(index int) (interface{}, error)

	// Set assigns a value at the given index, if index is not in range, it returns IndexOutOfBound error
	Set(index int, value interface{}) error

	// Remove removes the value at the given index, if index is not in range, it returns IndexOutOfBound error
	Remove(index int) (interface{}, error)

	// Contains returns true if the given value exists in the list, otherwise false
	Contains(value int) bool

	// IndexOf returns the index of the given value if it exists, otherwise it returns -1
	IndexOf(value int) int

	// Size returns the total number of elements in the list
	Size() int

	// Empty clears the list
	Empty()

	// WithInRange returns true if the given index is valid, otherwise false
	WithInRange(index int) bool

	// String returns the string representation of the list
	String() string
}
