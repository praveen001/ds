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
	if x == nil {
		return
	}
	x.left.inOrder(ll)
	ll.PushBack(fmt.Sprintf("%v(%v)", x.value, x.color))
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

func (x *Node) isLeft() bool {
	if x == nil || x.parent == nil {
		return false
	}
	if x.parent.right == x {
		return false
	}
	return true
}

func (x *Node) isRight() bool {
	if x == nil || x.parent == nil {
		return false
	}
	if x.parent.left == x {
		return false
	}
	return true
}

func (x *Node) isRoot() bool {
	if x == nil {
		return false
	}
	if x.parent == nil {
		return true
	}
	return false
}

func (x *Node) sibling() *Node {
	if x == nil || x.parent == nil {
		return nil
	}
	s := x.parent.left
	if s == x {
		s = x.parent.right
	}

	return s
}

func (x *Node) toRed() {
	if x == nil {
		return
	}
	x.color = red
}

func (x *Node) toBlack() {
	if x == nil {
		return
	}
	x.color = black
}

func (x *Node) isRed() bool {
	if x == nil || x.color != red {
		return false
	}
	return true
}

func (x *Node) isBlack() bool {
	if x == nil || x.color == black {
		return true
	}
	return false
}
