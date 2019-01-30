package stack

import (
	"testing"

	"github.com/praveen001/ds/list/doublylinkedlist"
)

func TestNew(t *testing.T) {
	s := New()

	if s == nil {
		t.Errorf("Expected new stack to be created")
	}
}

func TestNewWithConfig(t *testing.T) {
	s := NewWithConfig(&Config{
		List: doublylinkedlist.New(),
	})

	if s == nil {
		t.Errorf("Expected new stack to be created")
	}
}

func TestPush(t *testing.T) {
	s := New()

	s.Push(10)
	s.Push(20)

	if l := s.Len(); l != 2 {
		t.Errorf("Expected %v Got %v", 2, l)
	}
}

func TestPop(t *testing.T) {
	s := New()

	// Invalid case
	if val, ok := s.Pop(); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}

	s.Push(10)
	s.Push(20)

	// Valid case
	if val, ok := s.Pop(); val != 20 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 20, true, val, ok)
	}

	if val, ok := s.Pop(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if l := s.Len(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}
}

func TestPeek(t *testing.T) {
	s := New()

	// Invalid case
	if val, ok := s.Peek(); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}

	s.Push(10)
	s.Push(20)

	// Valid case
	if val, ok := s.Peek(); val != 20 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 20, true, val, ok)
	}

	if l := s.Len(); l != 2 {
		t.Errorf("Expected %v Got %v", 2, l)
	}
}

func TestLen(t *testing.T) {
	s := New()

	if l := s.Len(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}

	s.Push(10)

	if l := s.Len(); l != 1 {
		t.Errorf("Expected %v Got %v", 1, l)
	}
}

func TestClear(t *testing.T) {
	s := New()

	s.Push(10)

	if l := s.Len(); l != 1 {
		t.Errorf("Expected %v Got %v", 1, l)
	}

	s.Clear()

	if l := s.Len(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}
}

func TestValues(t *testing.T) {
	s := New()

	s.Push(10)
	s.Push(20)

	values := s.Values()

	if values[0] != 10 {
		t.Errorf("Expected %v Got %v", 10, values[0])
	}

	if values[1] != 20 {
		t.Errorf("Expected %v Got %v", 20, values[1])
	}
}

func TestString(t *testing.T) {
	s := New()

	s.Push(10)
	s.Push(20)

	if str := s.String(); str != "[10 20]" {
		t.Errorf("Expected %v Got %v", "[10 20]", str)
	}
}
