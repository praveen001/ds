package doublylinkedlist

import (
	"sync"

	"github.com/praveen001/ds/ds"
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
	mtx  sync.RWMutex
	sync bool
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

// Len returns the number of elements in list
func (dl *DoublyLinkedList) Len() int {
	dl.rlock()
	defer dl.runlock()

	return dl.len()
}

// Front returns the first element of list or nil if the list is empty
func (dl *DoublyLinkedList) Front() (interface{}, bool) {
	dl.rlock()
	defer dl.runlock()

	return dl.front()
}

// Back returns the last element of the list or nil if the list is empty
func (dl *DoublyLinkedList) Back() (interface{}, bool) {
	dl.rlock()
	defer dl.runlock()

	return dl.back()
}

// PushFront inserts a new element with value v at the front of the list
func (dl *DoublyLinkedList) PushFront(v interface{}) {
	dl.lock()
	defer dl.unlock()

	dl.pushFront(v)
}

// PushBack inserts a new element with value v at the front of the list
func (dl *DoublyLinkedList) PushBack(v interface{}) {
	dl.lock()
	defer dl.unlock()

	dl.pushBack(v)
}

// Set inserts a new element with value v at the given index i.
// if index i is out of bound, it returns false, otherwise true
func (dl *DoublyLinkedList) Set(i int, v interface{}) (ok bool) {
	dl.lock()
	defer dl.unlock()

	return dl.set(i, v)
}

// Get ..
func (dl *DoublyLinkedList) Get(i int) (v interface{}, ok bool) {
	dl.rlock()
	defer dl.unlock()

	return dl.get(i)
}

// Remove the element at given index i. Returns true if element was removed otherwise false.
func (dl *DoublyLinkedList) Remove(i int) (v interface{}, ok bool) {
	dl.lock()
	defer dl.unlock()

	return dl.remove(i)
}

// Clear the list
func (dl *DoublyLinkedList) Clear() {
	dl.lock()
	defer dl.unlock()

	dl.clear()
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (dl *DoublyLinkedList) PushBackList(l ds.List) {
	dl.lock()
	defer dl.unlock()

	dl.pushBackList(l)
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (dl *DoublyLinkedList) PushFrontList(l ds.List) {
	dl.lock()
	defer dl.unlock()

	dl.pushFrontList(l)
}

// Contains returns true if the given value exists in the list, otherwise false
func (dl *DoublyLinkedList) Contains(v interface{}) bool {
	dl.rlock()
	defer dl.runlock()

	return dl.contains(v)
}

// IndexOf returns the index of the given value v if it exists, otherwise it returns -1
func (dl *DoublyLinkedList) IndexOf(v interface{}) int {
	dl.rlock()
	defer dl.runlock()

	return dl.indexOf(v)
}

// Values returns all the values in the list as a slice
func (dl *DoublyLinkedList) Values() []interface{} {
	dl.rlock()
	defer dl.runlock()

	return dl.values()
}

// Clone creates a shallow copy and returns the reference
func (dl *DoublyLinkedList) Clone() ds.List {
	dl.rlock()
	defer dl.runlock()

	return dl.clone()
}

// Swap two values at two given indexes
func (dl *DoublyLinkedList) Swap(a, b int) bool {
	dl.rlock()
	defer dl.runlock()

	return dl.swap(a, b)
}

// String returns the string representation of the list
func (dl *DoublyLinkedList) String() string {
	dl.rlock()
	defer dl.runlock()

	return dl.string()
}

func (dl *DoublyLinkedList) lock() {
	if dl.sync {
		dl.mtx.Lock()
	}
}

func (dl *DoublyLinkedList) unlock() {
	if dl.sync {
		dl.mtx.Unlock()
	}
}

func (dl *DoublyLinkedList) rlock() {
	if dl.sync {
		dl.mtx.RLock()
	}
}

func (dl *DoublyLinkedList) runlock() {
	if dl.sync {
		dl.mtx.RUnlock()
	}
}
