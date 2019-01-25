package arraylist

import (
	"sync"

	"github.com/praveen001/ds/utils"

	"github.com/praveen001/ds/list"
)

// ArrayList holds the elements in the list
type ArrayList struct {
	elements []interface{}
	size     int
	sync.RWMutex
}

// New creates a new array list and returns it
func New() *ArrayList {
	return &ArrayList{}
}

// Append new values to the ending of the list
func (al *ArrayList) Append(values ...interface{}) {
	al.Lock()
	defer al.Unlock()

	al.append(values...)
}

// Prepend adds new values to the beginning of the list
func (al *ArrayList) Prepend(values ...interface{}) {
	al.Lock()
	defer al.Unlock()

	al.prepend(values...)
}

// Get returns the value at the given index
func (al *ArrayList) Get(index int) (interface{}, bool) {
	al.RLock()
	defer al.RUnlock()

	return al.get(index)
}

// Set assigns a value at the given index
func (al *ArrayList) Set(index int, value interface{}) bool {
	al.Lock()
	defer al.Unlock()

	return al.set(index, value)
}

// Remove removes the value at the given index
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

// Values returns all the values in the list as an array
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

// Filter creates a new list with every value that pass a test
func (al *ArrayList) Filter(fn utils.FilterFunc) list.List {
	al.RLock()
	defer al.RUnlock()

	return al.filter(fn)
}

// Concat joins two or more lists together
func (al *ArrayList) Concat(...list.List) {
	New()
}

// Reverse reverses the order of items in the list
func (al *ArrayList) Reverse() {

}

// Sort arrange the values in ascending or descending order
func (al *ArrayList) Sort(utils.Comparator) {

}

// Map creates a new list with every value returned by the MapFunc
func (al *ArrayList) Map(utils.MapFunc) list.List {
	return New()
}

// Clone creates a shallow copy and returns it
func (al *ArrayList) Clone() list.List {
	return New()
}
