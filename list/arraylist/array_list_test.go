package arraylist

import (
	"testing"
)

func TestNew(t *testing.T) {
	al := New()
	if al == nil {
		t.Errorf("New didn't create new array list \n")
	}
}

func TestLen(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushFront(20)

	if l := al.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}
}

func TestFront(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	if v, _ := al.Front(); v != 10 {
		t.Errorf("Expected front to be %v, but got %v", 10, v)
	}
}

func TestBack(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	if v, _ := al.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestPushFront(t *testing.T) {
	al := New()

	al.PushFront(10)
	al.PushBack(20)
	al.PushFront(-10)

	if v, _ := al.Get(0); v != -10 {
		t.Errorf("Expected front to be %v, but got %v", -10, v)
	}
}

func TestPushBack(t *testing.T) {
	al := New()

	al.PushFront(10)
	al.PushBack(20)
	al.PushFront(-10)

	if v, _ := al.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestSet(t *testing.T) {
	al := New()

	al.PushFront(10)
	al.PushBack(20)
	al.PushFront(-10)
	al.Set(1, 100)

	if v, _ := al.Get(1); v != 100 {
		t.Errorf("Expected value to be %v, but got %v", 100, v)
	}
}

func TestInsert(t *testing.T) {
	al := New()

	al.Insert(0, 10)
	if v, _ := al.Get(0); v != 10 {
		t.Errorf("Expected value to be %v, but got %v", 10, v)
	}

	al.Insert(0, 20)
	if v, _ := al.Get(0); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}

	if ok := al.Insert(3, 30); ok {
		t.Errorf("Expected %v, but got %v", false, ok)
	}
}

func TestRemove(t *testing.T) {
	al := New()

	al.PushFront(10)
	al.PushBack(20)
	al.PushFront(-10)
	al.Remove(1)

	if v, _ := al.Get(1); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}
}

func TestAt(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	if v, _ := al.Get(1); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}
}

func TestClear(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	al.Clear()

	if l := al.Len(); l != 0 {
		t.Errorf("Expected length to be %v, but got %v", 0, l)
	}
}

func TestPushBackList(t *testing.T) {
	al := New()
	al.PushBack(10)

	al2 := New()
	al2.PushBack(20)

	al.PushBackList(al2)

	if l := al.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}

	if v, _ := al.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestPushFrontList(t *testing.T) {
	al := New()
	al.PushBack(10)

	al2 := New()
	al2.PushBack(20)

	al.PushFrontList(al2)

	if l := al.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}

	if v, _ := al.Front(); v != 20 {
		t.Errorf("Expected front to be %v, but got %v", 20, v)
	}
}

func TestContains(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	if ok := al.Contains(20); ok != true {
		t.Errorf("Expected %v, but got %v", true, ok)
	}
}

func TestIndexOf(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	if i := al.IndexOf(20); i != 1 {
		t.Errorf("Expected %v, but got %v", 1, i)
	}
}

func TestValues(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	values := al.Values()
	if values[0] != 10 || values[1] != 20 {
		t.Errorf("Expected %v, but got %v", []interface{}{10, 20}, values)
	}
}

func TestClone(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	cl := al.Clone()
	cl.PushBack(30)

	if l := al.Len(); l != 2 {
		t.Errorf("Expected %v, but got %v", 2, l)
	}
}
func TestSwap(t *testing.T) {
	al := New()

	al.PushBack(10)
	al.PushBack(20)

	al.Swap(0, 1)

	if a, _ := al.Get(0); a != 20 {
		t.Errorf("Expected %v, but got %v", 20, a)
	}
}
