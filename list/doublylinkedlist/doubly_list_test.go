package doublylinkedlist

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	al := New()
	if al == nil {
		t.Errorf("New didn't create new array list \n")
	}
}

func TestAppend(t *testing.T) {
	al := New()

	// Empty append
	al.Append()
	if len := al.Length(); len != 0 {
		t.Errorf("Expected %v Got %v", 0, len)
	}

	// Single append
	al.Append(10)
	if len := al.Length(); len != 1 {
		t.Errorf("Expected %v Got %v", 1, len)
	}
	if val, ok := al.Get(0); val != 10 || !ok {
		t.Errorf("Expected %v Got %v", 10, val)
	}

	// Multiple append
	al.Append(20, 30)
	if len := al.Length(); len != 3 {
		t.Errorf("Expected %v Got %v", 3, len)
	}

	if val, ok := al.Get(al.Length() - 1); val != 30 || !ok {
		t.Errorf("Expected %v Got %v", 30, val)
	}
}

func TestPrepend(t *testing.T) {
	al := New()

	// Empty prepend
	al.Prepend()
	if len := al.Length(); len != 0 {
		t.Errorf("Expected %v Got %v", 0, len)
	}

	// Single prepend
	al.Prepend(10)
	if len := al.Length(); len != 1 {
		t.Errorf("Expected %v Got %v", 1, len)
	}
	if val, ok := al.Get(0); val != 10 || !ok {
		t.Errorf("Expected %v, %v got %v, %v", 10, true, val, ok)
	}

	// Multiple prepend
	al.Prepend(20, 30)
	if len := al.Length(); len != 3 {
		t.Errorf("Expected %v Got %v", 3, len)
	}
	if val, ok := al.Get(0); val != 20 || !ok {
		t.Errorf("Expected %v, %v got %v, %v", 20, true, val, ok)
	}
}

func TestGet(t *testing.T) {
	al := New()

	al.Append("aa", "bb", "cc")

	// Left out of bound
	if val, ok := al.Get(-1); val != nil || ok {
		t.Errorf("Expected %v, %v got %v, %v", nil, false, val, ok)
	}

	// Right out of bound
	if val, ok := al.Get(al.Length()); val != nil || ok {
		t.Errorf("Expected %v, %v got %v, %v", nil, false, val, ok)
	}

	// valid case
	if val, ok := al.Get(0); val != "aa" || !ok {
		t.Errorf("Expected %v, %v got %v %v", "aa", true, val, ok)
	}
}

func TestSet(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)

	// Invalid case
	if ok := al.Set(-1, 0); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}

	// Valid case
	if ok := al.Set(1, 1000); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}
	if val, ok := al.Get(1); val != 1000 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 1000, true, val, ok)
	}
}

func TestRemove(t *testing.T) {
	al := New()

	al.Append(10.5, 11.5, 12.5)

	// Invalid case
	if val, ok := al.Remove(-1); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}

	// Valid Case
	if val, ok := al.Remove(1); val == nil || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 11.5, true, val, ok)
	}
	if len := al.Length(); len != 2 {
		t.Errorf("Expected %v Got %v", 2, len)
	}
	if val, ok := al.Get(1); val == nil || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 12.5, true, val, ok)
	}
}

func TestContains(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)

	// Not found
	if ok := al.Contains(40); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}

	// Found
	if ok := al.Contains(30); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}
}

func TestIndexOf(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)

	// Not found
	if idx := al.IndexOf(40); idx != -1 {
		t.Errorf("Expected %v Got %v", -1, idx)
	}

	// Found
	if idx := al.IndexOf(30); idx != 2 {
		t.Errorf("Expected %v Got %v", 2, idx)
	}
}

func TestValues(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)
	values := al.Values()

	if l := len(values); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if values[0] != 10 {
		t.Errorf("Expected %v Got %v", 10, values[0])
	}

	if values[2] != 30 {
		t.Errorf("Expected %v Got %v", 30, values[2])
	}
}

func TestLength(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)
	if len := al.Length(); len != 3 {
		t.Errorf("Expected %v Got %v", 3, len)
	}
}

func TestClear(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)
	al.Clear()

	if len := al.Length(); len != 0 {
		t.Errorf("Expected %v Got %v", 0, len)
	}

	if val, ok := al.Get(0); val != nil || ok {
		t.Errorf("Expected %v, %v Got %v, %v", nil, false, val, ok)
	}
}

func TestWithInRange(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)

	// Invalid case
	if ok := al.WithInRange(-1); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}

	// Valid case
	if ok := al.WithInRange(2); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}
}

func TestString(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)

	if str := al.String(); str != "[10 20 30]" {
		t.Errorf("Expected %v Got %v", "[10 20 30]", str)
	}
}

func TestFilter(t *testing.T) {
	al := New()

	al.Append(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	nal := al.Filter(func(a interface{}) bool {
		if a.(int)%2 == 0 {
			return true
		}
		return false
	})

	if l := nal.Length(); l != 5 {
		t.Errorf("Expected %v Got %v", 5, l)
	}

	if l := al.Length(); l != 10 {
		t.Errorf("Expected %v Got %v", 10, l)
	}
}

func TestMap(t *testing.T) {
	al := New()

	al.Append(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	nal := al.Map(func(a interface{}) interface{} {
		return a.(int) * 2
	})

	if l := nal.Length(); l != 10 {
		t.Errorf("Expected %v Got %v", 10, l)
	}

	if val, ok := nal.Get(0); val != 2 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 2, true, val, ok)
	}

	if val, ok := al.Get(0); val != 1 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 1, true, val, ok)
	}
}

func TestConcat(t *testing.T) {
	al := New()
	al2 := New()
	al3 := New()

	al.Append(10, 20, 30)
	al2.Append(40, 50, 60)
	al3.Append(70, 80, 90)

	nal := al.Concat(al2, al3)

	if l := nal.Length(); l != 9 {
		t.Errorf("Expected %v Got %v", 9, nal)
	}
}

func TestClone(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)
	nal := al.Clone()

	if l := nal.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}

	if val, ok := nal.Get(0); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	nal.Clear()

	if l := al.Length(); l != 3 {
		t.Errorf("Expected %v Got %v", 3, l)
	}
}

func TestReverse(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)
	al.Prepend(-10, 0)
	nal := al.Reverse()
	fmt.Println(nal)

	if val, ok := nal.Get(0); val != 30 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 30, true, val, ok)
	}

	if val, ok := nal.Get(2); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}

	if l := nal.Length(); l != 5 {
		t.Errorf("Expected %v Got %v", 5, l)
	}
}

func TestSwap(t *testing.T) {
	al := New()

	al.Append(10, 20, 30)
	al.Swap(0, 2)

	if val, ok := al.Get(0); val != 30 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 30, true, val, ok)
	}

	if val, ok := al.Get(2); val != 10 || !ok {
		t.Errorf("Expected %v, %v Got %v, %v", 10, true, val, ok)
	}
}
