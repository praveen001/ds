package redblacktree

import (
	"testing"

	"github.com/praveen001/ds/utils"
)

func TestRedBlackTree(t *testing.T) {
	rbt := New(utils.IntComparator)

	// Simple insert
	rbt.Add(11)
	if str := rbt.InOrder().String(); str != "[11(black)]" {
		t.Errorf("Expected %v Got %v", "[11(black)]", str)
	}

	// Simple insert
	rbt.Add(2)
	if str := rbt.InOrder().String(); str != "[2(red) 11(black)]" {
		t.Errorf("Expected %v Got %v", "[2(red) 11(black)]", str)
	}

	// Simple insert
	rbt.Add(14)
	if str := rbt.InOrder().String(); str != "[2(red) 11(black) 14(red)]" {
		t.Errorf("Expected %v Got %v", "[2(red) 11(black) 14(red)]", str)
	}

	rbt.Add(1) // INSERT CASE 3
	if str := rbt.InOrder().String(); str != "[1(red) 2(black) 11(black) 14(black)]" {
		t.Errorf("Expected %v Got %v", "[1(red) 2(black) 11(black) 14(black)]", str)
	}

	rbt.Add(7) // INSERT CASE 2 // Simple
	if str := rbt.InOrder().String(); str != "[1(red) 2(black) 7(red) 11(black) 14(black)]" {
		t.Errorf("Expected %v Got %v", "[1(red) 2(black) 7(red) 11(black) 14(black)]", str)
	}

	rbt.Add(15) // INSERT CASE 2
	if str := rbt.InOrder().String(); str != "[1(red) 2(black) 7(red) 11(black) 14(black) 15(red)]" {
		t.Errorf("Expected %v Got %v", "[1(red) 2(black) 7(red) 11(black) 14(black) 15(red)]", str)
	}

	rbt.Add(5) // INSERT CASE 3
	if str := rbt.InOrder().String(); str != "[1(black) 2(red) 5(red) 7(black) 11(black) 14(black) 15(red)]" {
		t.Errorf("Expected %v Got %v", "[1(black) 2(red) 5(red) 7(black) 11(black) 14(black) 15(red)]", str)
	}

	rbt.Add(8) // INSERT CASE 2
	rbt.Add(4) // INSERT CASE 3 ==> INSERT CASE 5

	if str := rbt.InOrder().String(); str != "[1(black) 2(red) 4(red) 5(black) 7(black) 8(black) 11(red) 14(black) 15(red)]" {
		t.Errorf("Expected %v Got %v", "[1(black) 2(red) 4(red) 5(black) 7(black) 8(black) 11(red) 14(black) 15(red)]", str)
	}

	// Simple red leaf delete
	rbt.Remove(4)
	if str := rbt.InOrder().String(); str != "[1(black) 2(red) 5(black) 7(black) 8(black) 11(red) 14(black) 15(red)]" {
		t.Errorf("Expected %v Got %v", "[1(black) 2(red) 5(black) 7(black) 8(black) 11(red) 14(black) 15(red)]", str)
	}

	// Case 4
	rbt.Remove(5)
	if str := rbt.InOrder().String(); str != "[1(red) 2(black) 7(black) 8(black) 11(red) 14(black) 15(red)]" {
		t.Errorf("Expected %v Got %v", "[1(red) 2(black) 7(black) 8(black) 11(red) 14(black) 15(red)]", str)
	}

	// Case 6
	rbt.Remove(8)
	if str := rbt.InOrder().String(); str != "[1(red) 2(black) 7(black) 11(black) 14(red) 15(black)]" {
		t.Errorf("Expected %v Got %v", "[1(red) 2(black) 7(black) 11(black) 14(red) 15(black)]", str)
	}

	// Simple red leaf delete
	rbt.Remove(1)
	if str := rbt.InOrder().String(); str != "[2(black) 7(black) 11(black) 14(red) 15(black)]" {
		t.Errorf("Expected %v Got %v", "[2(black) 7(black) 11(black) 14(red) 15(black)]", str)
	}

	// Case 2 => Case 4
	rbt.Remove(2)
	if str := rbt.InOrder().String(); str != "[7(black) 11(red) 14(black) 15(black)]" {
		t.Errorf("Expected %v Got %v", "[7(black) 11(red) 14(black) 15(black)]", str)
	}

	if ok := rbt.Contains(10); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}

	// Insert Right-Left Rotation Case
	rbt.Add(10)
	if str := rbt.InOrder().String(); str != "[7(red) 10(black) 11(red) 14(black) 15(black)]" {
		t.Errorf("Expected %v Got %v", "[7(red) 10(black) 11(red) 14(black) 15(black)]", str)
	}

	if ok := rbt.Contains(10); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}
}
