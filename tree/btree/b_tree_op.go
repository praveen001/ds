package btree

import (
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/stack"

	"github.com/praveen001/ds/ds"
)

func (bt *BTree) set(key, value interface{}) bool {
	e := &entry{key, value}

	if bt.root == nil {
		nn := newNode()
		nn.entries = append(nn.entries, e)

		bt.root = nn
		bt.size++

		return true
	}

	inserted := bt.insert(bt.root, e)
	if inserted {
		bt.size++
	}
	return inserted
}

func (bt *BTree) get(key interface{}) (interface{}, bool) {
	n := bt.root
	for n != nil {
		foundChild := false
		childIdx := 0
		for _, e := range n.entries {
			comp := bt.compare(e.key, key)

			if comp == 0 {
				return e.value, true
			}

			if comp == 1 {
				foundChild = true
				break
			}

			childIdx++
		}

		if n.isLeaf() {
			break
		}

		if foundChild {
			n = n.children[childIdx]
		} else {
			n = n.children[len(n.children)-1]
		}
	}

	return nil, false
}

func (bt *BTree) remove(key interface{}) bool {
	s := stack.New()
	n := bt.root

	exists := false
	idx := 0
NodeSearch:
	for n != nil {
		s.Push(n)

		for i, e := range n.entries {
			if comp := bt.compare(e.key, key); comp == 1 {
				n = n.children[i]
				continue NodeSearch
			} else if comp == 0 {
				exists = true
				idx = i
				if n.isLeaf() {
					break NodeSearch
				}
				n.entries[i] = n.children[i+1].entries[0]
				key = n.entries[i].key
				n = n.children[i+1]
				continue NodeSearch
			}
		}
		if n.isLeaf() {
			return false
		}
		n = n.children[len(n.children)-1]
	}

	if !exists {
		return false
	}

	n.deleteEntry(idx)
	bt.size--

	if n.isLeaf() {
		bt.deletionRebalance(s)
	}

	return true
}

func (bt *BTree) min() (*entry, bool) {
	if bt.size == 0 {
		return nil, false
	}

	var node *Node
	for node = bt.root; !node.isLeaf(); node = node.children[0] {
	}
	return node.entries[0], true
}

func (bt *BTree) max() (*entry, bool) {
	if bt.size == 0 {
		return nil, false
	}

	var node *Node
	for node = bt.root; !node.isLeaf(); node = node.children[len(node.children)-1] {
	}
	return node.entries[len(node.entries)-1], true
}

func (bt *BTree) contains(key interface{}) bool {
	_, ok := bt.get(key)
	return ok
}

func (bt *BTree) len() int {
	return bt.size
}

func (bt *BTree) clear() {
	bt.root = nil
	bt.size = 0
}

func (bt *BTree) keys() ds.List {
	ll := linkedlist.New()
	bt.root.inOrderKeys(ll)
	return ll
}

func (bt *BTree) values() ds.List {
	ll := linkedlist.New()
	bt.root.inOrder(ll)
	return ll
}

func (bt *BTree) inOrder() ds.List {
	ll := linkedlist.New()
	bt.root.inOrder(ll)
	return ll
}

func (bt *BTree) preOrder() ds.List {
	ll := linkedlist.New()
	bt.root.preOrder(ll)
	return ll
}

func (bt *BTree) postOrder() ds.List {
	ll := linkedlist.New()
	bt.root.postOrder(ll)
	return ll
}

func (bt *BTree) insert(n *Node, e *entry) bool {
	s := stack.New()

NodeSearch:
	for !n.isLeaf() {
		s.Push(n)

		for i, en := range n.entries {
			if comp := bt.compare(en.key, e.key); comp == 1 {
				n = n.children[i]
				continue NodeSearch
			} else if comp == 0 {
				return false
			}
		}
		n = n.children[len(n.entries)]
	}
	s.Push(n)

	pos, exists := bt.getInsertPos(n, e)
	if exists {
		n.entries[pos] = e
		return false
	}

	n.insertEntryAt(pos, e)

	if len(n.entries) == bt.order {
		bt.rebalance(s)
	}

	return true
}

