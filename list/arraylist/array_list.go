package arraylist

import (
	"fmt"
)

// ArrayList holds the elements in the list
type ArrayList struct {
	elements []interface{}
	size     int
}

// New creates a new array list and returns it
func New() *ArrayList {
	return &ArrayList{}
}

// Append new values to the ending of the list
func (al *ArrayList) Append(values ...interface{}) {
	al.elements = append(al.elements, values...)
	al.size += len(values)

}

// Prepend adds new values to the beginning of the list
func (al *ArrayList) Prepend(values ...interface{}) {
	length := len(values)

	newSubArr := make([]interface{}, length)
	for i := 0; i < len(values); i++ {
		newSubArr[i] = values[length-1-i]
	}

	al.elements = append(newSubArr, al.elements...)
	al.size += len(values)
}

// Get returns the value at the given index
func (al *ArrayList) Get(index int) (interface{}, bool) {
	if al.WithInRange(index) {
		return al.elements[index], true
	}

	return nil, false
}

// Set assigns a value at the given index
func (al *ArrayList) Set(index int, value interface{}) bool {
	if al.WithInRange(index) {
		al.elements[index] = value
		return true
	}

	return false
}

// Remove removes the value at the given index
func (al *ArrayList) Remove(index int) (interface{}, bool) {
	if al.WithInRange(index) {
		value := al.elements[index]

		newElems := make([]interface{}, al.size)

		copy(newElems, al.elements[:index])
		copy(newElems[index:], al.elements[index+1:])

		al.elements = newElems
		al.size--

		return value, true
	}

	return nil, false
}

// Contains returns true if the given value exists in the list, otherwise false
func (al *ArrayList) Contains(x int) bool {
	if al.IndexOf(x) == -1 {
		return false
	}

	return true
}

// IndexOf returns the index of the given value if it exists, otherwise it returns -1
func (al *ArrayList) IndexOf(x int) int {
	for i, element := range al.elements {
		if element == x {
			return i
		}
	}

	return -1
}

// Size returns the total number of elements in the list
func (al *ArrayList) Size() int {
	return al.size
}

// Empty clears the list
func (al *ArrayList) Empty() {
	al.elements = make([]interface{}, 0)
	al.size = 0
}

// WithInRange returns true if the given index is valid, otherwise false
func (al *ArrayList) WithInRange(index int) bool {
	return index > -1 && index < al.size
}

// String returns the string representation of the list
func (al *ArrayList) String() string {
	return fmt.Sprintf("%v", al.elements)
}
