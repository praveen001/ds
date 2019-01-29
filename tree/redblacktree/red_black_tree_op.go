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
				rbt.deletionRebalance(node)
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
				rbt.deletionRebalance(node)
				return true
			} else if node.right == nil {
				if parent == nil {
					rbt.root = node.left
				} else if rbt.compare(parent.value, value) == 1 {
					parent.left = node.left
				} else {
					parent.right = node.left
				}
				rbt.deletionRebalance(node)
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
	for {
		p, g, u := rbt.getNeigbours(x)

		if p == nil || g == nil {
			return
		}
		p.recomputeHeight()
		g.recomputeHeight()

		// Case 2
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
			continue
		}

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
}

func (rbt *RedBlackTree) deletionRebalance(n *Node) {
	// No need to rebalance if deleted node is red
	if n.isRed() {
		return
	}

	// If it has one red child, just change the color
	if !n.hasRight() && n.hasLeft() && n.left.isRed() {
		n.left.toBlack()
		return
	}
	if !n.hasLeft() && n.hasRight() && n.right.isRed() {
		n.right.toBlack()
		return
	}

	for {
		// Case 1
		if n.isRoot() {
			n.toBlack() // To be safe
			return      // Terminal case
		}

		p := n.parent
		s := n.left
		if p.left == n {
			s = p.right
		}
		sx := s.left
		sy := s.right

		// Case 2:
		if p.isBlack() && s.isRed() && sx.isBlack() && sy.isBlack() {
			p.toRed()
			s.toBlack()
			if s.isLeft() {
				s.parent.left = p.rightRotate()
			} else {
				s.parent.right = p.leftRotate()
			}
		}

		// Case 3:
		if p.isBlack() && s.isBlack() && sx.isBlack() && sy.isBlack() {
			s.toRed()

			n = p
			continue
		}

		// Case 4: (no mirror)
		if p.isRed() && s.isBlack() && sx.isBlack() && sy.isBlack() {
			p.toBlack()
			s.toRed()
			return // This is terminal case
		}

		// Case 5: (has mirror)
		if p.isBlack() && s.isBlack() && sx.isRed() && sx.isBlack() {
			nr := s.rightRotate()
			s.parent.right = nr
			sx.color, s.color = s.color, sx.color
		}
		// if n.isLeft() && s.isBlack() && sx.isRed() && sy.isBlack() {
		// 	s.toRed()
		// 	sx.toBlack()
		// 	p.rightRotate()
		// }

		// Case 6: (has mirror)
		if s.isBlack() {
			var nr *Node

			// Rotate
			if sy.isRed() {
				nr = s.leftRotate()
			} else if sx.isRed() {
				nr = s.rightRotate()
			}
			if s.parent.left == s {
				s.parent.left = nr
			} else {
				s.parent.right = nr
			}

			// Recolouring
			s.toBlack()
			p.toBlack() // Safety
			if sy.isRed() {
				sy.toBlack()
			} else if sx.isRed() {
				sx.toBlack()
			}

			return // Terminal case
		}
	}
}
