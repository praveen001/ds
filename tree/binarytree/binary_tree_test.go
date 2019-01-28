package binarytree

import (
	"testing"

	"github.com/praveen001/ds/utils"
)

func TestNew(t *testing.T) {
	bt := New(utils.IntComparator)

	if bt == nil {
		t.Errorf("Expected new tree to be created")
	}
}

func TestSet(t *testing.T) {
	bt := New(utils.IntComparator)

	bt.Set(10, 10)

	if l := bt.Length(); l != 1 {
		t.Errorf("Expected %v Got %v", 1, l)
	}

	bt.Set(10, 10)
	bt.Set(20, 20)
	if l := bt.Length(); l != 2 {
		t.Errorf("Expected %v Got %v", 2, l)
	}
}

func TestGet(t *testing.T) {
	bt := New(utils.IntComparator)

	bt.Set(1, 10)
	bt.Set(2, 20)

	if val, ok := bt.Get(1); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if val, ok := bt.Get(3); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}
}

func TestRemove(t *testing.T) {
	bt := New(utils.IntComparator)

	bt.Set(1, 10)
	bt.Set(2, 20)

	if ok := bt.Remove(1); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}

	if ok := bt.Remove(1); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}
}

func TestHeight(t *testing.T) {
	bt := New(utils.IntComparator)

	if h := bt.Height(); h != 0 {
		t.Errorf("Expected %v Got %v", 0, h)
	}

	bt.Set(1, 10)
	bt.Set(2, 20)
	bt.Set(3, 30)

	if h := bt.Height(); h != 3 {
		t.Errorf("Expected %v Got %v", 3, h)
	}
}

func TestMin(t *testing.T) {
	bt := New(utils.IntComparator)

	if m, ok := bt.Min(); m != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, m, ok)
	}

	bt.Set(100, 100)
	bt.Set(10, 10)
	bt.Set(20, 20)
	bt.Set(5, 5)

	if m, ok := bt.Min(); m.value != 5 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 5, true, m.value, ok)
	}
}

func TestMax(t *testing.T) {
	bt := New(utils.IntComparator)

	if m, ok := bt.Max(); m != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, m, ok)
	}

	bt.Set(5, 5)
	bt.Set(10, 10)
	bt.Set(100, 100)
	bt.Set(20, 20)

	if m, ok := bt.Max(); m.value != 100 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 100, true, m.value, ok)
	}
}

func TestContains(t *testing.T) {
	bt := New(utils.IntComparator)

	if ok := bt.Contains(1); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}

	bt.Set(5, 5)
	bt.Set(10, 10)
	bt.Set(100, 100)
	bt.Set(20, 20)

	if ok := bt.Contains(20); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}

	bt.Remove(20)

	if ok := bt.Contains(20); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}
}

func TestLength(t *testing.T) {
	bt := New(utils.IntComparator)

	if l := bt.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}

	bt.Set(5, 5)
	bt.Set(10, 10)
	bt.Set(100, 100)
	bt.Set(20, 20)

	if l := bt.Length(); l != 4 {
		t.Errorf("Expected %v Got %v", 4, l)
	}

	bt.Remove(10)

	if l := bt.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}
}

func TestClear(t *testing.T) {
	bt := New(utils.IntComparator)

	if l := bt.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}

	bt.Set(5, 5)
	bt.Set(10, 10)
	bt.Set(100, 100)
	bt.Set(20, 20)

	if l := bt.Length(); l != 4 {
		t.Errorf("Expected %v Got %v", 4, l)
	}

	bt.Clear()

	if l := bt.Length(); l != 0 {
		t.Errorf("Expected %v Got %v", 0, l)
	}

	if str := bt.InOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}
}

func TestInOrder(t *testing.T) {
	bt := New(utils.IntComparator)

	if str := bt.InOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}

	bt.Set(5, 5)
	bt.Set(10, 10)
	bt.Set(100, 100)
	bt.Set(20, 20)

	if str := bt.InOrder(); str.String() != "[5 10 20 100]" {
		t.Errorf("Expected %v Got %v", "[5 10 20 100]", str.String())
	}
}

func TestPreOrder(t *testing.T) {
	bt := New(utils.IntComparator)

	if str := bt.PreOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}

	bt.Set(5, 5)
	bt.Set(10, 10)
	bt.Set(7, 7)
	bt.Set(100, 100)
	bt.Set(20, 20)

	if str := bt.PreOrder(); str.String() != "[5 10 7 100 20]" {
		t.Errorf("Expected %v Got %v", "[5 10 7 100 20]", str.String())
	}
}

func TestPostOrder(t *testing.T) {
	bt := New(utils.IntComparator)

	if str := bt.PostOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}

	bt.Set(5, 5)
	bt.Set(10, 10)
	bt.Set(7, 7)
	bt.Set(100, 100)
	bt.Set(20, 20)

	if str := bt.PostOrder(); str.String() != "[7 20 100 10 5]" {
		t.Errorf("Expected %v Got %v", "[7 20 100 10 5]", str.String())
	}
}
