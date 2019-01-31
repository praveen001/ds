package redblacktree

import (
	"fmt"

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

	ll.PushBack(fmt.Sprintf("%v(%v)", x.value, x.color))

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

func isLeft(x *Node) bool {
	if x.parent == nil {
		return false
	}
	if x.parent.right == x {
		return false
	}
	return true
}

func isRight(x *Node) bool {
	if x.parent == nil {
		return false
	}
	if x.parent.left == x {
		return false
	}
	return true
}

func isRoot(x *Node) bool {
	if x.parent == nil {
		return true
	}
	return false
}

func sibling(x *Node) *Node {
	if x.parent == nil {
		return nil
	}
	s := x.parent.left
	if s == x {
		s = x.parent.right
	}

	return s
}

func toRed(x *Node) {
	x.color = red
}

func toBlack(x *Node) {
	x.color = black
}

func isRed(x *Node) bool {
	if x != nil && x.color == red {
		return true
	}
	return false
}

func isBlack(x *Node) bool {
	if x == nil || x.color == black {
		return true
	}
	return false
}
