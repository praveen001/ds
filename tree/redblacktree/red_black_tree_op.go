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
				rbt.deletionRebalance(node)
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
				rbt.deletionRebalance(node)
				if parent == nil {
					rbt.root = node.right
				} else if rbt.compare(parent.value, value) == 1 {
					parent.left = node.right
				} else {
					parent.right = node.right
				}
				return true
			} else if node.right == nil {
				rbt.deletionRebalance(node)
				if parent == nil {
					rbt.root = node.left
				} else if rbt.compare(parent.value, value) == 1 {
					parent.left = node.left
				} else {
					parent.right = node.left
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
	for {
		p, g, u := rbt.getNeigbours(x)

		if p == nil {
			return
		}
		p.recomputeHeight()

		if g == nil {
			return
		}
		g.recomputeHeight()

		// Case 2
		if isBlack(p) {
			return
		}

		// Case 3
		if u != nil && isRed(u) { // Has a RED uncle?
			toBlack(p)
			toBlack(u)

			if !isRoot(g) {
				toRed(g)
			}

			x = g
			continue
		}

		// p = x
		// var nr *Node
		// Left-Left Case - Need Right Rotation
		if g.left == p && p.left == x {
			rbt.rightRotate(g)
			g.color, p.color = p.color, g.color
			return
		}

		// Left-Right Case - Need Left Rotation, then Right Rotation
		if g.left == p && p.right == x {
			rbt.leftRotate(g.left)
			rbt.rightRotate(g)
			p = x
			g.color, p.color = p.color, g.color
			return
		}

		// Right-Right Case - Need Left Rotation
		if g.right == p && p.right == x {
			rbt.leftRotate(g)
			g.color, p.color = p.color, g.color
			return
		}

		// Right-Left Case - Need right rotation, then left rotation
		if g.right == p && p.left == x {
			rbt.rightRotate(g.right)
			rbt.leftRotate(g)
			p = x
			g.color, p.color = p.color, g.color
			return
		}
		// g.color, p.color = p.color, g.color

	}
}

func (rbt *RedBlackTree) deletionRebalance(n *Node) {
	// No need to rebalance if deleted node is red
	// Because when we delete a RED node it doesn't affect the count of black nodes in any path
	if isRed(n) {
		return
	}

	// If it has one red child, just change the color
	if !hasRight(n) && hasLeft(n) && isRed(n.left) {
		toBlack(n.left)
		return
	}
	if !hasLeft(n) && hasRight(n) && isRed(n.right) {
		toBlack(n.right)
		return
	}

	// Double back fix
	rbt.deleteCase1(n)
}

func (rbt *RedBlackTree) deleteCase1(n *Node) {
	if !isRoot(n) {
		rbt.deleteCase2(n)
	}
}

func (rbt *RedBlackTree) deleteCase2(n *Node) {
	p := n.parent
	s := sibling(n)
	sx := s.left
	sy := s.right

	// Case 2:
	if isBlack(p) && isRed(s) && isBlack(sx) && isBlack(sy) {
		toRed(p)
		toBlack(s)

		if isLeft(n) {
			rbt.leftRotate(p)
		} else {
			rbt.rightRotate(p)
		}
	}

	rbt.deleteCase3(n)
}

func (rbt *RedBlackTree) deleteCase3(n *Node) {
	p := n.parent
	s := sibling(n)
	sx := s.left
	sy := s.right

	// Case 3:
	if isBlack(p) && isBlack(s) && isBlack(sx) && isBlack(sy) {
		toRed(s)

		n = p
		// continue
	} else {
		rbt.deleteCase4(n)
	}
}

func (rbt *RedBlackTree) deleteCase4(n *Node) {
	p := n.parent
	s := sibling(n)
	sx := s.left
	sy := s.right

	// Case 4:
	if isRed(p) && isBlack(s) && isBlack(sx) && isBlack(sy) {
		toBlack(p)
		toRed(s)
		return // This is terminal case
	}

	rbt.deleteCase5(n)
}

func (rbt *RedBlackTree) deleteCase5(n *Node) {
	p := n.parent
	s := sibling(n)
	sx := s.left
	sy := s.right

	// Case 5: (has mirror)
	if isBlack(p) && isLeft(n) && isBlack(s) && isRed(sx) && isBlack(sy) {
		toRed(s)
		toBlack(sx)
		rbt.rightRotate(s)
	} else if isBlack(p) && isRight(n) && isBlack(s) && isBlack(sx) && isRed(sy) {
		toRed(s)
		toBlack(sy)
		rbt.leftRotate(s)
	}

	rbt.deleteCase6(n)
}

func (rbt *RedBlackTree) deleteCase6(n *Node) {
	p := n.parent
	s := sibling(n)
	sx := s.left
	sy := s.right

	// Case 6: (has mirror)
	// Sibliing should be black, and it should have atleast 1 red child
	if isBlack(s) {

		// Rotate
		if isRed(sy) {
			rbt.leftRotate(p)
			toBlack(sy)
		} else if isRed(sx) {
			rbt.rightRotate(p)
			toBlack(sx)
		}

		s.color = p.color
		toBlack(p)

		// Terminal Case
	}
}

func (rbt *RedBlackTree) rightRotate(x *Node) {
	y := x.left
	z := y.right

	y.right = x
	y.parent = x.parent

	x.left = z
	x.parent = y

	if z != nil {
		z.parent = x
	}

	x.recomputeHeight()
	y.recomputeHeight()

	if y.parent == nil {
		rbt.root = y
	} else if y.parent.left == x {
		y.parent.left = y
	} else {
		y.parent.right = y
	}
}

func (rbt *RedBlackTree) leftRotate(x *Node) {
	y := x.right
	z := y.left

	y.left = x
	y.parent = x.parent

	x.right = z
	x.parent = y

	if z != nil {
		z.parent = x
		z.recomputeHeight()
	}

	x.recomputeHeight()
	y.recomputeHeight()

	if y.parent == nil {
		rbt.root = y
	} else if y.parent.left == x {
		y.parent.left = y
	} else {
		y.parent.right = y
	}
}
