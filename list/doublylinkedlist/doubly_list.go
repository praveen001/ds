package doublylinkedlist

import (
	"sync"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/utils"
)

// DoublyLinkedList ..
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

// New creates a new linked list and returns its
func New() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Append new values to the ending of the list
func (dl *DoublyLinkedList) Append(values ...interface{}) {
	dl.Lock()
	defer dl.Unlock()

	dl.append(values...)
}

// Prepend adds new values to the beginning of the list
func (dl *DoublyLinkedList) Prepend(values ...interface{}) {
	dl.Lock()
	defer dl.Unlock()

	dl.prepend(values...)
}

// Get returns the value at the given index, if index is not in range, it returns IndexOutOfBound error
func (dl *DoublyLinkedList) Get(index int) (interface{}, bool) {
	dl.RLock()
	defer dl.RUnlock()

	return dl.get(index)
}

// Set assigns a value at the given index, if index is not in range, it returns IndexOutOfBound error
func (dl *DoublyLinkedList) Set(index int, value interface{}) bool {
	dl.Lock()
	defer dl.Unlock()

	return dl.set(index, value)
}

// Remove removes the value at the given index, if index is not in range, it returns IndexOutOfBound error
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

// Values returns all the values in the list as an array
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

// Filter creates a new list with every value that pass a test
func (dl *DoublyLinkedList) Filter(fn utils.FilterFunc) list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.filter(fn)
}

// Concat joins two or more lists together
func (dl *DoublyLinkedList) Concat(lists ...list.List) list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.concat(lists...)
}

// Reverse reverses the order of items in the list
func (dl *DoublyLinkedList) Reverse() list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.reverse()
}

// Sort arrange the values in ascending or descending order
func (dl *DoublyLinkedList) Sort(utils.CompareFunc) {

}

// Map creates a new list with every value returned by the MapFunc
func (dl *DoublyLinkedList) Map(fn utils.MapFunc) list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.mp(fn)
}

// Clone creates a shallow copy and returns it
func (dl *DoublyLinkedList) Clone() list.List {
	dl.RLock()
	defer dl.RUnlock()

	return dl.clone()
}

// Swap two values based on index
func (dl *DoublyLinkedList) Swap(a, b int) bool {
	dl.Lock()
	defer dl.Unlock()

	return dl.swap(a, b)
}
