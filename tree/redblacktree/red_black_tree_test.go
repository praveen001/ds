package redblacktree

import (
	"testing"

	"github.com/praveen001/ds/utils"
)

func TestRedBlackTree(t *testing.T) {
	rbt := New(utils.IntComparator)

	// Simple insert
	rbt.Add(11)
	if str := rbt.InOrder().String(); str != "[11(black, 1)]" {
		t.Errorf("Expected %v Got %v", "[11(black, 1)]", str)
	}

	// Simple insert
	rbt.Add(2)
	if str := rbt.InOrder().String(); str != "[2(red, 1) 11(black, 2)]" {
		t.Errorf("Expected %v Got %v", "[2(red, 1) 11(black, 2)]", str)
	}

	// Simple insert
	rbt.Add(14)
	if str := rbt.InOrder().String(); str != "[2(red, 1) 11(black, 2) 14(red, 1)]" {
		t.Errorf("Expected %v Got %v", "[2(red, 1) 11(black, 2) 14(red, 1)]", str)
	}

	rbt.Add(1) // INSERT CASE 3
	if str := rbt.InOrder().String(); str != "[1(red, 1) 2(black, 2) 11(black, 3) 14(black, 1)]" {
		t.Errorf("Expected %v Got %v", "[1(red, 1) 2(black, 2) 11(black, 3) 14(black, 1)]", str)
	}

	rbt.Add(7) // INSERT CASE 2 // Simple
	if str := rbt.InOrder().String(); str != "[1(red, 1) 2(black, 2) 7(red, 1) 11(black, 3) 14(black, 1)]" {
		t.Errorf("Expected %v Got %v", "[1(red, 1) 2(black, 2) 7(red, 1) 11(black, 3) 14(black, 1)]", str)
	}

	rbt.Add(15) // INSERT CASE 2
	if str := rbt.InOrder().String(); str != "[1(red, 1) 2(black, 2) 7(red, 1) 11(black, 3) 14(black, 2) 15(red, 1)]" {
		t.Errorf("Expected %v Got %v", "[1(red, 1) 2(black, 2) 7(red, 1) 11(black, 3) 14(black, 2) 15(red, 1)]", str)
	}

	rbt.Add(5) // INSERT CASE 3
	if str := rbt.InOrder().String(); str != "[1(black, 1) 2(red, 3) 5(red, 1) 7(black, 2) 11(black, 4) 14(black, 2) 15(red, 1)]" {
		t.Errorf("Expected %v Got %v", "[1(black, 1) 2(red, 3) 5(red, 1) 7(black, 2) 11(black, 4) 14(black, 2) 15(red, 1)]", str)
	}

	rbt.Add(8) // INSERT CASE 2
	rbt.Add(4) // INSERT CASE 3 ==> INSERT CASE 5

	if str := rbt.InOrder().String(); str != "[1(black, 1) 2(red, 3) 4(red, 1) 5(black, 2) 7(black, 4) 8(black, 1) 11(red, 3) 14(black, 2) 15(red, 1)]" {
		t.Errorf("Expected %v Got %v", "[1(black, 1) 2(red, 3) 4(red, 1) 5(black, 2) 7(black, 4) 8(black, 1) 11(red, 3) 14(black, 2) 15(red, 1)]", str)
	}

	// Simple red leaf delete
	rbt.Remove(4)
	if str := rbt.InOrder().String(); str != "[1(black, 1) 2(red, 3) 5(black, 2) 7(black, 4) 8(black, 1) 11(red, 3) 14(black, 2) 15(red, 1)]" {
		t.Errorf("Expected %v Got %v", "[1(black, 1) 2(red, 3) 5(black, 2) 7(black, 4) 8(black, 1) 11(red, 3) 14(black, 2) 15(red, 1)]", str)
	}

	// Case 4
	rbt.Remove(5)
	if str := rbt.InOrder().String(); str != "[1(red, 1) 2(black, 3) 7(black, 4) 8(black, 1) 11(red, 3) 14(black, 2) 15(red, 1)]" {
		t.Errorf("Expected %v Got %v", "[1(red, 1) 2(black, 3) 7(black, 4) 8(black, 1) 11(red, 3) 14(black, 2) 15(red, 1)]", str)
	}

	// Case 6
	rbt.Remove(8)
	if str := rbt.InOrder().String(); str != "[1(red, 1) 2(black, 3) 7(black, 4) 11(black, 2) 14(red, 3) 15(black, 1)]" {
		t.Errorf("Expected %v Got %v", "[1(red, 1) 2(black, 3) 7(black, 4) 11(black, 2) 14(red, 3) 15(black, 1)]", str)
	}

	// Simple red leaf delete
	rbt.Remove(1)
	if str := rbt.InOrder().String(); str != "[2(black, 3) 7(black, 4) 11(black, 2) 14(red, 3) 15(black, 1)]" {
		t.Errorf("Expected %v Got %v", "[2(black, 3) 7(black, 4) 11(black, 2) 14(red, 3) 15(black, 1)]", str)
	}

	// Case 2 => Case 4
	rbt.Remove(2)
	if str := rbt.InOrder().String(); str != "[7(black, 4) 11(red, 1) 14(black, 5) 15(black, 1)]" {
		t.Errorf("Expected %v Got %v", "[7(black, 4) 11(red, 1) 14(black, 5) 15(black, 1)]", str)
	}

	// Insert Right-Left Rotation Case
	rbt.Add(10)
	if str := rbt.InOrder().String(); str != "[7(red, 1) 10(black, 2) 11(red, 1) 14(black, 5) 15(black, 1)]" {
		t.Errorf("Expected %v Got %v", "[7(red, 1) 10(black, 2) 11(red, 1) 14(black, 5) 15(black, 1)]", str)
	}
}
