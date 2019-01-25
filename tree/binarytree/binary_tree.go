package binarytree

import (
	"github.com/praveen001/ds/list"

	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/queue"
	"github.com/praveen001/ds/utils"
)

// BinaryTree represents a binary tree
type BinaryTree struct {
	root    *treeNode
	size    int
	compare utils.Comparator
}

// Node represents a node in a binary tree
type treeNode struct {
	value interface{}
	left  *treeNode
	right *treeNode
}

// New creates a new instance of binary tree and returns it
func New(c utils.Comparator) *BinaryTree {
	return &BinaryTree{
		compare: c,
	}
}

// NewNode returns a new binary tree node with given value
func newNode(value interface{}) *treeNode {
	return &treeNode{value, nil, nil}
}

// Insert given values into the tree
func (bt *BinaryTree) Insert(value interface{}) {
	bt.size++

	if bt.root == nil {
		bt.root = newNode(value)
		return
	}

	node := bt.root
	for {
		if bt.compare(node.value, value) == -1 {
			if node.right == nil {
				node.right = newNode(value)
				break
			} else {
				node = node.right
			}
		} else {
			if node.left == nil {
				node.left = newNode(value)
				break
			} else {
				node = node.left
			}
		}
	}
}

// Delete a value from the tree
func (bt *BinaryTree) Delete(value interface{}) bool {
	if bt.Size() == 0 {
		return false
	}

	bt.size--
	node := bt.root
	var parent *treeNode

	for node != nil {
		if comp := bt.compare(node.value, value); comp == 1 {
			parent = node
			node = node.left
		} else if comp == -1 {
			parent = node
			node = node.right
		} else {

			if node.left == nil && node.right == nil {
				if parent == nil {
					bt.root = nil
				} else if comp := bt.compare(parent.value, value); comp == 1 {
					parent.left = nil
				} else {
					parent.right = nil
				}
				return true
			}

			if node.left == nil {
				if parent == nil {
					bt.root = node.right
				} else if comp := bt.compare(parent.value, value); comp == 1 {
					parent.left = node.right
				} else {
					parent.right = node.right
				}
				return true
			} else if node.right == nil {
				if parent == nil {
					bt.root = node.left
				} else if comp := bt.compare(parent.value, value); comp == 1 {
					parent.left = node.left
				} else {
					parent.right = node.right
				}
				return true
			}

			min := node.right
			for {
				if min.left == nil {
					break
				}
				min = min.left
			}

			node.value = min.value
			value = min.value
			parent = node
			node = node.right

		}
	}

	bt.size++
	return false
}

// Contains returns true if the given value exists in the tree, otherwise false
func (bt *BinaryTree) Contains(value interface{}) bool {
	if bt.Size() == 0 {
		return false
	}

	node := bt.root
	for {
		if comp := bt.compare(node.value, value); comp == -1 {
			if node.right == nil {
				return false
			}
			node = node.right
		} else if comp == 1 {
			if node.left == nil {
				return false
			}
			node = node.left
		} else {
			return true
		}
	}
}

// Height returns the height of the binary tree (using node/level counts)
func (bt *BinaryTree) Height() int {
	if bt.Size() == 0 {
		return 0
	}

	q := queue.New()
	q.Enqueue(bt.root)
	height := 0

	for {
		nodeCount := q.Size()
		if nodeCount == 0 {
			break
		}

		height++
		for i := 0; i < nodeCount; i++ {
			n, _ := q.Dequeue()
			node := n.(*treeNode)

			if node.left != nil {
				q.Enqueue(node.left)
			}
			if node.right != nil {
				q.Enqueue(node.right)
			}
		}
	}

	return height
}

// Min returns the minimum value from the tree
func (bt *BinaryTree) Min() interface{} {
	if bt.Size() == 0 {
		return -1
	}

	node := bt.root
	for {
		if node.left == nil {
			return node.value
		}
		node = node.left
	}
}

// Max returns the maximum value from the tree
func (bt *BinaryTree) Max() interface{} {
	if bt.Size() == 0 {
		return -1
	}

	node := bt.root
	for {
		if node.right == nil {
			return node.value
		}
		node = node.right
	}
}

// Size returns the total number of values in the tree
func (bt *BinaryTree) Size() int {
	return bt.size
}

// Empty clears all the values in the tree
func (bt *BinaryTree) Empty() {
	bt.root = nil
	bt.size = 0
}

// InOrder ..
func (bt *BinaryTree) InOrder() list.List {
	ll := linkedlist.New()

	if bt.Size() != 0 {
		bt.root.inOrder(ll)
	}

	return ll
}

func (n *treeNode) inOrder(ll *linkedlist.LinkedList) {
	if n.left != nil {
		n.left.inOrder(ll)
	}

	ll.Append(n.value)

	if n.right != nil {
		n.right.inOrder(ll)
	}
}
