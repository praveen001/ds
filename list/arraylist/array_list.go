package arraylist

import (
	"sync"
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

// Count returns the total number of elements in the list
func (al *ArrayList) Count() int {
	al.RLock()
	defer al.RUnlock()

	return al.count()
}

// Empty clears the list
func (al *ArrayList) Empty() {
	al.Lock()
	defer al.Unlock()

	al.empty()
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
