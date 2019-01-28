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

func TestAdd(t *testing.T) {
	bt := New(utils.IntComparator)

	bt.Add(10)

	if l := bt.Length(); l != 1 {
		t.Errorf("Expected %v Got %v", 1, l)
	}

	bt.Add(10)
	bt.Add(20)
	if l := bt.Length(); l != 2 {
		t.Errorf("Expected %v Got %v", 2, l)
	}
}

func TestRemove(t *testing.T) {
	bt := New(utils.IntComparator)

	bt.Add(10)
	bt.Add(20)

	if ok := bt.Remove(10); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}

	if ok := bt.Remove(10); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}
}

func TestHeight(t *testing.T) {
	bt := New(utils.IntComparator)

	if h := bt.Height(); h != 0 {
		t.Errorf("Expected %v Got %v", 0, h)
	}

	bt.Add(10)
	bt.Add(20)
	bt.Add(30)

	if h := bt.Height(); h != 3 {
		t.Errorf("Expected %v Got %v", 3, h)
	}
}

func TestMin(t *testing.T) {
	bt := New(utils.IntComparator)

	if m, ok := bt.Min(); m != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, m, ok)
	}

	bt.Add(100)
	bt.Add(10)
	bt.Add(20)
	bt.Add(5)

	if m, ok := bt.Min(); m != 5 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 5, true, m, ok)
	}
}

func TestMax(t *testing.T) {
	bt := New(utils.IntComparator)

	if m, ok := bt.Max(); m != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, m, ok)
	}

	bt.Add(5)
	bt.Add(10)
	bt.Add(100)
	bt.Add(20)

	if m, ok := bt.Max(); m != 100 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 100, true, m, ok)
	}
}

func TestContains(t *testing.T) {
	bt := New(utils.IntComparator)

	if ok := bt.Contains(1); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}

	bt.Add(5)
	bt.Add(10)
	bt.Add(100)
	bt.Add(20)

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

	bt.Add(5)
	bt.Add(10)
	bt.Add(100)
	bt.Add(20)

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

	bt.Add(5)
	bt.Add(10)
	bt.Add(100)
	bt.Add(20)

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

	bt.Add(5)
	bt.Add(10)
	bt.Add(100)
	bt.Add(20)

	if str := bt.InOrder(); str.String() != "[5 10 20 100]" {
		t.Errorf("Expected %v Got %v", "[5 10 20 100]", str.String())
	}
}

func TestPreOrder(t *testing.T) {
	bt := New(utils.IntComparator)

	if str := bt.PreOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}

	bt.Add(5)
	bt.Add(10)
	bt.Add(7)
	bt.Add(100)
	bt.Add(20)

	if str := bt.PreOrder(); str.String() != "[5 10 7 100 20]" {
		t.Errorf("Expected %v Got %v", "[5 10 7 100 20]", str.String())
	}
}

func TestPostOrder(t *testing.T) {
	bt := New(utils.IntComparator)

	if str := bt.PostOrder(); str.String() != "[]" {
		t.Errorf("Expected %v Got %v", "[]", str.String())
	}

	bt.Add(5)
	bt.Add(10)
	bt.Add(7)
	bt.Add(100)
	bt.Add(20)

	if str := bt.PostOrder(); str.String() != "[7 20 100 10 5]" {
		t.Errorf("Expected %v Got %v", "[7 20 100 10 5]", str.String())
	}
}
