package redblacktree

import (
	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/stack"
)

func (rbt *RedBlackTree) add(value interface{}) bool {
	if rbt.root == nil {
		rbt.root = newNode(value, nil, black)
		rbt.size++
		return true
	}

	var newElem *Node

	node := rbt.root
	for {
		if comp := rbt.compare(node.value, value); comp == -1 {
			if node.right == nil {
				newElem = newNode(value, node, red)
				node.right = newElem
				break
			}
			node = node.right
		} else if comp == 1 {
			if node.left == nil {
				newElem = newNode(value, node, red)
				node.left = newElem
				break
			}
			node = node.left
		} else {
			return false
		}
	}
	rbt.size++
	rbt.rebalance(newElem)
	return true
}

func (rbt *RedBlackTree) remove(value interface{}) bool {
	if rbt.length() == 0 {
		return false
	}

	rbt.size--
	var parent *Node
	s := stack.New()

	node := rbt.root
	for node != nil {
		s.Push(node)
		if comp := rbt.compare(node.value, value); comp == 1 {
			parent = node
			node = node.left
		} else if comp == -1 {
			parent = node
			node = node.right
		} else {
			if node.left == nil && node.right == nil {
				if parent == nil {
					rbt.root = nil
				} else if rbt.compare(parent.value, value) == 1 {
					parent.left = nil
				} else {
					parent.right = nil
				}
				return true
			}

			if node.left == nil {
				if parent == nil {
					rbt.root = node.right
				} else if rbt.compare(parent.value, value) == 1 {
					parent.left = node.right
				} else {
					parent.right = node.right
				}
				return true
			} else if node.right == nil {
				if parent == nil {
					rbt.root = node.left
				} else if rbt.compare(parent.value, value) == 1 {
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

	rbt.size++
	return false
}

func (rbt *RedBlackTree) height() int {
	if rbt.length() == 0 {
		return 0
	}

	return rbt.root.height
}

func (rbt *RedBlackTree) min() (*Node, bool) {
	if rbt.length() == 0 {
		return nil, false
	}

	node := rbt.root
	for {
		if node.left == nil {
			return node, true
		}
		node = node.left
	}
}

func (rbt *RedBlackTree) max() (*Node, bool) {
	if rbt.length() == 0 {
		return nil, false
	}

	node := rbt.root
	for {
		if node.right == nil {
			return node, true
		}
		node = node.right
	}
}

func (rbt *RedBlackTree) contains(value interface{}) bool {
	if rbt.length() == 0 {
		return false
	}

	node := rbt.root
	for {
		if comp := rbt.compare(node.value, value); comp == -1 {
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

func (rbt *RedBlackTree) length() int {
	return rbt.size
}

func (rbt *RedBlackTree) clear() {
	rbt.root = nil
	rbt.size = 0
}

func (rbt *RedBlackTree) inOrder() list.List {
	ll := linkedlist.New()

	if rbt.length() != 0 {
		rbt.root.inOrder(ll)
	}

	return ll
}

func (rbt *RedBlackTree) preOrder() list.List {
	ll := linkedlist.New()

	if rbt.Length() != 0 {
		rbt.root.preOrder(ll)
	}

	return ll
}

func (rbt *RedBlackTree) postOrder() list.List {
	ll := linkedlist.New()

	if rbt.Length() != 0 {
		rbt.root.postOrder(ll)
	}

	return ll
}

func (rbt *RedBlackTree) getNeigbours(x *Node) (p, g, u *Node) {
	// Parent
	p = x.parent

	// Grandparent
	if p != nil {
		g = p.parent
	}

	// Uncle
	if g != nil {
		u = g.right
		if rbt.compare(p.value, g.value) == 1 {
			u = g.left
		}
	}

	return
}

func (rbt *RedBlackTree) rebalance(x *Node) {
	p, g, u := rbt.getNeigbours(x)
	for {
		if p == nil || g == nil {
			return
		}
		p.recomputeHeight()
		g.recomputeHeight()

		if p.isBlack() {
			return
		}

		if u != nil && u.isRed() { // Has a RED uncle?
			p.toBlack()
			u.toBlack()

			if !g.isRoot() {
				g.toRed()
			}

			x = g
		} else {
			var nr *Node

			// Left-Left Case - Need Right Rotation
			if rbt.compare(g.value, p.value) == 1 && rbt.compare(p.value, x.value) == 1 {
				nr = g.rightRotate()
			}

			// Left-Right Case - Need Left Rotation, then Right Rotation
			if rbt.compare(g.value, p.value) == 1 && rbt.compare(p.value, x.value) == -1 {
				g.left = g.left.leftRotate()
				p = x
				nr = g.rightRotate()
			}

			// Right-Right Case - Need Left Rotation
			if rbt.compare(g.value, p.value) == -1 && rbt.compare(p.value, x.value) == -1 {
				nr = g.leftRotate()
			}

			// Right-Left Case - Need right rotation, then left rotation
			if rbt.compare(g.value, p.value) == -1 && rbt.compare(p.value, x.value) == 1 {
				g.right = g.right.rightRotate()
				p = x
				nr = g.leftRotate()
			}
			g.color, p.color = p.color, g.color

			if nr.parent == nil {
				rbt.root = nr
			} else if rbt.compare(nr.parent.value, nr.value) == 1 {
				nr.parent.left = nr
			} else {
				nr.parent.right = nr
			}

		}
		p, g, u = rbt.getNeigbours(x)
	}
}
