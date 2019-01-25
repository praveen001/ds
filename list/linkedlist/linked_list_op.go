package linkedlist

import (
	"fmt"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/utils"
)

func (ll *LinkedList) append(values ...interface{}) {
	for _, value := range values {
		newElem := &element{value: value}

		if ll.size == 0 {
			ll.head = newElem
		} else {
			ll.tail.next = newElem
		}

		ll.tail = newElem
		ll.size++
	}
}

func (ll *LinkedList) prepend(values ...interface{}) {
	for _, value := range values {
		h := &element{value: value}
		h.next = ll.head
		ll.head = h
	}
	ll.size += len(values)
}

func (ll *LinkedList) get(index int) (interface{}, bool) {
	if ll.withInRange(index) {
		elem := ll.getElemByIdx(index)
		return elem.value, true
	}

	return nil, false
}

func (ll *LinkedList) set(index int, value interface{}) bool {
	if ll.withInRange(index) {
		elem := ll.getElemByIdx(index)
		elem.value = value
		return true
	}

	return false
}

func (ll *LinkedList) remove(index int) (interface{}, bool) {
	if ll.withInRange(index) {
		var value interface{}
		if index == 0 {
			elem := ll.head
			value = elem.value

			if elem.next != nil {
				ll.head = elem.next
			} else {
				ll.head = nil
			}
		} else {
			elem := ll.getElemByIdx(index - 1)
			value = elem.value

			if elem.next != nil {
				elem.next = elem.next.next
			} else {
				elem.next = nil
			}
		}

		ll.size--

		return value, true
	}

	return nil, false
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
	arr := make([]interface{}, ll.length())
	index := 0

	for elem := ll.head; elem != nil; elem = elem.next {
		arr[index] = elem.value
		index++
	}

	return arr
}

func (ll *LinkedList) length() int {
	return ll.size
}

func (ll *LinkedList) clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

func (ll *LinkedList) withInRange(index int) bool {
	return index > -1 && index < ll.size
}

func (ll *LinkedList) string() string {
	str := ""
	for elem := ll.head; elem != nil; elem = elem.next {
		str += fmt.Sprintf("%v ", elem.value)
	}

	return str
}

// getElemByIdx returns the element at the given index
func (ll *LinkedList) getElemByIdx(index int) *element {
	elem := ll.head
	for i := 0; i < index; i++ {
		elem = elem.next
	}

	return elem
}

func (ll *LinkedList) filter(fn utils.FilterFunc) list.List {
	l := New()
	for _, value := range ll.values() {
		if fn(value) {
			l.append(value)
		}
	}

	return l
}