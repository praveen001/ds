package daryheap

import (
	"testing"

	"github.com/praveen001/ds/heap"
	"github.com/praveen001/ds/list/linkedlist"
)

func comp(a, b interface{}) int {
	aa := a.(int)
	bb := b.(int)

	if aa > bb {
		return 1
	} else if aa < bb {
		return -1
	}
	return 0
}

func TestNew(t *testing.T) {
	dh := New(comp)

	if dh == nil {
		t.Errorf("Expected new binary heap to be created")
	}
}

func TestNewWithConfig(t *testing.T) {
	dh := NewWithConfig(&Config{
		List:       linkedlist.New(),
		Comparator: comp,
		Variant:    heap.MinHeap,
		D:          3,
	})

	if dh == nil {
		t.Errorf("Expected new binary heap to be created")
	}
}

func TestPush(t *testing.T) {
	dh := New(comp)

	dh.Push()
	dh.Push(30)
	dh.Push(10, 20)

	if l := dh.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}
}

func TestPop(t *testing.T) {
	dh := New(comp)

	dh.Push(30, 10, 20)

	if val, ok := dh.Pop(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if val, ok := dh.Pop(); val != 20 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 20, true, val, ok)
	}
}

func TestPeek(t *testing.T) {
	dh := New(comp)

	dh.Push(10, 20, 30)

	if val, ok := dh.Peek(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if l := dh.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if val, ok := dh.Peek(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}
}

func TestClear(t *testing.T) {
	dh := New(comp)

	dh.Push(10, 20, 30)

	if l := dh.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	dh.Clear()

	if l := dh.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}
}

func TestLength(t *testing.T) {
	dh := New(comp)

	dh.Push(10, 20, 30)

	if l := dh.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}
}

func TestString(t *testing.T) {
	dh := New(comp)

	dh.Push(10, 20, 30)

	if str := dh.String(); str != "[10 20 30]" {
		t.Errorf("Expected %v Got %v", "[10 20 30]", str)
	}
}