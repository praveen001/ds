package binarytree

import (
	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/stack"
)

func (bt *BinaryTree) set(key, value interface{}) bool {
	if bt.root == nil {
		bt.root = newNode(key, value)
		bt.size++
		return true
	}

	s := stack.New()

	node := bt.root
	for {
		s.Push(node)
		if comp := bt.compare(node.key, key); comp == -1 {
			if node.right == nil {
				node.right = newNode(key, value)
				break
			}
			node = node.right
		} else if comp == 1 {
			if node.left == nil {
				node.left = newNode(key, value)
				break
			}
			node = node.left
		} else {
			return false
		}
	}
	bt.size++

	bt.readjust(s)
	return true
}

func (bt *BinaryTree) get(key interface{}) (interface{}, bool) {
	node := bt.root
	for node != nil {
		if comp := bt.compare(node.key, key); comp == -1 {
			node = node.right
		} else if comp == 1 {
			node = node.left
		} else {
			return node.value, true
		}
	}

	return nil, false
}

func (bt *BinaryTree) remove(key interface{}) bool {
	var parent *Node
	s := stack.New()
	node := bt.root
	for node != nil {
		s.Push(node)
		if comp := bt.compare(node.key, key); comp == 1 {
			parent = node
			node = node.left
		} else if comp == -1 {
			parent = node
			node = node.right
		} else {
			if !hasLeft(node) || !hasRight(node) {
				if ch := child(node); parent == nil {
					bt.root = ch
				} else if bt.compare(parent.key, key) == 1 {
					parent.left = ch
				} else {
					parent.right = ch
				}
				bt.size--
				return true
			}

			min := node.right
			for {
				if min.left == nil {
					break
				}
				min = min.left
			}

			node.key, node.value = min.key, min.value
			key = min.key
			parent = node
			node = node.right
		}
	}

	return false
}

func (bt *BinaryTree) height() int {
	if bt.length() == 0 {
		return 0
	}

	return bt.root.height
	// q := queue.New()
	// q.Enqueue(bt.root)
	// height := 0

	// for {
	// 	nodeCount := q.Length()
	// 	if nodeCount == 0 {
	// 		break
	// 	}

	// 	height++
	// 	for i := 0; i < nodeCount; i++ {
	// 		n, _ := q.Dequeue()
	// 		node := n.(*Node)

	// 		if node.left != nil {
	// 			q.Enqueue(node.left)
	// 		}
	// 		if node.right != nil {
	// 			q.Enqueue(node.right)
	// 		}
	// 	}
	// }

	// return height
}

func (bt *BinaryTree) min() (*Node, bool) {
	if bt.length() == 0 {
		return nil, false
	}

	node := bt.root
	for {
		if node.left == nil {
			return node, true
		}
		node = node.left
	}
}

func (bt *BinaryTree) max() (*Node, bool) {
	if bt.length() == 0 {
		return nil, false
	}

	node := bt.root
	for {
		if node.right == nil {
			return node, true
		}
		node = node.right
	}
}

func (bt *BinaryTree) contains(key interface{}) bool {
	node := bt.root
	for node != nil {
		if comp := bt.compare(node.key, key); comp == -1 {
			node = node.right
		} else if comp == 1 {
			node = node.left
		} else {
			return true
		}
	}
	return false
}

func (bt *BinaryTree) length() int {
	return bt.size
}

func (bt *BinaryTree) clear() {
	bt.root = nil
	bt.size = 0
}

func (bt *BinaryTree) inOrder() list.List {
	ll := linkedlist.New()

	if bt.length() != 0 {
		bt.root.inOrder(ll)
	}

	return ll
}

func (bt *BinaryTree) preOrder() list.List {
	ll := linkedlist.New()

	if bt.Length() != 0 {
		bt.root.preOrder(ll)
	}

	return ll
}

func (bt *BinaryTree) postOrder() list.List {
	ll := linkedlist.New()

	if bt.Length() != 0 {
		bt.root.postOrder(ll)
	}

	return ll
}

func (bt *BinaryTree) readjust(s *stack.Stack) {
	var node *Node
	for {
		p, ok := s.Pop()
		if !ok {
			break
		}
		node = p.(*Node)
		node.recomputeHeight()
	}
}
