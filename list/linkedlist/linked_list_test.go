package linkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	ll := New()
	if ll == nil {
		t.Errorf("New didn't create new array list \n")
	}
}

func TestLen(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushFront(20)

	if l := ll.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}
}

func TestFront(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	if v, _ := ll.Front(); v != 10 {
		t.Errorf("Expected front to be %v, but got %v", 10, v)
	}
}

func TestBack(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	if v, _ := ll.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestPushFront(t *testing.T) {
	ll := New()

	ll.PushFront(10)
	ll.PushBack(20)
	ll.PushFront(-10)

	if v, _ := ll.Get(0); v != -10 {
		t.Errorf("Expected front to be %v, but got %v", -10, v)
	}
}

func TestPushBack(t *testing.T) {
	ll := New()

	ll.PushFront(10)
	ll.PushBack(20)
	ll.PushFront(-10)

	if v, _ := ll.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestSet(t *testing.T) {
	ll := New()

	ll.PushFront(10)
	ll.PushBack(20)
	ll.PushFront(-10)
	ll.Set(1, 100)

	if v, _ := ll.Get(1); v != 100 {
		t.Errorf("Expected value to be %v, but got %v", 100, v)
	}
}

func TestInsert(t *testing.T) {
	ll := New()

	ll.Insert(0, 10)
	if v, _ := ll.Get(0); v != 10 {
		t.Errorf("Expected value to be %v, but got %v", 10, v)
	}

	ll.Insert(1, 20)
	if v, _ := ll.Get(1); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}

	ll.Insert(1, 200)
	if v, _ := ll.Get(1); v != 200 {
		t.Errorf("Expected value to be %v, but got %v", 200, v)
	}

	if ok := ll.Insert(4, 30); ok {
		t.Errorf("Expected %v, but got %v", false, ok)
	}
}

func TestRemove(t *testing.T) {
	ll := New()

	ll.PushFront(10)
	ll.PushBack(20)
	ll.PushFront(-10)
	ll.Remove(1)

	if v, _ := ll.Get(1); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}
}

func TestAt(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	if v, _ := ll.Get(1); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}
}

func TestClear(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	ll.Clear()

	if l := ll.Len(); l != 0 {
		t.Errorf("Expected length to be %v, but got %v", 0, l)
	}
}

func TestPushBackList(t *testing.T) {
	ll := New()
	ll.PushBack(10)

	al2 := New()
	al2.PushBack(20)

	ll.PushBackList(al2)

	if l := ll.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}

	if v, _ := ll.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestPushFrontList(t *testing.T) {
	ll := New()
	ll.PushBack(10)

	al2 := New()
	al2.PushBack(20)

	ll.PushFrontList(al2)

	if l := ll.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}

	if v, _ := ll.Front(); v != 20 {
		t.Errorf("Expected front to be %v, but got %v", 20, v)
	}
}

func TestContains(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	if ok := ll.Contains(20); ok != true {
		t.Errorf("Expected %v, but got %v", true, ok)
	}
}

func TestIndexOf(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	if i := ll.IndexOf(20); i != 1 {
		t.Errorf("Expected %v, but got %v", 1, i)
	}
}

func TestValues(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	values := ll.Values()
	if values[0] != 10 || values[1] != 20 {
		t.Errorf("Expected %v, but got %v", []interface{}{10, 20}, values)
	}
}

func TestClone(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	cl := ll.Clone()
	cl.PushBack(30)

	if l := ll.Len(); l != 2 {
		t.Errorf("Expected %v, but got %v", 2, l)
	}
}
func TestSwap(t *testing.T) {
	ll := New()

	ll.PushBack(10)
	ll.PushBack(20)

	ll.Swap(0, 1)

	if a, _ := ll.Get(0); a != 20 {
		t.Errorf("Expected %v, but got %v", 20, a)
	}
}
