package arraylist

import (
	"sync"

	"github.com/praveen001/ds/utils"

	"github.com/praveen001/ds/list"
)

// ArrayList holds the set of elements in a slice
//
// It also stores the total number of elements currently present in the list, so length of the list is given in O(1)
//
// It uses sync.RWMutex to support concurrent access, all read operations acquires a RLock, and all write operatiors acquires a Lock()
type ArrayList struct {
	elements []interface{}
	size     int
	sync.RWMutex
}

// New creates a new empty list and return it's reference.
func New() *ArrayList {
	return &ArrayList{}
}

// Append a set of new values at the end of the list
func (al *ArrayList) Append(values ...interface{}) {
	al.Lock()
	defer al.Unlock()

	al.append(values...)
}

// Prepend a set of new values at the beginning of the list
func (al *ArrayList) Prepend(values ...interface{}) {
	al.Lock()
	defer al.Unlock()

	al.prepend(values...)
}

// Get the value at an index.
//
// Returns the value present at the index, if index is out of bound it will return nil.
//
// Second argument will be false if index is out of bound, otherwise it will be true.
func (al *ArrayList) Get(index int) (interface{}, bool) {
	al.RLock()
	defer al.RUnlock()

	return al.get(index)
}

// Set a value at an index
//
// Returns false if index is out of bound, otherwise it will be true.
func (al *ArrayList) Set(index int, value interface{}) bool {
	al.Lock()
	defer al.Unlock()

	return al.set(index, value)
}

// Remove the value at an index
//
// Returns the removed value present, if index is out of bound it will return nil.
//
// Second argument will be false if index is out of bound, otherwise it will be true.
func (al *ArrayList) Remove(index int) (interface{}, bool) {
	al.Lock()
	defer al.Unlock()

	return al.remove(index)
}

// Contains returns true if the given value exists in the list, otherwise false
func (al *ArrayList) Contains(value interface{}) bool {
	al.RLock()
	defer al.RUnlock()

	return al.contains(value)
}

// IndexOf returns the index of the given value if it exists, otherwise it returns -1
func (al *ArrayList) IndexOf(value interface{}) int {
	al.RLock()
	defer al.RUnlock()

	return al.indexOf(value)
}

// Values returns all the values in the list as a slice
func (al *ArrayList) Values() []interface{} {
	al.RLock()
	defer al.RUnlock()

	return al.values()
}

// Length returns the total number of elements in the list
func (al *ArrayList) Length() int {
	al.RLock()
	defer al.RUnlock()

	return al.length()
}

// Clear the list
func (al *ArrayList) Clear() {
	al.Lock()
	defer al.Unlock()

	al.clear()
}

// WithInRange returns true if the given index is valid, otherwise false
func (al *ArrayList) WithInRange(index int) bool {
	al.RLock()
	defer al.RUnlock()

	return al.withInRange(index)
}

// String returns the string representation of the list
func (al *ArrayList) String() string {
	al.RLock()
	defer al.RUnlock()

	return al.string()
}

// Filter creates a new list with every value that passes uitls.FilterFunc
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (al *ArrayList) Filter(fn utils.FilterFunc) list.List {
	al.RLock()
	defer al.RUnlock()

	return al.filter(fn)
}

// Concat joins two or more lists together
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (al *ArrayList) Concat(lists ...list.List) list.List {
	al.Lock()
	defer al.Unlock()

	return al.concat(lists...)
}

// Reverse the order of items in the list
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (al *ArrayList) Reverse() list.List {
	al.RLock()
	defer al.RUnlock()

	return al.reverse()
}

// Sort arrange the values in ascending or descending order
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (al *ArrayList) Sort(utils.CompareFunc) {

}

// Map creates a new list with values returned by the MapFunc
//
// Each value in the list is passed to the MapFunc, and values returned by MapFunc are used to construct a new list.
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (al *ArrayList) Map(fn utils.MapFunc) list.List {
	al.RLock()
	defer al.RUnlock()

	return al.mp(fn)
}

// Clone creates a shallow copy and returns the reference
func (al *ArrayList) Clone() list.List {
	al.RLock()
	defer al.RUnlock()

	return al.clone()
}

// Swap two values at two given indexes
//
// Returns false if swap doesn't succeed, otherwise true
func (al *ArrayList) Swap(a, b int) bool {
	al.Lock()
	defer al.Unlock()

	return al.swap(a, b)
}
