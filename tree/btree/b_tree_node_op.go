package btree

import (
	"github.com/praveen001/ds/list/linkedlist"
)

// Key of the node
func (e *entry) Key() interface{} {
	return e.key
}

// Value of the node
func (e *entry) Value() interface{} {
	return e.value
}

func (n *Node) isLeaf() bool {
	if len(n.children) == 0 {
		return true
	}
	return false
}

func (n *Node) insertEntryAt(i int, e *entry) {
	n.entries = append(n.entries, nil)
	copy(n.entries[i+1:], n.entries[i:])
	n.entries[i] = e
}

func (n *Node) inOrderKeys(ll *linkedlist.LinkedList) {
	if n == nil {
		return
	}

	for i, e := range n.entries {
		if !n.isLeaf() {
			n.children[i].inOrderKeys(ll)
		}
		ll.PushBack(e.key)
	}
	if !n.isLeaf() {
		n.children[len(n.children)-1].inOrderKeys(ll)
	}
}

func (n *Node) inOrder(ll *linkedlist.LinkedList) {
	if n == nil {
		return
	}

	for i, e := range n.entries {
		if !n.isLeaf() {
			n.children[i].inOrder(ll)
		}
		ll.PushBack(e.value)
	}
	if !n.isLeaf() {
		n.children[len(n.children)-1].inOrder(ll)
	}
}

func (n *Node) preOrder(ll *linkedlist.LinkedList) {
	if n == nil {
		return
	}

	for i, e := range n.entries {
		ll.PushBack(e.value)
		if !n.isLeaf() {
			n.children[i].preOrder(ll)
		}
	}
	if !n.isLeaf() {
		n.children[len(n.children)-1].preOrder(ll)
	}
}

func (n *Node) postOrder(ll *linkedlist.LinkedList) {
	if n == nil {
		return
	}

	for i := range n.entries {
		if !n.isLeaf() {
			n.children[i].postOrder(ll)
		}
	}
	if !n.isLeaf() {
		n.children[len(n.children)-1].postOrder(ll)
	}

	for _, e := range n.entries {
		ll.PushBack(e.value)
	}
}
