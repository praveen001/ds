# Data structures
- List 
	- Array List
	- Singly Linked List
- Queue
- Stack
- Tree
	- Binary Tree
	- AVL Tree
- Heap
	- Binary Heap
- Hash
	- HashMap
		
* All data structures are safe for concurrent access. Mutex based thread-safety

# Interfaces
```go
type List interface {
	// Append new values to the ending of the list
	Append(values ...interface{})

	// Prepend adds new values to the beginning of the list
	Prepend(values ...interface{})

	// Get returns the value at the given index
	Get(index int) (interface{}, bool)

	// Set assigns a value at the given index
	Set(index int, value interface{}) bool

	// Remove removes the value at the given index
	Remove(index int) (interface{}, bool)

	// Contains returns true if the given value exists in the list, otherwise false
	Contains(value interface{}) bool

	// IndexOf returns the index of the given value if it exists, otherwise it returns -1
	IndexOf(value interface{}) int

	// Values returns all the values in the list as an array
	Values() []interface{}

	// Length returns the total number of elements in the list
	Length() int

	// Clear the list
	Clear()

	// WithInRange returns true if the given index is valid, otherwise false
	WithInRange(index int) bool

	// String returns the string representation of the list
	String() string

	// Concat joins two or more lists together
	Concat(...List) List

	// Reverse reverses the order of items in the list
	Reverse() List

	// Sort arrange the values in ascending or descending order
	Sort(utils.CompareFunc)

	// Filter creates a new list with every value that pass a test
	Filter(utils.FilterFunc) List

	// Map creates a new list with every value returned by the MapFunc
	Map(utils.MapFunc) List

	// Clone creates a shallow copy and returns it
	Clone() List

	Swap(a, b int) bool
}

```

## Array List
  - Array list implements `Tree` interface.
 
**Usage**

## Singly Linked List
  - Array list implements `Tree` interface.
  
**Usage**

## Queue
  - Queue internally uses `linkedlist.List` to store data.

**Usage**

## Stack
  - Stack internally uses `linkedlist.List` to store data.
  
**Usage**

## Binary Tree
  - Insertions and deletions are iterative. 
  - Node in tree stores - value, left child, right child, and height.

**Usage**

## AVL Tree
  - Insertions, deletions, balancing are all iterative.
  - Node in tree stores - value, left child, right child, height and balance factor.

**Usage**

## Binary Heap
  - Binary heap internally uses a `arraylist.ArrayList` to store data.
  
**Usage**

## Hash Map
**Usage**

