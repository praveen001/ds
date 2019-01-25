package avltree

import (
	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/queue"
)

func (at *AvlTree) insert(value interface{}) {
	at.size++

	if at.root == nil {
		at.root = newNode(value)
		return
	}

	node := at.root
	for {
		if at.compare(node.value, value) == -1 {
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

func (at *AvlTree) delete(value interface{}) bool {
	if at.count() == 0 {
		return false
	}

	at.size--
	node := at.root
	var parent *treeNode

	for node != nil {
		if comp := at.compare(node.value, value); comp == 1 {
			parent = node
			node = node.left
		} else if comp == -1 {
			parent = node
			node = node.right
		} else {

			if node.left == nil && node.right == nil {
				if parent == nil {
					at.root = nil
				} else if comp := at.compare(parent.value, value); comp == 1 {
					parent.left = nil
				} else {
					parent.right = nil
				}
				return true
			}

			if node.left == nil {
				if parent == nil {
					at.root = node.right
				} else if comp := at.compare(parent.value, value); comp == 1 {
					parent.left = node.right
				} else {
					parent.right = node.right
				}
				return true
			} else if node.right == nil {
				if parent == nil {
					at.root = node.left
				} else if comp := at.compare(parent.value, value); comp == 1 {
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

	at.size++
	return false
}

func (at *AvlTree) contains(value interface{}) bool {
	if at.count() == 0 {
		return false
	}

	node := at.root
	for {
		if comp := at.compare(node.value, value); comp == -1 {
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

func (at *AvlTree) height() int {
	if at.count() == 0 {
		return 0
	}

	q := queue.New()
	q.Enqueue(at.root)
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

func (at *AvlTree) min() (interface{}, bool) {
	if at.count() == 0 {
		return nil, false
	}

	node := at.root
	for {
		if node.left == nil {
			return node.value, true
		}
		node = node.left
	}
}

func (at *AvlTree) max() (interface{}, bool) {
	if at.count() == 0 {
		return nil, false
	}

	node := at.root
	for {
		if node.right == nil {
			return node.value, true
		}
		node = node.right
	}
}

func (at *AvlTree) count() int {
	return at.size
}

func (at *AvlTree) empty() {
	at.root = nil
	at.size = 0
}

func (at *AvlTree) inOrder() list.List {
	ll := linkedlist.New()

	if at.count() != 0 {
		at.root.inOrder(ll)
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
