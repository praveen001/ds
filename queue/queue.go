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
	return q.list.Get(q.list.Count() - 1)
}

// Count returns the total number of values in queue
func (q *Queue) Count() int {
	return q.list.Count()
}

// Empty clears the queue
func (q *Queue) Empty() {
	q.list.Empty()
}

// Values returns all the values in the queue as a slice
func (q *Queue) Values() []interface{} {
	return q.list.Values()
}

// String returns the string representation of the stack
func (q *Queue) String() string {
	return q.list.String()
}
