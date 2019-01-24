package queue

import (
	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/linkedlist"
)

// Queue ..
type Queue struct {
	list list.List
}

// New creates a new queue and returns it
func New() *Queue {
	return &Queue{
		list: linkedlist.New(),
	}
}

// Enqueue adds the given value at the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	q.list.Append(value)
}

// Dequeue removes and returns the first value from the queue
func (q *Queue) Dequeue() (interface{}, bool) {
	return q.list.Remove(0)
}

// Peek returns the first value without remove it from the queue
func (q *Queue) Peek() (interface{}, bool) {
	return q.list.Get(q.list.Size() - 1)
}

// Size returns the total number of values in queue
func (q *Queue) Size() int {
	return q.list.Size()
}

// Empty clears the queue
func (q *Queue) Empty() {
	q.list.Empty()
}

// String returns the string representation of the stack
func (q *Queue) String() string {
	return q.list.String()
}
