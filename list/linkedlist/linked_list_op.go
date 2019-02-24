package linkedlist

import (
	"fmt"

	"github.com/praveen001/ds/ds"
)

func (ll *LinkedList) len() int {
	return ll.size
}

func (ll *LinkedList) front() (interface{}, bool) {
	return ll.get(0)
}

func (ll *LinkedList) back() (interface{}, bool) {
	return ll.get(ll.len() - 1)
}

func (ll *LinkedList) pushFront(v interface{}) {
	elem := &element{v, ll.head}

	if ll.len() == 0 {
		ll.tail = elem
	}
	ll.head = elem

	ll.size++
}

func (ll *LinkedList) pushBack(v interface{}) {
	elem := &element{v, nil}

	if ll.len() == 0 {
		ll.head = elem
	} else {
		ll.tail.next = elem
	}
	ll.tail = elem
	ll.size++
}

func (ll *LinkedList) set(i int, v interface{}) bool {
	if ll.withInRange(i) {
		elem := ll.getElemByIdx(i)
		elem.value = v
		return true
	}

	return false
}

func (ll *LinkedList) get(index int) (interface{}, bool) {
	if ll.withInRange(index) {
		elem := ll.getElemByIdx(index)
		return elem.value, true
	}

	return nil, false
}

func (ll *LinkedList) insert(i int, v interface{}) bool {
	if i >= 0 && i <= ll.len() {
		if i == 0 {
			ll.pushFront(v)
		} else if i == ll.len() {
			ll.pushBack(v)
		} else {
			var n *element
			idx := 0
			for n = ll.head; idx < i-1; n = n.next {
				idx++
			}
			elem := &element{v, n.next}
			n.next = elem

			ll.size++
		}

		return true
	}

	return false
}

func (ll *LinkedList) remove(i int) (interface{}, bool) {
	if ll.withInRange(i) {
		var value interface{}

		if i == 0 {
			elem := ll.head
			value = elem.value

			ll.head = elem.next
		} else {
			elem := ll.getElemByIdx(i - 1)
			value = elem.next.value

			elem.next = elem.next.next
		}
		ll.size--

		return value, true
	}

	return nil, false
}

func (ll *LinkedList) clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

func (ll *LinkedList) pushBackList(l ds.List) {
	for _, v := range l.Values() {
		ll.pushBack(v)
	}
}

func (ll *LinkedList) pushFrontList(l ds.List) {
	for _, v := range l.Values() {
		ll.pushFront(v)
	}
}

func (ll *LinkedList) contains(value interface{}) bool {
	for elem := ll.head; elem != nil; elem = elem.next {
		if elem.value == value {
			return true
		}
	}

	return false
}

func (ll *LinkedList) indexOf(value interface{}) int {
	elem := ll.head
	for i := 0; i < ll.size; i++ {
		if elem.value == value {
			return i
		}
		elem = elem.next
	}

	return -1
}

func (ll *LinkedList) values() []interface{} {
	arr := make([]interface{}, ll.len())
	index := 0

	for elem := ll.head; elem != nil; elem = elem.next {
		arr[index] = elem.value
		index++
	}

	return arr
}

func (ll *LinkedList) clone() *LinkedList {
	nal := New()
	nal.PushBackList(ll)

	return nal
}

func (ll *LinkedList) swap(a, b int) bool {
	if ll.withInRange(a) && ll.withInRange(b) {
		an := ll.getElemByIdx(a)
		bn := ll.getElemByIdx(b)

		an.value, bn.value = bn.value, an.value
		return true
	}
	return false
}

// getElemByIdx returns the element at the given index
func (ll *LinkedList) getElemByIdx(index int) *element {
	elem := ll.head
	for i := 0; i < index; i++ {
		elem = elem.next
	}

	return elem
}

func (ll *LinkedList) withInRange(index int) bool {
	return index > -1 && index < ll.size
}

func (ll *LinkedList) string() string {
	str := "["
	for elem := ll.head; elem != nil; elem = elem.next {
		str += fmt.Sprintf("%v", elem.value)
		if elem.next != nil {
			str += " "
		}
	}
	str += "]"
	return str
}
