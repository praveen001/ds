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
		left.entries = n.entries[:medIdx]

		right := newNode()
		right.entries = n.entries[medIdx+1:]

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
