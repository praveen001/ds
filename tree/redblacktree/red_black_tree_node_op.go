package redblacktree

import (
	"fmt"

	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/utils"
)

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

func (x *Node) isRed() bool {
	if x.color == red {
		return true
	}
	return false
}

func (x *Node) isBlack() bool {
	if x.color == black {
		return true
	}
	return false
}

func (x *Node) toRed() {
	x.color = red
}

func (x *Node) toBlack() {
	x.color = black
}

func (x *Node) isRoot() bool {
	if x.parent == nil {
		return true
	}
	return false
}

func (x *Node) leftRotate() *Node {
	y := x.right
	z := y.left

	y.left = x
	y.parent = x.parent

	x.right = z
	x.parent = y

	if z != nil {
		z.parent = x
		z.recomputeHeight()
	}

	x.recomputeHeight()
	y.recomputeHeight()

	return y
}

func (x *Node) rightRotate() *Node {
	y := x.left
	z := y.right

	y.right = x
	y.parent = x.parent

	x.left = z
	x.parent = y

	if z != nil {
		z.parent = x
	}

	x.recomputeHeight()
	y.recomputeHeight()

	return y
}
