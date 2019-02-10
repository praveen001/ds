package btree

import "testing"

func TestBTree(t *testing.T) {
	bt := New(func(a, b interface{}) int {
		astr := a.(string)
		bstr := b.(string)

		if []rune(astr)[0] > []rune(bstr)[0] {
			return 1
		} else if []rune(astr)[0] < []rune(bstr)[0] {
			return -1
		} else {
			return 0
		}
	}, 5)

	if bt == nil {
		t.Errorf("Expected a btree to be created, but got nil")
	}

	if s := bt.Len(); s != 0 {
		t.Errorf("Expected btree size to be 0, but got %d", s)
	}

	if ok := bt.contains("a"); ok == true {
		t.Errorf("Expected item to not exist, but it exists")
	}

	if m, ok := bt.Min(); ok != false {
		t.Errorf("Expected min to not exist, but got %v", m)
	}

	if m, ok := bt.Max(); ok != false {
		t.Errorf("Expected max to not exist, but got %v", m)
	}

	bt.Set("c", "c")

	if s := bt.Len(); s != 1 {
		t.Errorf("Expected btree size to be %v, but got %d", 1, s)
	}

	if ok := bt.contains("c"); ok != true {
		t.Errorf("Expected item to exist, but it does not exists")
	}

	if m, ok := bt.Min(); ok != true || m.Value().(string) != "c" {
		t.Errorf("Expected min to be %v, but got %v", "c", m)
	}

	if m, ok := bt.Max(); ok != true || m.Value().(string) != "c" {
		t.Errorf("Expected max to be %v, but got %v", "c", m)
	}

	bt.Set("n", "n")
	bt.Set("g", "g")
	bt.Set("a", "a")

	if s := bt.Len(); s != 4 {
		t.Errorf("Expected btree size to be %v, but got %d", 4, s)
	}

	if ok := bt.contains("a"); ok != true {
		t.Errorf("Expected item to exist, but it does not exists")
	}

	if m, ok := bt.Min(); ok != true || m.Value().(string) != "a" {
		t.Errorf("Expected min to be %v, but got %v", "a", m)
	}

	if m, ok := bt.Max(); ok != true || m.Value().(string) != "n" {
		t.Errorf("Expected max to be %v, but got %v", "n", m)
	}

	bt.Set("h", "h")

	if s := bt.Len(); s != 5 {
		t.Errorf("Expected btree size to be %v, but got %d", 5, s)
	}

	if s := bt.InOrder().String(); s != "[a c g h n]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c g h n]", s)
	}

	bt.Set("e", "e")
	bt.Set("k", "k")
	bt.Set("q", "q")

	if s := bt.Len(); s != 8 {
		t.Errorf("Expected btree size to be %v, but got %d", 8, s)
	}

	if ok := bt.contains("k"); ok != true {
		t.Errorf("Expected item to exist, but it does not exists")
	}

	if m, ok := bt.Min(); ok != true || m.Value().(string) != "a" {
		t.Errorf("Expected min to be %v, but got %v", "a", m)
	}

	if m, ok := bt.Max(); ok != true || m.Value().(string) != "q" {
		t.Errorf("Expected max to be %v, but got %v", "q", m)
	}

	if s := bt.InOrder().String(); s != "[a c e g h k n q]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c e g h k n q]", s)
	}

	bt.Set("m", "m")

	if s := bt.InOrder().String(); s != "[a c e g h k m n q]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c e g h k m n q]", s)
	}

	bt.Set("f", "f")
	bt.Set("w", "w")
	bt.Set("l", "l")
	bt.Set("t", "t")

	if s := bt.Len(); s != 13 {
		t.Errorf("Expected btree size to be %v, but got %d", 13, s)
	}

	if s := bt.InOrder().String(); s != "[a c e f g h k l m n q t w]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c e f g h k l m n q t w]", s)
	}

	bt.Set("z", "z")

	if s := bt.InOrder().String(); s != "[a c e f g h k l m n q t w z]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c e f g h k l m n q t w z]", s)
	}

	bt.Set("d", "d")
	bt.Set("p", "p")
	bt.Set("r", "r")
	bt.Set("x", "x")
	bt.Set("y", "y")
	bt.Set("s", "s")

	if s := bt.Len(); s != 20 {
		t.Errorf("Expected btree size to be %v, but got %d", 20, s)
	}

	if ok := bt.contains("k"); ok != true {
		t.Errorf("Expected item to exist, but it does not exists")
	}

	if m, ok := bt.Min(); ok != true || m.Value().(string) != "a" {
		t.Errorf("Expected min to be %v, but got %v", "a", m)
	}

	if m, ok := bt.Max(); ok != true || m.Value().(string) != "z" {
		t.Errorf("Expected max to be %v, but got %v", "z", m)
	}

	if s := bt.InOrder().String(); s != "[a c d e f g h k l m n p q r s t w x y z]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c d e f g h k l m n p q r s t w x y z]", s)
	}
}
