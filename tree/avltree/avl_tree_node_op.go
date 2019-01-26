package avltree

import (
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/utils"
)

func (x *Node) inOrder(ll *linkedlist.LinkedList) {
	if x.left != nil {
		x.left.inOrder(ll)
	}

	ll.Append([]interface{}{x.value, x.height, x.bFactor})

	if x.right != nil {
		x.right.inOrder(ll)
	}
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
	x.bFactor = x.recomputeBFactor()
}

func (x *Node) recomputeBFactor() int {
	if x.left == nil && x.right == nil {
		return 0
	} else if x.left == nil {
		return 0 - x.right.height
	} else if x.right == nil {
		return x.left.height
	} else {
		return x.left.height - x.right.height
	}
}
