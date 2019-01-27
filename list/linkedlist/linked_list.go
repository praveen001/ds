package linkedlist

import (
	"sync"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/utils"
)

// LinkedList holds the set of elements in a slice
//
// It also stores the total number of elements currently present in the list, so length of the list is given in O(1)
//
// It uses sync.RWMutex to support concurrent access, all read operations acquires a RLock, and all write operatiors acquires a Lock()
type LinkedList struct {
	head *element
	tail *element
	size int
	sync.RWMutex
}

type element struct {
	value interface{}
	next  *element
}

// New creates a new empty list and return it's reference.
func New() *LinkedList {
	return &LinkedList{}
}

// Append a set of new values at the end of the list
func (ll *LinkedList) Append(values ...interface{}) {
	ll.Lock()
	defer ll.Unlock()

	ll.append(values...)
}

// Prepend a set of new values at the beginning of the list
func (ll *LinkedList) Prepend(values ...interface{}) {
	ll.Lock()
	defer ll.Unlock()

	ll.prepend(values...)
}

// Get the value at an index.
//
// Returns the value present at the index, if index is out of bound it will return nil.
//
// Second argument will be false if index is out of bound, otherwise it will be true.
func (ll *LinkedList) Get(index int) (interface{}, bool) {
	ll.RLock()
	defer ll.RUnlock()

	return ll.get(index)
}

// Set a value at an index
//
// Returns false if index is out of bound, otherwise it will be true.
func (ll *LinkedList) Set(index int, value interface{}) bool {
	ll.Lock()
	defer ll.Unlock()

	return ll.set(index, value)
}

// Remove the value at an index
//
// Returns the removed value present, if index is out of bound it will return nil.
//
// Second argument will be false if index is out of bound, otherwise it will be true
func (ll *LinkedList) Remove(index int) (interface{}, bool) {
	ll.Lock()
	defer ll.Unlock()

	return ll.remove(index)
}

// Contains returns true if the given value exists in the list, otherwise false
func (ll *LinkedList) Contains(value interface{}) bool {
	ll.RLock()
	defer ll.RUnlock()

	return ll.contains(value)
}

// IndexOf returns the index of the given value if it exists, otherwise it returns -1
func (ll *LinkedList) IndexOf(value interface{}) int {
	ll.RLock()
	defer ll.RUnlock()

	return ll.indexOf(value)
}

// Values returns all the values in the list as a slice
func (ll *LinkedList) Values() []interface{} {
	ll.RLock()
	defer ll.RUnlock()

	return ll.values()
}

// Length returns the total number of elements in the list
func (ll *LinkedList) Length() int {
	ll.RLock()
	defer ll.RUnlock()

	return ll.length()
}

// Clear the list
func (ll *LinkedList) Clear() {
	ll.Lock()
	defer ll.Unlock()

	ll.clear()
}

// WithInRange returns true if the given index is valid, otherwise false
func (ll *LinkedList) WithInRange(index int) bool {
	ll.RLock()
	defer ll.RUnlock()

	return ll.withInRange(index)
}

// String returns the string representation of the list
func (ll *LinkedList) String() string {
	ll.RLock()
	defer ll.RUnlock()

	return ll.string()
}

// Filter creates a new list with every value that passes uitls.FilterFunc
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (ll *LinkedList) Filter(fn utils.FilterFunc) list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.filter(fn)
}

// Concat joins two or more lists together
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (ll *LinkedList) Concat(lists ...list.List) list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.concat(lists...)
}

// Reverse the order of items in the list
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (ll *LinkedList) Reverse() list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.reverse()
}

// Sort arrange the values in ascending or descending order
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (ll *LinkedList) Sort(utils.CompareFunc) {

}

// Map creates a new list with values returned by the MapFunc
//
// Each value in the list is passed to the MapFunc, and values returned by MapFunc are used to construct a new list.
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (ll *LinkedList) Map(fn utils.MapFunc) list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.mp(fn)
}

// Clone creates a shallow copy and returns the reference
func (ll *LinkedList) Clone() list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.clone()
}

// Swap two values at two given indexes
//
// Returns false if swap doesn't succeed, otherwise true
func (ll *LinkedList) Swap(a, b int) bool {
	ll.Lock()
	defer ll.Unlock()

	return ll.swap(a, b)
}
