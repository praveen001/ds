package binaryheap

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
	bh := New(comp)

	if bh == nil {
		t.Errorf("Expected new binary heap to be created")
	}
}

func TestNewWithConfig(t *testing.T) {
	bh := NewWithConfig(&Config{
		List:       linkedlist.New(),
		Comparator: comp,
		Variant:    heap.MinHeap,
	})

	if bh == nil {
		t.Errorf("Expected new binary heap to be created")
	}
}

func TestPush(t *testing.T) {
	dh := New(comp)

	dh.Push()
	dh.Push(30, 10, 20, 5, 2, 1)

	if l := dh.Len(); l != 6 {
		t.Errorf("Expected %v Got %v", 6, l)
	}

	if val, ok := dh.Peek(); val != 1 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 1, true, val, ok)
	}
}

func TestPop(t *testing.T) {
	dh := New(comp)

	dh.Push(30, 10, 20, 5, 2, 1)

	if val, ok := dh.Pop(); val != 1 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 1, true, val, ok)
	}

	if val, ok := dh.Pop(); val != 2 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 2, true, val, ok)
	}

	if val, ok := dh.Pop(); val != 5 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 5, true, val, ok)
	}

	if val, ok := dh.Pop(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if val, ok := dh.Pop(); val != 20 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 20, true, val, ok)
	}

	if val, ok := dh.Pop(); val != 30 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 30, true, val, ok)
	}

	if val, ok := dh.Pop(); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}

	if str := dh.String(); str != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str)
	}

}

func TestPeek(t *testing.T) {
	bh := New(comp)

	bh.Push(10, 20, 30)

	if val, ok := bh.Peek(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if l := bh.Len(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if val, ok := bh.Peek(); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}
}

func TestClear(t *testing.T) {
	bh := New(comp)

	bh.Push(10, 20, 30)

	if l := bh.Len(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	bh.Clear()

	if l := bh.Len(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}
}

func TestLen(t *testing.T) {
	bh := New(comp)

	bh.Push(10, 20, 30)

	if l := bh.Len(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}
}

func TestString(t *testing.T) {
	bh := New(comp)

	bh.Push(10, 20, 30)

	if str := bh.String(); str != "[10 20 30]" {
		t.Errorf("Expected %v Got %v", "[10 20 30]", str)
	}
}