func (bt *BTree) rebalance(s ds.Stack) {
	var nr *Node

	for {
		v, ok := s.Pop()
		if !ok {
			break
		}

		n := v.(*Node)
		if nr != nil {
			pos, _ := bt.getInsertPos(n, nr.entries[0])
			n.insertEntryAt(pos, nr.entries[0])

			n.children[pos] = nr.children[0]
			n.children = append(n.children, nil)
			copy(n.children[pos+2:], n.children[pos+1:])
			n.children[pos+1] = nr.children[1]

			nr = nil
		}

		if len(n.entries) < bt.order {
			break
		}

		medIdx := (bt.order) / 2

		left := newNode()
		left.entries = make([]*entry, medIdx)
		copy(left.entries, n.entries[:medIdx])

		right := newNode()
		right.entries = make([]*entry, len(n.entries[medIdx+1:]))
		copy(right.entries, n.entries[medIdx+1:])

		if !n.isLeaf() {
			left.children = n.children[:medIdx+1]
			right.children = n.children[medIdx+1:]
		}

		nr = newNode()
		nr.entries = append(nr.entries, n.entries[medIdx])
		nr.children = append(nr.children, left, right)
	}

	if nr != nil {
		bt.root = nr
	}
}

func (bt *BTree) deletionRebalance(s ds.Stack) {
	var prev *Node
	for {
		v, ok := s.Pop()
		if !ok {
			break
		}

		n := v.(*Node)
		if len(n.entries) >= bt.order/2 {
			break
		}

		v, ok = s.Peek()
		if !ok {
			// If empty and is root, use prev node as new root
			if len(n.entries) == 0 {
				bt.root = prev
			}
			return
		}

		if ok {
			parent := v.(*Node)

			// Left sibling
			left, leftChildIdx := n.leftSibling(parent)
			if left != nil && len(left.entries) > bt.order/2 {
				// Move key down from parent
				n.entries = append([]*entry{parent.entries[leftChildIdx]}, n.entries...)

				// Move key from left sibling to parent
				parent.entries[leftChildIdx] = left.entries[len(left.entries)-1]

				// Delete from left sibling
				left.deleteEntry(len(left.entries) - 1)

				// Have to move a child?
				if !left.isLeaf() {
					n.children = append([]*Node{left.children[len(left.children)-1]}, n.children...)
					left.deleteChild(len(left.children) - 1)
				}

				return
			}

			right, rightChildIdx := n.rightSibling(parent)
			if right != nil && len(right.entries) > bt.order/2 {
				// Move key down from parent
				n.entries = append(n.entries, parent.entries[rightChildIdx-1])

				// Move key from right sibling to parent
				parent.entries[rightChildIdx-1] = right.entries[0]

				// Delete from right sibling
				right.deleteEntry(0)

				// Have to move a child
				if !right.isLeaf() {
					n.children = append(n.children, right.children[0])
					right.deleteChild(0)
				}

				return
			}

			if left != nil {
				// Append separator key
				left.entries = append(left.entries, parent.entries[leftChildIdx])

				// Append all keys from current node to left sibling
				left.entries = append(left.entries, n.entries...)

				// Append all children from current to node to left sibling
				left.children = append(left.children, n.children...)

				// Delete separator from parent
				parent.deleteEntry(leftChildIdx)

				// Delete child, cause parent lost a node
				parent.deleteChild(leftChildIdx + 1)
			} else if right != nil {
				// Append separator key
				n.entries = append(n.entries, parent.entries[rightChildIdx-1])

				// Append all keys from right sibling to current node
				n.entries = append(n.entries, right.entries...)

				// Append all children from right sibling to current node
				n.children = append(n.children, parent.children[rightChildIdx].children...)

				// Delete separator from parent
				parent.deleteEntry(rightChildIdx - 1)

				// Delete child, cause parent lost a node
				parent.deleteChild(rightChildIdx)
			}
		}

		prev = n
	}
}

func (bt *BTree) getInsertPos(n *Node, e *entry) (int, bool) {
	for i, v := range n.entries {
		if comp := bt.compare(v.key, e.key); comp == 1 {
			return i, false
		} else if comp == 0 {
			return i, true
		}
	}

	return len(n.entries), false
}
