package avltree

import (
	"github.com/praveen001/ds/stack"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/linkedlist"
)

func (at *AvlTree) insert(value interface{}) bool {
	if at.length() == 0 {
		at.root = newNode(value)
		at.size++
		return true
	}

	s := stack.New()

	node := at.root
	for {
		s.Push(node)
		if comp := at.compare(node.value, value); comp == 1 {
			if node.left == nil {
				node.left = newNode(value)
				break
			} else {
				node = node.left
			}
		} else if comp == -1 {
			if node.right == nil {
				node.right = newNode(value)
				break
			} else {
				node = node.right
			}
		} else {
			return false
		}
	}
	at.size++

	at.rebalance(s)
	return true
}

func (at *AvlTree) rinsert(node *Node, value interface{}) *Node {
	if node == nil {
		at.size++
		return newNode(value)
	}

	if comp := at.compare(node.value, value); comp == 1 {
		node.left = at.rinsert(node.left, value)
	} else if comp == -1 {
		node.right = at.rinsert(node.right, value)
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

func (at *AvlTree) delete(value interface{}) bool {
	if at.length() == 0 {
		return false
	}

	at.size--
	var parent *Node
	s := stack.New()

	node := at.root
	for node != nil {
		s.Push(node)
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
				} else if at.compare(parent.value, value) == 1 {
					parent.left = nil
				} else {
					parent.right = nil
				}
				at.rebalance(s)
				return true
			}

			if node.left == nil {
				if parent == nil {
					at.root = node.right
				} else if at.compare(parent.value, value) == 1 {
					parent.left = node.right
				} else {
					parent.right = node.right
				}
				at.rebalance(s)
				return true
			} else if node.right == nil {
				if parent == nil {
					at.root = node.left
				} else if at.compare(parent.value, value) == 1 {
					parent.left = node.left
				} else {
					parent.right = node.left
				}
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
	if at.length() == 0 {
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
	if at.root == nil {
		return 0
	}
	return at.root.height
}

func (at *AvlTree) min() (interface{}, bool) {
	if at.length() == 0 {
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
	if at.length() == 0 {
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

func (at *AvlTree) length() int {
	return at.size
}

func (at *AvlTree) clear() {
	at.root = nil
	at.size = 0
}

func (at *AvlTree) inOrder() list.List {
	ll := linkedlist.New()

	if at.length() != 0 {
		at.root.inOrder(ll)
	}

	return ll
}

func (at *AvlTree) preOrder() list.List {
	ll := linkedlist.New()

	if at.Length() != 0 {
		at.root.preOrder(ll)
	}

	return ll
}

func (at *AvlTree) postOrder() list.List {
	ll := linkedlist.New()

	if at.Length() != 0 {
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
			if comp := at.compare(node.value, prev.value); comp == 1 {
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
