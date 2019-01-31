package ds

// Queue interface
type Queue interface {
	// Len returns the total number of values in queue
	Len() (l int)

	// Enqueue adds the given value at the end of the queue
	Enqueue(v interface{})

	// Dequeue removes and returns the first value from the queue
	Dequeue() (v interface{}, ok bool)

	// Peek returns the first value without remove it from the queue
	Peek() (v interface{}, ok bool)

	// Clear the queue
	Clear()

	// Values returns all the values in the queue as a slice
	Values() (vs []interface{})

	// String returns the string representation of the stack
	String() (s string)
}
