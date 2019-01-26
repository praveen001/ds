package linkedlist

import (
	"sync"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/utils"
)

// LinkedList ..
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

// New creates a new linked list and returns its
func New() *LinkedList {
	return &LinkedList{}
}

// Append new values to the ending of the list
func (ll *LinkedList) Append(values ...interface{}) {
	ll.Lock()
	defer ll.Unlock()

	ll.append(values...)
}

// Prepend adds new values to the beginning of the list
func (ll *LinkedList) Prepend(values ...interface{}) {
	ll.Lock()
	defer ll.Unlock()

	ll.prepend(values...)
}

// Get returns the value at the given index, if index is not in range, it returns IndexOutOfBound error
func (ll *LinkedList) Get(index int) (interface{}, bool) {
	ll.RLock()
	defer ll.RUnlock()

	return ll.get(index)
}

// Set assigns a value at the given index, if index is not in range, it returns IndexOutOfBound error
func (ll *LinkedList) Set(index int, value interface{}) bool {
	ll.Lock()
	defer ll.Unlock()

	return ll.set(index, value)
}

// Remove removes the value at the given index, if index is not in range, it returns IndexOutOfBound error
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

// Values returns all the values in the list as an array
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

// Filter creates a new list with every value that pass a test
func (ll *LinkedList) Filter(fn utils.FilterFunc) list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.filter(fn)
}

// Concat joins two or more lists together
func (ll *LinkedList) Concat(lists ...list.List) list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.concat(lists...)
}

// Reverse reverses the order of items in the list
func (ll *LinkedList) Reverse() list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.reverse()
}

// Sort arrange the values in ascending or descending order
func (ll *LinkedList) Sort(utils.CompareFunc) {

}

// Map creates a new list with every value returned by the MapFunc
func (ll *LinkedList) Map(fn utils.MapFunc) list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.mp(fn)
}

// Clone creates a shallow copy and returns it
func (ll *LinkedList) Clone() list.List {
	ll.RLock()
	defer ll.RUnlock()

	return ll.clone()
}

// Swap two values based on index
func (ll *LinkedList) Swap(a, b int) bool {
	ll.Lock()
	defer ll.Unlock()

	return ll.swap(a, b)
}
