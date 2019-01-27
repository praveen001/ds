package doublylinkedlist

import (
	"sync"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/utils"
)

// DoublyLinkedList holds the set of elements in a slice
//
// It also stores the total number of elements currently present in the list, so length of the list is given in O(1)
//
// It uses sync.RWMutex to support concurrent access, all read operations acquires a RLock, and all write operatiors acquires a Lock()
type DoublyLinkedList struct {
	head *element
	tail *element
	size int
	sync.RWMutex
}

type element struct {
	value interface{}
	next  *element
	prev  *element
}

// New creates a new empty list and return it's reference.
func New() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Append a set of new values at the end of the list
func (dl *DoublyLinkedList) Append(values ...interface{}) {
	dl.Lock()
	defer dl.Unlock()

	dl.append(values...)
}

// Prepend a set of new values at the beginning of the list
func (dl *DoublyLinkedList) Prepend(values ...interface{}) {
	dl.Lock()
	defer dl.Unlock()

	dl.prepend(values...)
}

// Get the value at an index.
//
// Returns the value present at the index, if index is out of bound it will return nil.
//
// Second argument will be false if index is out of bound, otherwise it will be true.
func (dl *DoublyLinkedList) Get(index int) (interface{}, bool) {
	dl.RLock()
	defer dl.RUnlock()

	return dl.get(index)
}

// Set a value at an index
//
// Returns false if index is out of bound, otherwise it will be true.
func (dl *DoublyLinkedList) Set(index int, value interface{}) bool {
	dl.Lock()
	defer dl.Unlock()

	return dl.set(index, value)
}

// Remove the value at an index
//
// Returns the removed value present, if index is out of bound it will return nil.
//
// Second argument will be false if index is out of bound, otherwise it will be true.
func (dl *DoublyLinkedList) Remove(index int) (interface{}, bool) {
	dl.Lock()
	defer dl.Unlock()

	return dl.remove(index)
}

// Contains returns true if the given value exists in the list, otherwise false
func (dl *DoublyLinkedList) Contains(value interface{}) bool {
	dl.RLock()
	defer dl.RUnlock()

	return dl.contains(value)
}

// IndexOf returns the index of the given value if it exists, otherwise it returns -1
func (dl *DoublyLinkedList) IndexOf(value interface{}) int {
	dl.RLock()
	defer dl.RUnlock()

	return dl.indexOf(value)
}

// Values returns all the values in the list as a slice
func (dl *DoublyLinkedList) Values() []interface{} {
	dl.RLock()
	defer dl.RUnlock()

	return dl.values()
}

// Length returns the total number of elements in the list
func (dl *DoublyLinkedList) Length() int {
	dl.RLock()
	defer dl.RUnlock()

	return dl.length()
}

// Clear the list
func (dl *DoublyLinkedList) Clear() {
	dl.Lock()
	defer dl.Unlock()

	dl.clear()
}

// WithInRange returns true if the given index is valid, otherwise false
func (dl *DoublyLinkedList) WithInRange(index int) bool {
	dl.RLock()
	defer dl.RUnlock()

	return dl.withInRange(index)
}

// String returns the string representation of the list
func (dl *DoublyLinkedList) String() string {
	dl.RLock()
	defer dl.RUnlock()

	return dl.string()
}

// Filter creates a new list with every value that passes uitls.FilterFunc
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (dl *DoublyLinkedList) Filter(fn utils.FilterFunc) list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.filter(fn)
}

// Concat joins two or more lists together
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (dl *DoublyLinkedList) Concat(lists ...list.List) list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.concat(lists...)
}

// Reverse the order of items in the list
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (dl *DoublyLinkedList) Reverse() list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.reverse()
}

// Sort arrange the values in ascending or descending order
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (dl *DoublyLinkedList) Sort(utils.CompareFunc) {

}

// Map creates a new list with values returned by the MapFunc
//
// Each value in the list is passed to the MapFunc, and values returned by MapFunc are used to construct a new list.
//
// It doesn't mutate the list, instead it creates a new list and returns it's reference.
func (dl *DoublyLinkedList) Map(fn utils.MapFunc) list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.mp(fn)
}

// Clone creates a shallow copy and returns the reference
func (dl *DoublyLinkedList) Clone() list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.clone()
}

// Swap two values at two given indexes
//
// Returns false if swap doesn't succeed, otherwise true
func (dl *DoublyLinkedList) Swap(a, b int) bool {
	dl.Lock()
	defer dl.Unlock()

	return dl.swap(a, b)
}
