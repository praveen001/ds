package ds

// MinHeap used to configure heap as Min Heap
const MinHeap = 1

// MaxHeap used to configure heap as Max Heap
const MaxHeap = -1

// Heap interface has to be implemented by all heap variants.
type Heap interface {
	// Push a given values into the tree.
	//
	// Returns true if push succeeds.
	Push(vs ...interface{}) (ok bool)

	// Pop removes first value from tree, and returns it.
	//
	// Second return value will be false if tree is empty.
	Pop() (v interface{}, ok bool)

	// Peek returns the first value in heap without removing it
	//
	// Second return value will be false if tree is empty.
	Peek() (v interface{}, ok bool)

	// Clear removes all the values from heap
	Clear()

	// Len gives the number of values in heap
	Len() (l int)

	// String returns the string representation of the heap
	String() (s string)
}
