package avltree

import (
	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/stack"

	"github.com/praveen001/ds/list/linkedlist"
)

func (at *AvlTree) set(key, value interface{}) bool {
	if at.length() == 0 {
		at.root = newNode(key, value)
		at.size++
		return true
	}

	s := stack.New()

	node := at.root
	for {
		s.Push(node)
		if comp := at.compare(node.key, key); comp == 1 {
			if !hasLeft(node) {
				node.left = newNode(key, value)
				break
			}
			node = node.left
		} else if comp == -1 {
			if !hasRight(node) {
				node.right = newNode(key, value)
				break
			}
			node = node.right
		} else {
			return false
		}
	}
	at.size++

	at.rebalance(s)
	return true
}

func (at *AvlTree) rset(node *Node, key, value interface{}) *Node {
	if node == nil {
		at.size++
		return newNode(key, value)
	}

	if comp := at.compare(node.key, key); comp == 1 {
		node.left = at.rset(node.left, key, value)
	} else if comp == -1 {
		node.right = at.rset(node.right, key, value)
	} else {
		return nil
	}

	node.recomputeHeight()
	if node.bFactor > 1 {
		if node.left.bFactor < 0 {
			node.left = at.leftRotate(node.left)
		}
		return at.rightRotate(node)
	} else if node.bFactor < -1 {
		if node.right.bFactor > 0 {
			node.right = at.rightRotate(node.right)
		}
		return at.leftRotate(node)
	} else {
		return node
	}
}

func (at *AvlTree) get(key interface{}) (interface{}, bool) {
	node := at.root
	for node != nil {
		if comp := at.compare(node.key, key); comp == -1 {
			node = node.right
		} else if comp == 1 {
			node = node.left
		} else {
			return node.value, true
		}
	}

	return nil, false
}

func (at *AvlTree) remove(key interface{}) bool {
	var parent *Node
	s := stack.New()
	node := at.root
	for node != nil {
		s.Push(node)
		if comp := at.compare(node.key, key); comp == 1 {
			parent = node
			node = node.left
		} else if comp == -1 {
			parent = node
			node = node.right
		} else {
			if !hasLeft(node) || !hasRight(node) {
				if ch := child(node); parent == nil {
					at.root = ch
				} else if at.compare(parent.key, key) == 1 {
					parent.left = ch
				} else {
					parent.right = ch
				}
				at.size--
				at.rebalance(s)
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

func (at *AvlTree) height() int {
	if at.Len() == 0 {
		return 0
	}
	return at.root.height
}

func (at *AvlTree) min() (node *Node, ok bool) {
	for node = at.root; hasLeft(node); node = node.left {
		ok = true
	}
	return
}

func (at *AvlTree) max() (node *Node, ok bool) {
	for node = at.root; hasRight(node); node = node.right {
		ok = true
	}
	return
}

func (at *AvlTree) contains(key interface{}) bool {
	node := at.root
	for node != nil {
		if comp := at.compare(node.key, key); comp == -1 {
			node = node.right
		} else if comp == 1 {
			node = node.left
		} else {
			return true
		}
	}
	return false
}

func (at *AvlTree) length() int {
	return at.size
}

func (at *AvlTree) clear() {
	at.root = nil
	at.size = 0
}

func (at *AvlTree) inOrder() ds.List {
	ll := linkedlist.New()

	if at.length() != 0 {
		at.root.inOrder(ll)
	}

	return ll
}

func (at *AvlTree) preOrder() ds.List {
	ll := linkedlist.New()

	if at.Len() != 0 {
		at.root.preOrder(ll)
	}

	return ll
}

func (at *AvlTree) postOrder() ds.List {
	ll := linkedlist.New()

	if at.Len() != 0 {
		at.root.postOrder(ll)
	}

	return ll
}

func (at *AvlTree) leftRotate(x *Node) *Node {
	y := x.right
	z := y.left

	y.left = x
	x.right = z

	x.recomputeHeight()
	y.recomputeHeight()

	return y
}

func (at *AvlTree) rightRotate(x *Node) *Node {
	y := x.left
	z := y.right

	y.right = x
	x.left = z

	x.recomputeHeight()
	y.recomputeHeight()

	return y
}

func (at *AvlTree) rebalance(s *stack.Stack) {
	var prev, node *Node
	for {
		p, ok := s.Pop()
		if !ok {
			break
		}
		node = p.(*Node)

		if prev != nil {
			if comp := at.compare(node.key, prev.key); comp == 1 {
				node.left = prev
			} else if comp == -1 {
				node.right = prev
			}
			prev = nil
		}
		node.recomputeHeight()

		prev = nil
		if node.bFactor > 1 {
			if node.left.bFactor < 0 {
				node.left = at.leftRotate(node.left)
			}
			prev = at.rightRotate(node)
		} else if node.bFactor < -1 {
			if node.right.bFactor > 0 {
				node.right = at.rightRotate(node.right)
			}
			prev = at.leftRotate(node)
		}
	}
	if prev != nil {
		at.root = prev
	}
}
