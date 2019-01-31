package redblacktree

import (
	"testing"

	"github.com/praveen001/ds/utils"
)

func TestRedBlackTree(t *testing.T) {
	rbt := New(utils.IntComparator)

	// Simple insert
	rbt.Set(11, "Eleven")
	if str := rbt.InOrder().String(); str != "[Eleven(black)]" {
		t.Errorf("Expected %v Got %v", "[Eleven(black)]", str)
	}

	// Simple insert
	rbt.Set(2, "Two")
	if str := rbt.InOrder().String(); str != "[Two(red) Eleven(black)]" {
		t.Errorf("Expected %v Got %v", "[Two(red) Eleven(black)]", str)
	}

	// Simple insert
	rbt.Set(14, "Fourteen")
	if str := rbt.InOrder().String(); str != "[Two(red) Eleven(black) Fourteen(red)]" {
		t.Errorf("Expected %v Got %v", "[Two(red) Eleven(black) Fourteen(red)]", str)
	}

	rbt.Set(1, "One") // INSERT CASE 3
	if str := rbt.InOrder().String(); str != "[One(red) Two(black) Eleven(black) Fourteen(black)]" {
		t.Errorf("Expected %v Got %v", "[One(red) Two(black) Eleven(black) Fourteen(black)]", str)
	}

	rbt.Set(7, "Seven") // INSERT CASE 2 // Simple
	if str := rbt.InOrder().String(); str != "[One(red) Two(black) Seven(red) Eleven(black) Fourteen(black)]" {
		t.Errorf("Expected %v Got %v", "[One(red) Two(black) Seven(red) Eleven(black) Fourteen(black)]", str)
	}

	rbt.Set(15, "Fifteen") // INSERT CASE 2
	if str := rbt.InOrder().String(); str != "[One(red) Two(black) Seven(red) Eleven(black) Fourteen(black) Fifteen(red)]" {
		t.Errorf("Expected %v Got %v", "[One(red) Two(black) Seven(red) Eleven(black) Fourteen(black) Fifteen(red)]", str)
	}

	rbt.Set(5, "Five") // INSERT CASE 3
	if str := rbt.InOrder().String(); str != "[One(black) Two(red) Five(red) Seven(black) Eleven(black) Fourteen(black) Fifteen(red)]" {
		t.Errorf("Expected %v Got %v", "[One(black) Two(red) Five(red) Seven(black) Eleven(black) Fourteen(black) Fifteen(red)]", str)
	}

	rbt.Set(8, "Eight") // INSERT CASE 2
	rbt.Set(4, "Four")  // INSERT CASE 3 ==> INSERT CASE 5

	if str := rbt.InOrder().String(); str != "[One(black) Two(red) Four(red) Five(black) Seven(black) Eight(black) Eleven(red) Fourteen(black) Fifteen(red)]" {
		t.Errorf("Expected %v Got %v", "[One(black) Two(red) Four(red) Five(black) Seven(black) Eight(black) Eleven(red) Fourteen(black) Fifteen(red)]", str)
	}

	// Simple red leaf delete
	rbt.Remove(4)
	if str := rbt.InOrder().String(); str != "[One(black) Two(red) Five(black) Seven(black) Eight(black) Eleven(red) Fourteen(black) Fifteen(red)]" {
		t.Errorf("Expected %v Got %v", "[One(black) Two(red) Five(black) Seven(black) Eight(black) Eleven(red) Fourteen(black) Fifteen(red)]", str)
	}

	// Case 4
	rbt.Remove(5)
	if str := rbt.InOrder().String(); str != "[One(red) Two(black) Seven(black) Eight(black) Eleven(red) Fourteen(black) Fifteen(red)]" {
		t.Errorf("Expected %v Got %v", "[One(red) Two(black) Seven(black) Eight(black) Eleven(red) Fourteen(black) Fifteen(red)]", str)
	}

	// Case 6
	rbt.Remove(8)
	if str := rbt.InOrder().String(); str != "[One(red) Two(black) Seven(black) Eleven(black) Fourteen(red) Fifteen(black)]" {
		t.Errorf("Expected %v Got %v", "[One(red) Two(black) Seven(black) Eleven(black) Fourteen(red) Fifteen(black)]", str)
	}

	// Simple red leaf delete
	rbt.Remove(1)
	if str := rbt.InOrder().String(); str != "[Two(black) Seven(black) Eleven(black) Fourteen(red) Fifteen(black)]" {
		t.Errorf("Expected %v Got %v", "[Two(black) Seven(black) Eleven(black) Fourteen(red) Fifteen(black)]", str)
	}

	// Case 2 => Case 4
	rbt.Remove(2)
	if str := rbt.InOrder().String(); str != "[Seven(black) Eleven(red) Fourteen(black) Fifteen(black)]" {
		t.Errorf("Expected %v Got %v", "[Seven(black) Eleven(red) Fourteen(black) Fifteen(black)]", str)
	}

	if ok := rbt.Contains(10); ok {
		t.Errorf("Expected %v Got %v", false, ok)
	}

	// Insert Right-Left Rotation Case
	rbt.Set(10, "Ten")
	if str := rbt.InOrder().String(); str != "[Seven(red) Ten(black) Eleven(red) Fourteen(black) Fifteen(black)]" {
		t.Errorf("Expected %v Got %v", "[Seven(red) Ten(black) Eleven(red) Fourteen(black) Fifteen(black)]", str)
	}

	if ok := rbt.Contains(10); !ok {
		t.Errorf("Expected %v Got %v", true, ok)
	}
}
