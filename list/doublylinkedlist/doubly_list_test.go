package doublylinkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	dl := New()
	if dl == nil {
		t.Errorf("New didn't create new array list \n")
	}
}

func TestLen(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushFront(20)

	if l := dl.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}
}

func TestFront(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	if v, _ := dl.Front(); v != 10 {
		t.Errorf("Expected front to be %v, but got %v", 10, v)
	}
}

func TestBack(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	if v, _ := dl.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestPushFront(t *testing.T) {
	dl := New()

	dl.PushFront(10)
	dl.PushBack(20)
	dl.PushFront(-10)

	if v, _ := dl.Get(0); v != -10 {
		t.Errorf("Expected front to be %v, but got %v", -10, v)
	}
}

func TestPushBack(t *testing.T) {
	dl := New()

	dl.PushFront(10)
	dl.PushBack(20)
	dl.PushFront(-10)

	if v, _ := dl.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestSet(t *testing.T) {
	dl := New()

	dl.PushFront(10)
	dl.PushBack(20)
	dl.PushFront(-10)
	dl.Set(1, 100)

	if v, _ := dl.Get(1); v != 100 {
		t.Errorf("Expected value to be %v, but got %v", 100, v)
	}
}

func TestInsert(t *testing.T) {
	dl := New()

	dl.Insert(0, 10)
	if v, _ := dl.Get(0); v != 10 {
		t.Errorf("Expected value to be %v, but got %v", 10, v)
	}

	dl.Insert(1, 20)
	if v, _ := dl.Get(1); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}

	dl.Insert(1, 200)
	if v, _ := dl.Get(1); v != 200 {
		t.Errorf("Expected value to be %v, but got %v", 200, v)
	}

	if ok := dl.Insert(4, 30); ok {
		t.Errorf("Expected %v, but got %v", false, ok)
	}
}

func TestRemove(t *testing.T) {
	dl := New()

	dl.PushFront(10)
	dl.PushBack(20)
	dl.PushFront(-10)
	dl.Remove(1)

	if v, _ := dl.Get(1); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}
}

func TestAt(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	if v, _ := dl.Get(1); v != 20 {
		t.Errorf("Expected value to be %v, but got %v", 20, v)
	}
}

func TestClear(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	dl.Clear()

	if l := dl.Len(); l != 0 {
		t.Errorf("Expected length to be %v, but got %v", 0, l)
	}
}

func TestPushBackList(t *testing.T) {
	dl := New()
	dl.PushBack(10)

	al2 := New()
	al2.PushBack(20)

	dl.PushBackList(al2)

	if l := dl.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}

	if v, _ := dl.Back(); v != 20 {
		t.Errorf("Expected back to be %v, but got %v", 20, v)
	}
}

func TestPushFrontList(t *testing.T) {
	dl := New()
	dl.PushBack(10)

	al2 := New()
	al2.PushBack(20)

	dl.PushFrontList(al2)

	if l := dl.Len(); l != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, l)
	}

	if v, _ := dl.Front(); v != 20 {
		t.Errorf("Expected front to be %v, but got %v", 20, v)
	}
}

func TestContains(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	if ok := dl.Contains(20); ok != true {
		t.Errorf("Expected %v, but got %v", true, ok)
	}
}

func TestIndexOf(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	if i := dl.IndexOf(20); i != 1 {
		t.Errorf("Expected %v, but got %v", 1, i)
	}
}

func TestValues(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	values := dl.Values()
	if values[0] != 10 || values[1] != 20 {
		t.Errorf("Expected %v, but got %v", []interface{}{10, 20}, values)
	}
}

func TestClone(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	cl := dl.Clone()
	cl.PushBack(30)

	if l := dl.Len(); l != 2 {
		t.Errorf("Expected %v, but got %v", 2, l)
	}
}

func TestSwap(t *testing.T) {
	dl := New()

	dl.PushBack(10)
	dl.PushBack(20)

	dl.Swap(0, 1)

	if a, _ := dl.Get(0); a != 20 {
		t.Errorf("Expected %v, but got %v", 20, a)
	}
}
