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

func (x *Node) inOrder(ll *linkedlist.LinkedList) {
	if hasLeft(x) {
		x.left.inOrder(ll)
	}

	ll.PushBack(x.value)

	if hasRight(x) {
		x.right.inOrder(ll)
	}
}

func (x *Node) preOrder(ll *linkedlist.LinkedList) {
	ll.PushBack(x.value)

	if hasLeft(x) {
		x.left.preOrder(ll)
	}

	if hasRight(x) {
		x.right.preOrder(ll)
	}
}

func (x *Node) postOrder(ll *linkedlist.LinkedList) {
	if hasLeft(x) {
		x.left.postOrder(ll)
	}

	if hasRight(x) {
		x.right.postOrder(ll)
	}

	ll.PushBack(x.value)
}

func hasLeft(x *Node) bool {
	if x == nil || x.left == nil {
		return false
	}
	return true
}

func hasRight(x *Node) bool {
	if x == nil || x.right == nil {
		return false
	}
	return true
}

func child(x *Node) *Node {
	if x == nil {
		return nil
	}
	c := x.left
	if c == nil {
		c = x.right
	}

	return c
}

func isLeaf(x *Node) bool {
	return !hasLeft(x) && !hasRight(x)
}
