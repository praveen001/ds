package arraylist

import (
	"sync"

	"github.com/praveen001/ds/ds"
)

// ArrayList holds the set of elements in a slice
//
// It also stores the total number of elements currently present in the list, so length of the list is given in O(1)
//
// It uses sync.RWMutex to support concurrent access, all read operations acquires a RLock, and all write operatiors acquires a Lock()
type ArrayList struct {
	elements []interface{}
	size     int
	mtx      sync.RWMutex
	sync     bool
}

// New creates a new empty list and return it's reference.
func New() *ArrayList {
	return &ArrayList{sync: true}
}

// Len returns the number of elements in list
func (al *ArrayList) Len() int {
	al.rlock()
	defer al.runlock()

	return al.len()
}

// Front returns the first element of list or nil if the list is empty
func (al *ArrayList) Front() (interface{}, bool) {
	al.rlock()
	defer al.runlock()

	return al.front()
}

// Back returns the last element of the list or nil if the list is empty
func (al *ArrayList) Back() (interface{}, bool) {
	al.rlock()
	defer al.runlock()

	return al.back()
}

// PushFront inserts a new element with value v at the front of the list
func (al *ArrayList) PushFront(v interface{}) {
	al.lock()
	defer al.unlock()

	al.pushFront(v)
}

// PushBack inserts a new element with value v at the front of the list
func (al *ArrayList) PushBack(v interface{}) {
	al.lock()
	defer al.unlock()

	al.pushBack(v)
}

// Set inserts a new element with value v at the given index i.
// if index i is out of bound, it returns false, otherwise true
func (al *ArrayList) Set(i int, v interface{}) (ok bool) {
	al.lock()
	defer al.unlock()

	return al.set(i, v)
}

// Get ..
func (al *ArrayList) Get(i int) (v interface{}, ok bool) {
	al.rlock()
	defer al.runlock()

	return al.get(i)
}

// Remove the element at given index i. Returns true if element was removed otherwise false.
func (al *ArrayList) Remove(i int) (v interface{}, ok bool) {
	al.lock()
	defer al.unlock()

	return al.remove(i)
}

// Clear the list
func (al *ArrayList) Clear() {
	al.lock()
	defer al.unlock()

	al.clear()
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (al *ArrayList) PushBackList(l ds.List) {
	al.lock()
	defer al.unlock()

	al.pushBackList(l)
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (al *ArrayList) PushFrontList(l ds.List) {
	al.lock()
	defer al.unlock()

	al.pushFrontList(l)
}

// Contains returns true if the given value exists in the list, otherwise false
func (al *ArrayList) Contains(v interface{}) bool {
	al.rlock()
	defer al.runlock()

	return al.contains(v)
}

// IndexOf returns the index of the given value v if it exists, otherwise it returns -1
func (al *ArrayList) IndexOf(v interface{}) int {
	al.rlock()
	defer al.runlock()

	return al.indexOf(v)
}

// Values returns all the values in the list as a slice
func (al *ArrayList) Values() []interface{} {
	al.rlock()
	defer al.runlock()

	return al.values()
}

// Clone creates a shallow copy and returns the reference
func (al *ArrayList) Clone() ds.List {
	al.rlock()
	defer al.runlock()

	return al.clone()
}

// Swap two values at two given indexes
func (al *ArrayList) Swap(a, b int) bool {
	al.rlock()
	defer al.runlock()

	return al.swap(a, b)
}

// String returns the string representation of the list
func (al *ArrayList) String() string {
	al.rlock()
	defer al.runlock()

	return al.string()
}

func (al *ArrayList) lock() {
	if al.sync {
		al.mtx.Lock()
	}
}

func (al *ArrayList) unlock() {
	if al.sync {
		al.mtx.Unlock()
	}
}

func (al *ArrayList) rlock() {
	if al.sync {
		al.mtx.RLock()
	}
}

func (al *ArrayList) runlock() {
	if al.sync {
		al.mtx.RUnlock()
	}
}
