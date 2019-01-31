package redblacktree

import (
	"github.com/praveen001/ds/ds"
	"github.com/praveen001/ds/list/linkedlist"
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
	node := rbt.root
	for node != nil {
		if comp := rbt.compare(node.value, value); comp == 1 {
			node = node.left
		} else if comp == -1 {
			node = node.right
		} else {
			if !node.hasLeft() || !node.hasRight() {
				rbt.deletionRebalance(node)

				if ch := node.child(); node.parent == nil {
					rbt.root = ch
				} else if rbt.compare(node.parent.value, value) == 1 {
					node.parent.left = ch
				} else {
					node.parent.right = ch
				}
				rbt.size--
				return true
			}

			min := node.right
			for min.hasLeft() {
				min = min.left
			}

			node.value = min.value
			value = min.value
			node = node.right
		}
	}

	return false
}

func (rbt *RedBlackTree) min() (node *Node, ok bool) {
	for node = rbt.root; node.hasLeft(); node = node.left {
		ok = true
	}
	return
}

func (rbt *RedBlackTree) max() (node *Node, ok bool) {
	for node = rbt.root; node.hasRight(); node = node.right {
		ok = true
	}
	return
}

func (rbt *RedBlackTree) contains(value interface{}) bool {
	node := rbt.root
	for node != nil {
		if comp := rbt.compare(node.value, value); comp == -1 {
			node = node.right
		} else if comp == 1 {
			node = node.left
		} else {
			return true
		}
	}
	return false
}

func (rbt *RedBlackTree) len() int {
	return rbt.size
}

func (rbt *RedBlackTree) clear() {
	rbt.root = nil
	rbt.size = 0
}

func (rbt *RedBlackTree) inOrder() ds.List {
	ll := linkedlist.New()
	rbt.root.inOrder(ll)
	return ll
}

func (rbt *RedBlackTree) preOrder() ds.List {
	ll := linkedlist.New()
	rbt.root.preOrder(ll)
	return ll
}

func (rbt *RedBlackTree) postOrder() ds.List {
	ll := linkedlist.New()
	rbt.root.postOrder(ll)
	return ll
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
	}

	if y.isRoot() {
		rbt.root = y
	} else if y.parent.left == x {
		y.parent.left = y
	} else {
		y.parent.right = y
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

	if y.isRoot() {
		rbt.root = y
	} else if y.parent.left == x {
		y.parent.left = y
	} else {
		y.parent.right = y
	}
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
		parent, grandParent, uncle := rbt.getNeigbours(x)

		// Case 2 - Parent is black, then we have inserted a red node, so no need to rebalance the tree
		if parent.isBlack() || grandParent == nil {
			return
		}

		// Case 3 - Uncle is red
		if uncle.isRed() {
			parent.toBlack()
			uncle.toBlack()

			if !grandParent.isRoot() {
				grandParent.toRed()
			}

			x = grandParent
			continue
		}

		// p = x
		// var nr *Node
		// Left-Left Case - Need Right Rotation
		if parent.isLeft() && x.isLeft() {
			rbt.rightRotate(grandParent)
			grandParent.color, parent.color = parent.color, grandParent.color
			return
		}

		// Left-Right Case - Need Left Rotation, then Right Rotation
		if parent.isLeft() && x.isRight() {
			rbt.leftRotate(grandParent.left)
			rbt.rightRotate(grandParent)
			parent = x
			grandParent.color, parent.color = parent.color, grandParent.color
			return
		}

		// Right-Right Case - Need Left Rotation
		if parent.isRight() && x.isRight() {
			rbt.leftRotate(grandParent)
			grandParent.color, parent.color = parent.color, grandParent.color
			return
		}

		// Right-Left Case - Need right rotation, then left rotation
		if parent.isRight() && x.isLeft() {
			rbt.rightRotate(grandParent.right)
			rbt.leftRotate(grandParent)
			parent = x
			grandParent.color, parent.color = parent.color, grandParent.color
			return
		}

	}
}

func (rbt *RedBlackTree) deletionRebalance(n *Node) {
	// No need to rebalance if deleted node is red
	// Because when we delete a RED node it doesn't affect the count of black nodes in any path
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

	// Double back fix
	rbt.deleteCase1(n)
}

func (rbt *RedBlackTree) deleteCase1(n *Node) {
	if !n.isRoot() {
		rbt.deleteCase2(n)
	}
}

func (rbt *RedBlackTree) deleteCase2(n *Node) {
	parent := n.parent
	sib := n.sibling()
	sibx := sib.left
	siby := sib.right

	// Case 2:
	if parent.isBlack() && sib.isRed() && sibx.isBlack() && siby.isBlack() {
		parent.toRed()
		sib.toBlack()

		if n.isLeft() {
			rbt.leftRotate(parent)
		} else {
			rbt.rightRotate(parent)
		}
	}

	rbt.deleteCase3(n)
}

func (rbt *RedBlackTree) deleteCase3(n *Node) {
	parent := n.parent
	sib := n.sibling()
	sibx := sib.left
	siby := sib.right

	// Case 3:
	if parent.isBlack() && sib.isBlack() && sibx.isBlack() && siby.isBlack() {
		sib.toRed()

		n = parent
		// continue
		rbt.deleteCase1(n)
	} else {
		rbt.deleteCase4(n)
	}
}

func (rbt *RedBlackTree) deleteCase4(n *Node) {
	parent := n.parent
	sib := n.sibling()
	sibx := sib.left
	siby := sib.right

	// Case 4:
	if parent.isRed() && sib.isBlack() && sibx.isBlack() && siby.isBlack() {
		parent.toBlack()
		sib.toRed()
		return // This is terminal case
	}

	rbt.deleteCase5(n)
}

func (rbt *RedBlackTree) deleteCase5(n *Node) {
	parent := n.parent
	sib := n.sibling()
	sibx := sib.left
	siby := sib.right

	// Case 5: (has mirror)
	if parent.isBlack() && n.isLeft() && sib.isBlack() && sibx.isRed() && siby.isBlack() {
		sib.toRed()
		sibx.toBlack()
		rbt.rightRotate(sib)
	} else if parent.isBlack() && n.isRight() && sib.isBlack() && sibx.isBlack() && siby.isRed() {
		sib.toRed()
		siby.toBlack()
		rbt.leftRotate(sib)
	}

	rbt.deleteCase6(n)
}

func (rbt *RedBlackTree) deleteCase6(n *Node) {
	parent := n.parent
	sib := n.sibling()
	sibx := sib.left
	siby := sib.right

	// Case 6: (has mirror)
	// Sibliing should be black, and it should have atleast 1 red child
	if sib.isBlack() {

		// Rotate
		if siby.isRed() {
			rbt.leftRotate(parent)
			siby.toBlack()
		} else if sibx.isRed() {
			rbt.rightRotate(parent)
			sibx.toBlack()
		}

		sib.color = parent.color
		parent.toBlack()

		// Terminal Case
	}
}
