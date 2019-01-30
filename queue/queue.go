package queue

import (
	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/list/linkedlist"
)

// Queue ..
type Queue struct {
	list ds.List
}

// Config for queue
type Config struct {
	List ds.List
}

// New creates a new queue and returns it
func New() *Queue {
	return &Queue{
		list: linkedlist.New(),
	}
}

// NewWithConfig creates a new queue with given config
func NewWithConfig(c *Config) *Queue {
	return &Queue{
		list: c.List,
	}
}

// Enqueue adds the given value at the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	q.list.PushBack(value)
}

// Dequeue removes and returns the first value from the queue
func (q *Queue) Dequeue() (interface{}, bool) {
	return q.list.Remove(0)
}

// Peek returns the first value without remove it from the queue
func (q *Queue) Peek() (interface{}, bool) {
	return q.list.At(0)
}

// Len returns the total number of values in queue
func (q *Queue) Len() int {
	return q.list.Len()
}

// Clear the queue
func (q *Queue) Clear() {
	q.list.Clear()
}

// Values returns all the values in the queue as a slice
func (q *Queue) Values() []interface{} {
	return q.list.Values()
}

// String returns the string representation of the stack
func (q *Queue) String() string {
	return q.list.String()
}
