package linkedlist

import (
	"sync"

	"github.com/praveen001/ds/ds"
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
	mtx  sync.RWMutex
	sync bool
}

type element struct {
	value interface{}
	next  *element
}

// New creates a new empty list and return it's reference.
func New() *LinkedList {
	return &LinkedList{sync: true}
}

// Len returns the number of elements in list
func (ll *LinkedList) Len() int {
	ll.rlock()
	defer ll.runlock()

	return ll.len()
}

// Front returns the first element of list or nil if the list is empty
func (ll *LinkedList) Front() (interface{}, bool) {
	ll.rlock()
	defer ll.runlock()

	return ll.front()
}

// Back returns the last element of the list or nil if the list is empty
func (ll *LinkedList) Back() (interface{}, bool) {
	ll.rlock()
	defer ll.runlock()

	return ll.back()
}

// PushFront inserts a new element with value v at the front of the list
func (ll *LinkedList) PushFront(v interface{}) {
	ll.lock()
	defer ll.unlock()

	ll.pushFront(v)
}

// PushBack inserts a new element with value v at the front of the list
func (ll *LinkedList) PushBack(v interface{}) {
	ll.lock()
	defer ll.unlock()

	ll.pushBack(v)
}

// Set inserts a new element with value v at the given index i.
// if index i is out of bound, it returns false, otherwise true
func (ll *LinkedList) Set(i int, v interface{}) (ok bool) {
	ll.lock()
	defer ll.unlock()

	return ll.set(i, v)
}

// Get ..
func (ll *LinkedList) Get(i int) (v interface{}, ok bool) {
	ll.rlock()
	defer ll.runlock()

	return ll.get(i)
}

// Insert value v at index i
func (ll *LinkedList) Insert(i int, v interface{}) (ok bool) {
	ll.lock()
	defer ll.unlock()

	return ll.insert(i, v)
}

// Remove the element at given index i. Returns true if element was removed otherwise false.
func (ll *LinkedList) Remove(i int) (v interface{}, ok bool) {
	ll.lock()
	defer ll.unlock()

	return ll.remove(i)
}

// Clear the list
func (ll *LinkedList) Clear() {
	ll.lock()
	defer ll.unlock()

	ll.clear()
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (ll *LinkedList) PushBackList(l ds.List) {
	ll.lock()
	defer ll.unlock()

	ll.pushBackList(l)
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (ll *LinkedList) PushFrontList(l ds.List) {
	ll.lock()
	defer ll.unlock()

	ll.pushFrontList(l)
}

// Contains returns true if the given value exists in the list, otherwise false
func (ll *LinkedList) Contains(v interface{}) bool {
	ll.rlock()
	defer ll.runlock()

	return ll.contains(v)
}

// IndexOf returns the index of the given value v if it exists, otherwise it returns -1
func (ll *LinkedList) IndexOf(v interface{}) int {
	ll.rlock()
	defer ll.runlock()

	return ll.indexOf(v)
}

// Values returns all the values in the list as a slice
func (ll *LinkedList) Values() []interface{} {
	ll.rlock()
	defer ll.runlock()

	return ll.values()
}

// Clone creates a shallow copy and returns the reference
func (ll *LinkedList) Clone() ds.List {
	ll.rlock()
	defer ll.runlock()

	return ll.clone()
}

// Swap two values at two given indexes
func (ll *LinkedList) Swap(a, b int) bool {
	ll.rlock()
	defer ll.runlock()

	return ll.swap(a, b)
}

// String returns the string representation of the list
func (ll *LinkedList) String() string {
	ll.rlock()
	defer ll.runlock()

	return ll.string()
}

func (ll *LinkedList) lock() {
	if ll.sync {
		ll.mtx.Lock()
	}
}

func (ll *LinkedList) unlock() {
	if ll.sync {
		ll.mtx.Unlock()
	}
}

func (ll *LinkedList) rlock() {
	if ll.sync {
		ll.mtx.RLock()
	}
}

func (ll *LinkedList) runlock() {
	if ll.sync {
		ll.mtx.RUnlock()
	}
}
