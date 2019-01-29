package binarytree

import (
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/utils"
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
	if x.left != nil {
		x.left.inOrder(ll)
	}

	ll.Append(x.value)

	if x.right != nil {
		x.right.inOrder(ll)
	}
}

func (x *Node) preOrder(ll *linkedlist.LinkedList) {
	ll.Append(x.value)

	if x.left != nil {
		x.left.preOrder(ll)
	}

	if x.right != nil {
		x.right.preOrder(ll)
	}
}

func (x *Node) postOrder(ll *linkedlist.LinkedList) {
	if x.left != nil {
		x.left.postOrder(ll)
	}

	if x.right != nil {
		x.right.postOrder(ll)
	}

	ll.Append(x.value)
}

func (x *Node) recomputeHeight() {
	if x.left == nil && x.right == nil {
		x.height = 1
	} else if x.left == nil {
		x.height = x.right.height + 1
	} else if x.right == nil {
		x.height = x.left.height + 1
	} else {
		x.height = utils.MaxInt(x.left.height, x.right.height) + 1
	}
}

func hasLeft(x *Node) bool {
	if x.left == nil {
		return false
	}
	return true
}

func hasRight(x *Node) bool {
	if x.right == nil {
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
