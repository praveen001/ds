package queue

import (
	"testing"

	"github.com/praveen001/ds/list/doublylinkedlist"
)

func TestNew(t *testing.T) {
	q := New()

	if q == nil {
		t.Errorf("Expected new queue to be created")
	}
}

func TestNewWithConfig(t *testing.T) {
	q := NewWithConfig(&Config{
		List: doublylinkedlist.New(),
	})

	if q == nil {
		t.Errorf("Expected new queue to be created")
	}
}

func TestEnqueue(t *testing.T) {
	q := New()

	q.Enqueue(10)
	q.Enqueue(20)

	if l := q.Length(); l != 2 {
		t.Errorf("Expected %v Got %v", 2, l)
	}
}

func TestDequeue(t *testing.T) {
	q := New()

	// Invalid case
	if val, ok := q.Dequeue(); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}

	q.Enqueue(10)
	q.Enqueue(20)

	// Valid case
	if val, ok := q.Dequeue(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if val, ok := q.Dequeue(); val != 20 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 20, true, val, ok)
	}

	if l := q.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}
}

func TestPeek(t *testing.T) {
	q := New()

	// Invalid case
	if val, ok := q.Peek(); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}

	q.Enqueue(10)
	q.Enqueue(20)

	// Valid case
	if val, ok := q.Peek(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if l := q.Length(); l != 2 {
		t.Errorf("Expected %v Got %v", 2, l)
	}
}

func TestLength(t *testing.T) {
	q := New()

	if l := q.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}

	q.Enqueue(10)

	if l := q.Length(); l != 1 {
		t.Errorf("Expected %v Got %v", 1, l)
	}
}

func TestClear(t *testing.T) {
	q := New()

	q.Enqueue(10)

	if l := q.Length(); l != 1 {
		t.Errorf("Expected %v Got %v", 1, l)
	}

	q.Clear()

	if l := q.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}
}

func TestValues(t *testing.T) {
	q := New()

	q.Enqueue(10)
	q.Enqueue(20)

	values := q.Values()

	if values[0] != 10 {
		t.Errorf("Expected %v Got %v", 10, values[0])
	}

	if values[1] != 20 {
		t.Errorf("Expected %v Got %v", 20, values[1])
	}
}

func TestString(t *testing.T) {
	q := New()

	q.Enqueue(10)
	q.Enqueue(20)

	if str := q.String(); str != "[10 20]" {
		t.Errorf("Expected %v Got %v", "[10 20]", str)
	}
}
