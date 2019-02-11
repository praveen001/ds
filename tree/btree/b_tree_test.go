package btree

import (
	"testing"

	"github.com/praveen001/ds/utils"
)

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

	CheckNode(t, bt.compare, bt.root, 1, []interface{}{"c"})

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

	CheckNode(t, bt.compare, bt.root, 4, []interface{}{"a", "c", "g", "n"})

	bt.Set("h", "h")

	if s := bt.Len(); s != 5 {
		t.Errorf("Expected btree size to be %v, but got %d", 5, s)
	}

	CheckNode(t, bt.compare, bt.root, 1, []interface{}{"g"})
	CheckNode(t, bt.compare, bt.root.children[0], 2, []interface{}{"a", "c"})
	CheckNode(t, bt.compare, bt.root.children[1], 2, []interface{}{"h", "n"})

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

	CheckNode(t, bt.compare, bt.root, 1, []interface{}{"g"})
	CheckNode(t, bt.compare, bt.root.children[0], 3, []interface{}{"a", "c", "e"})
	CheckNode(t, bt.compare, bt.root.children[1], 4, []interface{}{"h", "k", "n", "q"})

	bt.Set("m", "m")

	CheckNode(t, bt.compare, bt.root, 2, []interface{}{"g", "m"})
	CheckNode(t, bt.compare, bt.root.children[0], 3, []interface{}{"a", "c", "e"})
	CheckNode(t, bt.compare, bt.root.children[1], 2, []interface{}{"h", "k"})
	CheckNode(t, bt.compare, bt.root.children[2], 2, []interface{}{"n", "q"})

	bt.Set("f", "f")
	bt.Set("w", "w")
	bt.Set("l", "l")
	bt.Set("t", "t")

	if s := bt.Len(); s != 13 {
		t.Errorf("Expected btree size to be %v, but got %d", 13, s)
	}

	CheckNode(t, bt.compare, bt.root, 2, []interface{}{"g", "m"})
	CheckNode(t, bt.compare, bt.root.children[0], 4, []interface{}{"a", "c", "e", "f"})
	CheckNode(t, bt.compare, bt.root.children[1], 3, []interface{}{"h", "k", "l"})
	CheckNode(t, bt.compare, bt.root.children[2], 4, []interface{}{"n", "q", "t", "w"})

	bt.Set("z", "z")

	CheckNode(t, bt.compare, bt.root, 3, []interface{}{"g", "m", "t"})
	CheckNode(t, bt.compare, bt.root.children[0], 4, []interface{}{"a", "c", "e", "f"})
	CheckNode(t, bt.compare, bt.root.children[1], 3, []interface{}{"h", "k", "l"})
	CheckNode(t, bt.compare, bt.root.children[2], 2, []interface{}{"n", "q"})
	CheckNode(t, bt.compare, bt.root.children[3], 2, []interface{}{"w", "z"})

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

	CheckNode(t, bt.compare, bt.root, 1, []interface{}{"m"})
	CheckNode(t, bt.compare, bt.root.children[0], 2, []interface{}{"d", "g"})
	CheckNode(t, bt.compare, bt.root.children[1], 2, []interface{}{"q", "t"})
	CheckNode(t, bt.compare, bt.root.children[0].children[0], 2, []interface{}{"a", "c"})
	CheckNode(t, bt.compare, bt.root.children[0].children[1], 2, []interface{}{"e", "f"})
	CheckNode(t, bt.compare, bt.root.children[0].children[2], 3, []interface{}{"h", "k", "l"})
	CheckNode(t, bt.compare, bt.root.children[1].children[0], 2, []interface{}{"n", "p"})
	CheckNode(t, bt.compare, bt.root.children[1].children[1], 2, []interface{}{"r", "s"})
	CheckNode(t, bt.compare, bt.root.children[1].children[2], 4, []interface{}{"w", "x", "y", "z"})

	// Deletion
	bt.Remove("h")

	if s := bt.Len(); s != 19 {
		t.Errorf("Expected btree size to be %v, but got %d", 19, s)
	}

	if s := bt.InOrder().String(); s != "[a c d e f g k l m n p q r s t w x y z]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c d e f g k l m n p q r s t w x y z]", s)
	}

	CheckNode(t, bt.compare, bt.root, 1, []interface{}{"m"})
	CheckNode(t, bt.compare, bt.root.children[0], 2, []interface{}{"d", "g"})
	CheckNode(t, bt.compare, bt.root.children[1], 2, []interface{}{"q", "t"})
	CheckNode(t, bt.compare, bt.root.children[0].children[0], 2, []interface{}{"a", "c"})
	CheckNode(t, bt.compare, bt.root.children[0].children[1], 2, []interface{}{"e", "f"})
	CheckNode(t, bt.compare, bt.root.children[0].children[2], 2, []interface{}{"k", "l"})
	CheckNode(t, bt.compare, bt.root.children[1].children[0], 2, []interface{}{"n", "p"})
	CheckNode(t, bt.compare, bt.root.children[1].children[1], 2, []interface{}{"r", "s"})
	CheckNode(t, bt.compare, bt.root.children[1].children[2], 4, []interface{}{"w", "x", "y", "z"})

	bt.Remove("t")

	if s := bt.Len(); s != 18 {
		t.Errorf("Expected btree size to be %v, but got %d", 18, s)
	}

	if ok := bt.contains("t"); ok == true {
		t.Errorf("Expected item to not exist, but it does exist")
	}

	if s := bt.InOrder().String(); s != "[a c d e f g k l m n p q r s w x y z]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c d e f g k l m n p q r s w x y z]", s)
	}

	CheckNode(t, bt.compare, bt.root, 1, []interface{}{"m"})
	CheckNode(t, bt.compare, bt.root.children[0], 2, []interface{}{"d", "g"})
	CheckNode(t, bt.compare, bt.root.children[1], 2, []interface{}{"q", "w"})
	CheckNode(t, bt.compare, bt.root.children[0].children[0], 2, []interface{}{"a", "c"})
	CheckNode(t, bt.compare, bt.root.children[0].children[1], 2, []interface{}{"e", "f"})
	CheckNode(t, bt.compare, bt.root.children[0].children[2], 2, []interface{}{"k", "l"})
	CheckNode(t, bt.compare, bt.root.children[1].children[0], 2, []interface{}{"n", "p"})
	CheckNode(t, bt.compare, bt.root.children[1].children[1], 2, []interface{}{"r", "s"})
	CheckNode(t, bt.compare, bt.root.children[1].children[2], 3, []interface{}{"x", "y", "z"})

	bt.Remove("r")

	if s := bt.Len(); s != 17 {
		t.Errorf("Expected btree size to be %v, but got %d", 17, s)
	}

	if ok := bt.contains("r"); ok == true {
		t.Errorf("Expected item to not exist, but it does exist")
	}

	if s := bt.InOrder().String(); s != "[a c d e f g k l m n p q s w x y z]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c d e f g k l m n p q s w x y z]", s)
	}

	CheckNode(t, bt.compare, bt.root, 1, []interface{}{"m"})
	CheckNode(t, bt.compare, bt.root.children[0], 2, []interface{}{"d", "g"})
	CheckNode(t, bt.compare, bt.root.children[1], 2, []interface{}{"q", "x"})
	CheckNode(t, bt.compare, bt.root.children[0].children[0], 2, []interface{}{"a", "c"})
	CheckNode(t, bt.compare, bt.root.children[0].children[1], 2, []interface{}{"e", "f"})
	CheckNode(t, bt.compare, bt.root.children[0].children[2], 2, []interface{}{"k", "l"})
	CheckNode(t, bt.compare, bt.root.children[1].children[0], 2, []interface{}{"n", "p"})
	CheckNode(t, bt.compare, bt.root.children[1].children[1], 2, []interface{}{"s", "w"})
	CheckNode(t, bt.compare, bt.root.children[1].children[2], 2, []interface{}{"y", "z"})

	bt.Remove("e")

	if s := bt.Len(); s != 16 {
		t.Errorf("Expected btree size to be %v, but got %d", 16, s)
	}

	if ok := bt.contains("e"); ok == true {
		t.Errorf("Expected item to not exist, but it does exist")
	}

	if s := bt.InOrder().String(); s != "[a c d f g k l m n p q s w x y z]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c d f g k l m n p q s w x y z]", s)
	}

	CheckNode(t, bt.compare, bt.root, 4, []interface{}{"g", "m", "q", "x"})
	CheckNode(t, bt.compare, bt.root.children[0], 4, []interface{}{"a", "c", "d", "f"})
	CheckNode(t, bt.compare, bt.root.children[1], 2, []interface{}{"k", "l"})
	CheckNode(t, bt.compare, bt.root.children[2], 2, []interface{}{"n", "p"})
	CheckNode(t, bt.compare, bt.root.children[3], 2, []interface{}{"s", "w"})
	CheckNode(t, bt.compare, bt.root.children[4], 2, []interface{}{"y", "z"})

	bt.Remove("k")

	if s := bt.Len(); s != 15 {
		t.Errorf("Expected btree size to be %v, but got %d", 15, s)
	}

	if ok := bt.contains("k"); ok == true {
		t.Errorf("Expected item to not exist, but it does exist")
	}

	if s := bt.InOrder().String(); s != "[a c d f g l m n p q s w x y z]" {
		t.Errorf("Expected ordering to be %v, but got %v", "[a c d f g l m n p q s w x y z]", s)
	}

	CheckNode(t, bt.compare, bt.root, 4, []interface{}{"f", "m", "q", "x"})
	CheckNode(t, bt.compare, bt.root.children[0], 3, []interface{}{"a", "c", "d"})
	CheckNode(t, bt.compare, bt.root.children[1], 2, []interface{}{"g", "l"})
	CheckNode(t, bt.compare, bt.root.children[2], 2, []interface{}{"n", "p"})
	CheckNode(t, bt.compare, bt.root.children[3], 2, []interface{}{"s", "w"})
	CheckNode(t, bt.compare, bt.root.children[4], 2, []interface{}{"y", "z"})
}

func CheckNode(t *testing.T, comparator utils.CompareFunc, node *Node, size int, keys []interface{}) {
	if l := len(node.entries); l != size {
		t.Errorf(" Expected number of keys in node to be %v, but got %v ", size, l)
	}

	for i, e := range node.entries {
		if comp := comparator(e.key, keys[i]); comp != 0 {
			t.Errorf(" Expected key no. %v to be %v, but got %v", i, keys[i], e.key)
		}
	}
}
