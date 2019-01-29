package redblacktree

import (
	"fmt"

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

	ll.Append(fmt.Sprintf("%v(%v, %v)", x.value, x.color, x.height))

	if x.right != nil {
		x.right.inOrder(ll)
	}
}

func (x *Node) preOrder(ll *linkedlist.LinkedList) {
	ll.Append([]interface{}{x.value, x.height})

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

	ll.Append([]interface{}{x.value, x.height})
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

func toRed(x *Node) {
	x.color = red
}

func toBlack(x *Node) {
	x.color = black
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

func isRoot(x *Node) bool {
	if x.parent == nil {
		return true
	}
	return false
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
