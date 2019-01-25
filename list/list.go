package list

// List ..
type List interface {
	// Append new values to the ending of the list
	Append(values ...interface{})

	// Prepend adds new values to the beginning of the list
	Prepend(values ...interface{})

	// Get returns the value at the given index
	Get(index int) (interface{}, bool)

	// Set assigns a value at the given index
	Set(index int, value interface{}) bool

	// Remove removes the value at the given index
	Remove(index int) (interface{}, bool)

	// Contains returns true if the given value exists in the list, otherwise false
	Contains(value interface{}) bool

	// IndexOf returns the index of the given value if it exists, otherwise it returns -1
	IndexOf(value interface{}) int

	// Values returns all the values in the list as an array
	Values() []interface{}

	// Count returns the total number of elements in the list
	Count() int

	// Empty clears the list
	Empty()

	// WithInRange returns true if the given index is valid, otherwise false
	WithInRange(index int) bool

	// String returns the string representation of the list
	String() string
}
