package list

import (
	"github.com/praveen001/ds/utils"
)

// List interface is implemented by all list data structures.
type List interface {
	// Append a set of new values at the end of the list
	Append(values ...interface{})

	// Prepend a set of new values at the beginning of the list
	Prepend(values ...interface{})

	// Get the value at an index.
	//
	// Returns the value present at the index, if index is out of bound it will return nil.
	//
	// Second argument will be false if index is out of bound, otherwise it will be true.
	Get(index int) (interface{}, bool)

	// Set a value at an index
	//
	// Returns false if index is out of bound, otherwise it will be true.
	Set(index int, value interface{}) bool

	// Remove the value at an index
	//
	// Returns the removed value present, if index is out of bound it will return nil.
	//
	// Second argument will be false if index is out of bound, otherwise it will be true
	Remove(index int) (interface{}, bool)

	// Contains returns true if the given value exists in the list, otherwise false
	Contains(value interface{}) bool

	// IndexOf returns the index of the given value if it exists, otherwise it returns -1
	IndexOf(value interface{}) int

	// Values returns all the values in the list as a slice
	Values() []interface{}

	// Length returns the total number of elements in the list
	Length() int

	// Clear the list
	Clear()

	// WithInRange returns true if the given index is valid, otherwise false
	WithInRange(index int) bool

	// String returns the string representation of the list
	String() string

	// Concat joins two or more lists together
	//
	// It doesn't mutate the list, instead it creates a new list and returns it's reference.
	Concat(...List) List

	// Reverse the order of items in the list
	//
	// It doesn't mutate the list, instead it creates a new list and returns it's reference.
	Reverse() List

	// Sort arrange the values in ascending or descending order
	//
	// It doesn't mutate the list, instead it creates a new list and returns it's reference.
	Sort(utils.CompareFunc)

	// Filter creates a new list with every value that passes uitls.FilterFunc
	//
	// It doesn't mutate the list, instead it creates a new list and returns it's reference.
	Filter(utils.FilterFunc) List

	// Map creates a new list with values returned by the MapFunc
	//
	// Each value in the list is passed to the MapFunc, and values returned by MapFunc are used to construct a new list.
	//
	// It doesn't mutate the list, instead it creates a new list and returns it's reference.
	Map(utils.MapFunc) List

	// Clone creates a shallow copy and returns the reference
	Clone() List

	// Swap two values at two given indexes
	//
	// Returns false if swap doesn't succeed, otherwise true
	Swap(a, b int) bool
}
