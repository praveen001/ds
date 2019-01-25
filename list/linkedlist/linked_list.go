package linkedlist

import (
	"sync"
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

// Count returns the total number of elements in the list
func (ll *LinkedList) Count() int {
	ll.RLock()
	defer ll.RUnlock()

	return ll.count()
}

// Empty clears the list
func (ll *LinkedList) Empty() {
	ll.Lock()
	defer ll.RLock()

	ll.empty()
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
