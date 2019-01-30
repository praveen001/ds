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
    -   [Binary Tree](#binary-tree)
    -   [AVL Tree](#avl-tree)
-   [Heap](#heap-interface)
    -   [Binary Heap](#binary-heap)
    -   [D-ary Heap](#d-ary-heap)
-   [Hash]()
    -   [HashMap](#hash-map)

# List Interface

```go
type List interface {
	// Len returns the number of elements in list
	Len() int

	// Front returns the first element of list or nil if the list is empty
	Front() (interface{}, bool)

	// Back returns the last element of the list or nil if the list is empty
	Back() (interface{}, bool)

	// PushFront inserts a new element with value v at the front of the list
	PushFront(v interface{})

	// PushBack inserts a new element with value v at the back of the list
	PushBack(v interface{})

	// Set a new element with value v at the given index i.
	// if index i is out of bound, it returns false, otherwise true
	Set(i int, v interface{}) (ok bool)

	// Get ..
	Get(i int) (v interface{}, ok bool)

	// Remove the element at given index i. Returns true if element was removed otherwise false.
	Remove(i int) (v interface{}, ok bool)

	// Clear the list
	Clear()

	// PushBackList inserts a copy of an other list at the back of list l.
	// The lists l and other may be the same. They must not be nil.
	PushBackList(l List)

	// PushFrontList inserts a copy of an other list at the front of list l.
	// The lists l and other may be the same. They must not be nil.
	PushFrontList(l List)

	// Contains returns true if the given value exists in the list, otherwise false
	Contains(value interface{}) bool

	// IndexOf returns the index of the given value if it exists, otherwise it returns -1
	IndexOf(value interface{}) int

	// Values returns all the values in the list as a slice
	Values() []interface{}

	// Clone creates a shallow copy and returns the reference
	Clone() List

	// Swap two values at two given indexes
	Swap(a, b int) bool

	// String ..
	String() string
}

```

## Array List

-   Array list implements [Tree](#list-interface) interface.
-   It uses Go slices to store the values.

**Usage**

## Singly Linked List

-   Singly linked list implements [Tree](#list-interface) interface.

**Usage**

## Doubly Linked List

-   Doubly linked list implements [Tree](#list-interface) interface.

**Usage**

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
	// Set a value into the tree
	//
	// Returns false if key already exists in tree, otherwise true
	Set(key, value interface{}) bool

	// Get a value by key
	//
	// Returns value if key exists, otherwise it returns nil, false
	Get(key interface{}) (interface{}, bool)

	// Remove a value from the tree
	//
	// Returns true if value was removed, otherwise false.
	Remove(value interface{}) bool

	// Height returns the height of the tree (node/level count) in O(1) Time Complexity.
	Height() int

	// Min returns the minimum value present in the tree
	//
	// Returns false if tree is empty
	Min() (Node, bool)

	// Max returns the maximum value present in the tree
	//
	// Returns false if tree is empty
	Max() (Node, bool)

	// Contains return true if key exists in tree, otherwise false
	Contains(key interface{}) bool

	// Len returns the total number of nodes in tree
	Len() int

	// Clear all the nodes from tree
	Clear()

	// InOrder returns a ds.List with all values in-order
	InOrder() List

	// PreOrder returns a ds.List with all values in pre order
	PreOrder() List

	// PostOrder returns a ds.List with all values in post order
	PostOrder() List
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

	// Len gives the number of values in heap
	Len() int

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

	// Len returns total number of entries in map
	Len() int

	// Clear removes all the entries from map
	Clear()
}
```

## Hash Map

**Usage**
