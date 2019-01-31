package ds

// Stack interface
type Stack interface {
	// Len returns the total values in the stack
	Len() (l int)

	// Push the value to the top of stack
	Push(v interface{})

	// Pop removes the element from the top of the stack, and returns it
	Pop() (v interface{}, ok bool)

	// Peek returns the top element without removing it
	Peek() (v interface{}, ok bool)

	// Clear the stack
	Clear()

	// Values returns all the values in the stack as a slice
	Values() (vs []interface{})

	// String returns the string representation of the stack
	String() string
}
