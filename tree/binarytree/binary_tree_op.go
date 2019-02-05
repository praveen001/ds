package binarytree

import (
	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/list/linkedlist"
)

func (bt *BinaryTree) set(key, value interface{}) bool {
	if bt.root == nil {
		bt.root = newNode(key, value)
		bt.size++
		return true
	}

	node := bt.root
	for {
		if comp := bt.compare(node.key, key); comp == -1 {
			if !node.hasRight() {
				node.right = newNode(key, value)
				break
			}
			node = node.right
		} else if comp == 1 {
			if !node.hasLeft() {
				node.left = newNode(key, value)
				break
			}
			node = node.left
		} else {
			return false
		}
	}
	bt.size++

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
	node := bt.root
	for node != nil {
		if comp := bt.compare(node.key, key); comp == 1 {
			parent = node
			node = node.left
		} else if comp == -1 {
			parent = node
			node = node.right
		} else {
			if !node.hasLeft() || !node.hasRight() {
				if ch := node.child(); parent == nil {
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
			for min.hasLeft() {
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

func (bt *BinaryTree) min() (node *Node, ok bool) {
	for node = bt.root; node.hasLeft(); node = node.left {
		ok = true
	}
	return
}

func (bt *BinaryTree) max() (node *Node, ok bool) {
	for node = bt.root; node.hasRight(); node = node.right {
		ok = true
	}
	return
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

func (bt *BinaryTree) len() int {
	return bt.size
}

func (bt *BinaryTree) clear() {
	bt.root = nil
	bt.size = 0
}

func (bt *BinaryTree) keys() ds.List {
	ll := linkedlist.New()
	bt.root.inOrderKeys(ll)
	return ll
}

func (bt *BinaryTree) values() ds.List {
	ll := linkedlist.New()
	bt.root.inOrder(ll)
	return ll
}

func (bt *BinaryTree) inOrder() ds.List {
	ll := linkedlist.New()
	bt.root.inOrder(ll)
	return ll
}

func (bt *BinaryTree) preOrder() ds.List {
	ll := linkedlist.New()
	bt.root.preOrder(ll)
	return ll
}

func (bt *BinaryTree) postOrder() ds.List {
	ll := linkedlist.New()
	bt.root.postOrder(ll)
	return ll
}
