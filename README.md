[![GoDoc](https://godoc.org/github.com/praveen001/ds?status.svg)](https://godoc.org/github.com/praveen001/ds) [![Build Status](https://travis-ci.org/praveen001/ds.svg?branch=master)](https://travis-ci.org/praveen001/ds)

# Data structures

All data structure implementations are safe for concurrent access (through mutex).

-   [List ](#list-interface)
    -   [Array List](#array-list)
    -   [Singly Linked List](#singly-linked-list)
    -   [Doubly Linked List](#doubly-linked-list)
-   [Queue](#queue)
-   [Stack](#stack)
-   [Tree](#tree-interface)
    -   [Binary Tree](https://github.com/praveen001/ds/blob/master/README.md#binary-tree)
    -   [AVL Tree](#avl-tree)
-   [Heap](#heap-interface)
    -   [Binary Heap](#binary-heap)
    -   [D-ary Heap](#d-ary-heap)
-   [Hash]()
    -   [HashMap](#hash-map)

# List Interface

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

-   Array list implements [Tree](#list-interface) interface.
-   It uses Go slices to store the values.

**Usage**

```go
al := arraylist.New()                       // []
al.Append(0)                                // [0]
al.Append(10, 20, 30)                       // [10, 20, 30]
al.Prepend(-20, -10)                        // [-20, -10, 0, 10, 20, 30]
val, ok := al.Get(1)                        // val = 10, ok = true
ok := al.Set(1, 100)                        // [-20, 100, 0, 10, 20, 30]
val, ok := al.Remove(0)                     // val = 20, ok = true, [100, 0, 10, 20, 30]
ok := al.Contains(100)                      // true
ok := al.Contains(-100)                     // false
ok := al.IndexOf(100)                       // 0
ok := al.IndexOf(200)                       // -1
values := al.Values()                       // [100, 0, 10, 20, 30]
len := al.Length()                          // 5
al.Clear()                                  // []
ok := al.WithInRange(1)                     // false
```

**Advanced Usage**

```go
// al = [100, 0, 10, 20, 30]

// Filter the list
newAl := al.Filter(func(a interface{}) bool {           // [0, 10, 20]
    if a.(int) < 25 {
        return true
    }
    return false
})

// Concat multiple lists together
newAl2 := al.concat(al, newAl)                         // [100, 0, 10, 20, 30, 0, 10, 20]

// Reverse the list
newAl3 := al.Reverse()                                 // [30, 20, 10, 0, 100]

// Map every value using map function
newAl4 := al.Map(func(a interface{}) interface{} {     // [200, 0, 20, 40, 60]
    return a.(int) * 2
})

// Create a clone
newAl5 := al.Clone()

// Swap values
al.Swap(0, 4)                                         // [60, 0, 20, 40, 200]
```

## Singly Linked List

-   Singly linked list implements [Tree](#list-interface) interface.

**Usage**

```go
al := linkedlist.New()                       // []
al.Append(0)                                // [0]
al.Append(10, 20, 30)                       // [10, 20, 30]
al.Prepend(-20, -10)                        // [-20, -10, 0, 10, 20, 30]
val, ok := al.Get(1)                        // val = 10, ok = true
ok := al.Set(1, 100)                        // [-20, 100, 0, 10, 20, 30]
val, ok := al.Remove(0)                     // val = 20, ok = true, [100, 0, 10, 20, 30]
ok := al.Contains(100)                      // true
ok := al.Contains(-100)                     // false
ok := al.IndexOf(100)                       // 0
ok := al.IndexOf(200)                       // -1
values := al.Values()                       // [100, 0, 10, 20, 30]
len := al.Length()                          // 5
al.Clear()                                  // []
ok := al.WithInRange(1)                     // false
```

**Advanced Usage**

```go
// al = [100, 0, 10, 20, 30]

// Filter the list
newAl := al.Filter(func(a interface{}) bool {           // [0, 10, 20]
    if a.(int) < 25 {
        return true
    }
    return false
})

// Concat multiple lists together
newAl2 := al.concat(al, newAl)                         // [100, 0, 10, 20, 30, 0, 10, 20]

// Reverse the list
newAl3 := al.Reverse()                                 // [30, 20, 10, 0, 100]

// Map every value using map function
newAl4 := al.Map(func(a interface{}) interface{} {     // [200, 0, 20, 40, 60]
    return a.(int) * 2
})

// Create a clone
newAl5 := al.Clone()

// Swap values
al.Swap(0, 4)                                         // [60, 0, 20, 40, 200]
```

## Doubly Linked List

-   Doubly linked list implements [Tree](#list-interface) interface.

**Usage**

```go
al := doublylinkedlist.New()                       // []
al.Append(0)                                // [0]
al.Append(10, 20, 30)                       // [10, 20, 30]
al.Prepend(-20, -10)                        // [-20, -10, 0, 10, 20, 30]
val, ok := al.Get(1)                        // val = 10, ok = true
ok := al.Set(1, 100)                        // [-20, 100, 0, 10, 20, 30]
val, ok := al.Remove(0)                     // val = 20, ok = true, [100, 0, 10, 20, 30]
ok := al.Contains(100)                      // true
ok := al.Contains(-100)                     // false
ok := al.IndexOf(100)                       // 0
ok := al.IndexOf(200)                       // -1
values := al.Values()                       // [100, 0, 10, 20, 30]
len := al.Length()                          // 5
al.Clear()                                  // []
ok := al.WithInRange(1)                     // false
```

**Advanced Usage**

```go
// al = [100, 0, 10, 20, 30]

// Filter the list
newAl := al.Filter(func(a interface{}) bool {           // [0, 10, 20]
    if a.(int) < 25 {
        return true
    }
    return false
})

// Concat multiple lists together
newAl2 := al.concat(al, newAl)                         // [100, 0, 10, 20, 30, 0, 10, 20]

// Reverse the list
newAl3 := al.Reverse()                                 // [30, 20, 10, 0, 100]

// Map every value using map function
newAl4 := al.Map(func(a interface{}) interface{} {     // [200, 0, 20, 40, 60]
    return a.(int) * 2
})

// Create a clone
newAl5 := al.Clone()

// Swap values
al.Swap(0, 4)                                         // [60, 0, 20, 40, 200]
```

## Queue

-   Queue internally uses `linkedlist.List` to store data.
-   You can configure queue to use any type that implement [List](#list-interface)

**Usage**

```go
q := queue.New()                        // []
q.Enqueue(10)                           // [10]
q.Enqueue(20)                           // [10, 20]
val, ok := q.Dequeue()                  // val = 10, ok = true, [20]
val, ok := q.Peek()                     // val = 20, ok = true, [20]
len := q.Length()                       // len = 1
q.Values()                              // [20]
q.Clear()                               // []
```

**Advanced Usage**

```go
// Make queue use other type of lists
q := queue.NewWithConfig(&queue.Config{
    List: doublylinkedlist.New(),
})
```

## Stack

-   Stack internally uses `linkedlist.List` to store data.
-   You can configure stack to use any type that implement [List](#list-interface)

**Usage**

```go
s := queue.New()                        // []
s.Push(10)                              // [10]
s.Push(20)                              // [10, 20]
val, ok := s.Pop()                      // val = 10, ok = true, [20]
val, ok := s.Peek()                     // val = 20, ok = true, [20]
len := s.Length()                       // len = 1
s.Values()                              // [20]
s.Clear()                               // []
```

**Advanced Usage**

```go
// Make stack use other type of lists
s := stack.NewWithConfig(&stack.Config{
    List: doublylinkedlist.New(),
})
```

# Tree Interface

```go
type Tree interface {
	// Add a value into the tree
	//
	// Returns false is value already exists in tree, otherwise true
	Add(value interface{}) bool

	// Remove a value from the tree
	//
	// Returns true if value was removed, otherwise false.
	Remove(value interface{}) bool

	// Height returns the height of the tree (node/level count) in O(1) Time Complexity.
	Height() int

	// Min returns the minimum value present in the tree
	//
	// Returns false if tree is empty
	Min() (interface{}, bool)

	// Max returns the maximum value present in the tree
	//
	// Returns false if tree is empty
	Max() (interface{}, bool)

	// Contains return true if value exists in tree, otherwise false
	Contains(value interface{}) bool

	// Length returns the total number of nodes in tree
	Length() int

	// Clear all the nodes from tree
	Clear()

	// InOrder returns a list.List with all values in-order
	InOrder() list.List

	// PreOrder returns a list.List with all values in pre order
	PreOrder() list.List

	// PostOrder returns a list.List with all values in post order
	PostOrder() list.List
}

```

## Binary Tree

-   Insertions and deletions are iterative.
-   Node in tree stores - value, left child, right child, and height.

**Usage**

## AVL Tree

-   Insertions, deletions, balancing are all iterative.
-   Node in tree stores - value, left child, right child, height and balance factor.

**Usage**

# Heap Interface

```go
type Heap interface {
	// Push a given values into the tree.
	//
	// Returns true if push succeeds.
	Push(values ...interface{}) bool

	// Pop removes first value from tree, and returns it.
	//
	// Second return value will be false if tree is empty.
	Pop() (interface{}, bool)

	// Peek returns the first value in heap without removing it
	//
	// Second return value will be false if tree is empty.
	Peek() (interface{}, bool)

	// Clear removes all the values from heap
	Clear()

	// Length gives the number of values in heap
	Length() int

	// String returns the string representation of the heap
	String() string
}

```

## Binary Heap

-   Binary heap by default uses `linkedlist.LinkedList` to store data.

**Usage**

## D-ary Heap

-   D-ary heap by default uses `linkedlist.LinkedList` to store data.

**Usage**

# Map Interface

```go
// Map ..
type Map interface {
	// Set a value in map
	Set(key, val interface{})

	// Get finds value by key and returns it, if found, otherwise it returns nil
	Get(key interface{}) (interface{}, bool)

	// Remove a value from map
	Remove(key interface{}) bool

	// Keys returns list of keys present in map
	Keys() []interface{}

	// Values returns list of values present in map
	Values() []interface{}

	// Contains return whether given key exists in map
	Contains(key interface{}) bool

	// Length returns total number of entries in map
	Length() int

	// Clear removes all the entries from map
	Clear()
}
```

## Hash Map

**Usage**
