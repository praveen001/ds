package binarytree

import (
	"github.com/praveen001/ds/list/linkedlist"
)

// Key of the node
func (x *Node) Key() interface{} {
	return x.key
}

// Value of the node
func (x *Node) Value() interface{} {
	return x.value
}

func (x *Node) inOrderKeys(ll *linkedlist.LinkedList) {
	if x == nil {
		return
	}
	x.left.inOrderKeys(ll)
	ll.PushBack(x.key)
	x.right.inOrderKeys(ll)
}

func (x *Node) inOrder(ll *linkedlist.LinkedList) {
	if x == nil {
		return
	}
	x.left.inOrder(ll)
	ll.PushBack(x.value)
	x.right.inOrder(ll)
}

func (x *Node) preOrder(ll *linkedlist.LinkedList) {
	if x == nil {
		return
	}
	ll.PushBack(x.value)
	x.left.preOrder(ll)
	x.right.preOrder(ll)
}

func (x *Node) postOrder(ll *linkedlist.LinkedList) {
	if x == nil {
		return
	}
	x.left.postOrder(ll)
	x.right.postOrder(ll)
	ll.PushBack(x.value)
}

func (x *Node) hasLeft() bool {
	if x == nil || x.left == nil {
		return false
	}
	return true
}

func (x *Node) hasRight() bool {
	if x == nil || x.right == nil {
		return false
	}
	return true
}

func (x *Node) child() *Node {
	if x == nil {
		return nil
	}
	c := x.left
	if c == nil {
		c = x.right
	}

	return c
}

func (x *Node) isLeaf() bool {
	return !x.hasLeft() && !x.hasRight()
}
