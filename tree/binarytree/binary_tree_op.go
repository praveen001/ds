package binarytree

import (
	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/queue"
)

func (bt *BinaryTree) insert(value interface{}) {
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

func (bt *BinaryTree) delete(value interface{}) bool {
	if bt.count() == 0 {
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

func (bt *BinaryTree) contains(value interface{}) bool {
	if bt.count() == 0 {
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

func (bt *BinaryTree) height() int {
	if bt.count() == 0 {
		return 0
	}

	q := queue.New()
	q.Enqueue(bt.root)
	height := 0

	for {
		nodeCount := q.Count()
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

func (bt *BinaryTree) min() (interface{}, bool) {
	if bt.count() == 0 {
		return nil, false
	}

	node := bt.root
	for {
		if node.left == nil {
			return node.value, true
		}
		node = node.left
	}
}

func (bt *BinaryTree) max() (interface{}, bool) {
	if bt.count() == 0 {
		return nil, false
	}

	node := bt.root
	for {
		if node.right == nil {
			return node.value, true
		}
		node = node.right
	}
}

func (bt *BinaryTree) count() int {
	return bt.size
}

func (bt *BinaryTree) empty() {
	bt.root = nil
	bt.size = 0
}

func (bt *BinaryTree) inOrder() list.List {
	ll := linkedlist.New()

	if bt.count() != 0 {
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
