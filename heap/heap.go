package heap

const MinHeap = 1
const MaxHeap = -1

// Heap interface
type Heap interface {
	Push(values ...interface{}) bool

	Pop() (interface{}, bool)

	Peek() (interface{}, bool)

	Clear()

	Length() int

	String() string
}
