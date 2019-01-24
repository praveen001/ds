package linkedlist

import (
	"errors"
	"fmt"

	"github.com/praveen001/ds/dserrors"
)

// LinkedList ..
type LinkedList struct {
	head *element
	tail *element
	size int
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
	for _, value := range values {
		newElem := &element{value: value}

		if ll.size == 0 {
			ll.head = newElem
		} else {
			ll.tail.next = newElem
		}

		ll.tail = newElem
		ll.size++
	}
}

// Prepend adds new values to the beginning of the list
func (ll *LinkedList) Prepend(values ...interface{}) {
	for _, value := range values {
		h := &element{value: value}
		h.next = ll.head
		ll.head = h
	}
	ll.size += len(values)
}

// Get returns the value at the given index, if index is not in range, it returns IndexOutOfBound error
func (ll *LinkedList) Get(index int) (interface{}, error) {
	if ll.WithInRange(index) {
		elem := ll.getElemByIdx(index)
		return elem.value, nil
	}

	return nil, errors.New(dserrors.IndexOutOfBound)
}

// Set assigns a value at the given index, if index is not in range, it returns IndexOutOfBound error
func (ll *LinkedList) Set(index int, value interface{}) error {
	if ll.WithInRange(index) {
		elem := ll.getElemByIdx(index)
		elem.value = value
		return nil
	}

	return errors.New(dserrors.IndexOutOfBound)
}

// Remove removes the value at the given index, if index is not in range, it returns IndexOutOfBound error
func (ll *LinkedList) Remove(index int) (interface{}, error) {
	if ll.WithInRange(index) {
		var value interface{}
		if index == 0 {
			elem := ll.head
			value = elem.value

			if elem.next != nil {
				ll.head = elem.next
			} else {
				ll.head = nil
			}
		} else {
			elem := ll.getElemByIdx(index - 1)
			value = elem.value

			if elem.next != nil {
				elem.next = elem.next.next
			} else {
				elem.next = nil
			}
		}

		ll.size--

		return value, nil
	}

	return nil, errors.New(dserrors.IndexOutOfBound)
}

// Contains returns true if the given value exists in the list, otherwise false
func (ll *LinkedList) Contains(x int) bool {
	for elem := ll.head; elem != nil; elem = elem.next {
		if elem.value == x {
			return true
		}
	}

	return false
}

// IndexOf returns the index of the given value if it exists, otherwise it returns -1
func (ll *LinkedList) IndexOf(x int) int {
	elem := ll.head
	for i := 0; i < ll.size; i++ {
		if elem.value == x {
			return i
		}
		elem = elem.next
	}

	return -1
}

// Size returns the total number of elements in the list
func (ll *LinkedList) Size() int {
	return ll.size
}

// Empty clears the list
func (ll *LinkedList) Empty() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

// WithInRange returns true if the given index is valid, otherwise false
func (ll *LinkedList) WithInRange(index int) bool {
	return index > -1 && index < ll.size
}

// String returns the string representation of the list
func (ll *LinkedList) String() string {
	str := ""
	for elem := ll.head; elem != nil; elem = elem.next {
		str += fmt.Sprintf("%v ", elem.value)
	}

	return str
}

// getElemByIdx returns the element at the given index
func (ll *LinkedList) getElemByIdx(index int) *element {
	elem := ll.head
	for i := 0; i < index; i++ {
		elem = elem.next
	}

	return elem
}
