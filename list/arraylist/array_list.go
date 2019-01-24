package arraylist

import "fmt"

// ArrayList holds the elements in the list
type ArrayList struct {
	elements []interface{}
	size     int
}

// New creates a new array list and returns it
func New() *ArrayList {
	return &ArrayList{}
}

// Add new value
func (al *ArrayList) Add(values ...interface{}) *ArrayList {
	al.elements = append(al.elements, values...)
	al.size += len(values)

	return al
}

// Remove removes the element at the given index
func (al *ArrayList) Remove(index int) *ArrayList {
	if al.withInRange(index) {
		al.size--

		newElems := make([]interface{}, al.size)

		copy(newElems, al.elements[:index])
		copy(newElems[index:], al.elements[index+1:])

		al.elements = newElems
	}

	return al
}

// Get return the element at the given index
func (al *ArrayList) Get(index int) (interface{}, bool) {
	if al.withInRange(index) {
		return al.elements[index], true
	}
	return nil, false
}

// Contains check if the given value exists in the list
func (al *ArrayList) Contains(x int) bool {
	if al.IndexOf(x) == -1 {
		return false
	}

	return true
}

// IndexOf finds the given element's index in the list
// If element is not found, it returns -1
func (al *ArrayList) IndexOf(x int) int {
	for i, element := range al.elements {
		if element == x {
			return i
		}
	}

	return -1
}

// Set puts the given value at the given index
func (al *ArrayList) Set(index int, value interface{}) bool {
	if al.withInRange(index) {
		al.elements[index] = value
		return true
	}

	return false
}

// Push inserts the given value at the end of the list
func (al *ArrayList) Push(value interface{}) *ArrayList {
	return al.Add(value)
}

// Pop removes 1 element from the end, and returns it
func (al *ArrayList) Pop() interface{} {
	if al.size < 0 {
		return nil
	}

	end := al.size - 1
	x := al.elements[end]
	al.Remove(end)

	return x
}

// Size returns the number of elements stored in the list
func (al *ArrayList) Size() int {
	return al.size
}

// withInRange returns true if the given index is a valid index
func (al *ArrayList) withInRange(index int) bool {
	return index > -1 && index < al.size
}

// String returns the string representation of the list
func (al *ArrayList) String() string {
	return fmt.Sprintf("%v", al.elements)
}
