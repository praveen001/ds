package doublylinkedlist

import (
	"fmt"

	"github.com/praveen001/ds/list"
)

func (dl *DoublyLinkedList) len() int {
	return dl.size
}

func (dl *DoublyLinkedList) front() interface{} {
	return dl.get(0)
}

func (dl *DoublyLinkedList) back() interface{} {
	return dl.get(dl.len() - 1)
}

func (dl *DoublyLinkedList) pushFront(v interface{}) {
	elem := &element{v, dl.head, nil}

	if dl.len() == 0 {
		dl.tail = elem
	} else {
		dl.head.prev = elem
	}
	dl.head = elem

	dl.size++
}

func (dl *DoublyLinkedList) pushBack(v interface{}) {
	elem := &element{v, nil, dl.tail}

	if dl.len() == 0 {
		dl.head = elem
	} else {
		dl.tail.next = elem
	}
	dl.tail = elem

	dl.size++
}

func (dl *DoublyLinkedList) insert(i int, v interface{}) bool {
	return dl.set(i, v)
}

func (dl *DoublyLinkedList) remove(i int) (interface{}, bool) {
	if dl.withInRange(i) {
		var value interface{}

		if i == 0 {
			elem := dl.head
			value = elem.value

			if elem.next != nil {
				elem.prev = nil
			}
			dl.head = elem.next
		} else {
			elem := dl.getElemByIdx(i)
			value = elem.value

			if elem.next != nil {
				elem.next.prev = elem.prev
			}
			elem.prev.next = elem.next
		}
		dl.size--

		return value, true
	}

	return nil, false
}

func (dl *DoublyLinkedList) at(i int) (interface{}, bool) {
	if dl.withInRange(i) {
		return dl.get(i), true
	}
	return nil, false
}

func (dl *DoublyLinkedList) clear() {
	dl.head = nil
	dl.tail = nil
	dl.size = 0
}

func (dl *DoublyLinkedList) pushBackList(l list.List) {
	for _, v := range l.Values() {
		dl.pushBack(v)
	}
}

func (dl *DoublyLinkedList) pushFrontList(l list.List) {
	for _, v := range l.Values() {
		dl.pushFront(v)
	}
}

func (dl *DoublyLinkedList) contains(value interface{}) bool {
	for elem := dl.head; elem != nil; elem = elem.next {
		if elem.value == value {
			return true
		}
	}

	return false
}

func (dl *DoublyLinkedList) indexOf(value interface{}) int {
	elem := dl.head
	for i := 0; i < dl.size; i++ {
		if elem.value == value {
			return i
		}
		elem = elem.next
	}

	return -1
}

func (dl *DoublyLinkedList) values() []interface{} {
	arr := make([]interface{}, dl.len())
	index := 0

	for elem := dl.head; elem != nil; elem = elem.next {
		arr[index] = elem.value
		index++
	}

	return arr
}

func (dl *DoublyLinkedList) clone() *DoublyLinkedList {
	nal := New()
	nal.pushFrontList(dl)

	return nal
}

func (dl *DoublyLinkedList) swap(a, b int) bool {
	if dl.withInRange(a) && dl.withInRange(b) {
		an := dl.getElemByIdx(a)
		bn := dl.getElemByIdx(b)

		an.value, bn.value = bn.value, an.value
		return true
	}
	return false
}

func (dl *DoublyLinkedList) set(index int, value interface{}) bool {
	if dl.withInRange(index) {
		elem := dl.getElemByIdx(index)
		elem.value = value
		return true
	}

	return false
}

func (dl *DoublyLinkedList) get(index int) interface{} {
	if dl.withInRange(index) {
		elem := dl.getElemByIdx(index)
		return elem.value
	}

	return nil
}

// getElemByIdx returns the element at the given index
func (dl *DoublyLinkedList) getElemByIdx(index int) *element {
	elem := dl.head
	for i := 0; i < index; i++ {
		elem = elem.next
	}

	return elem
}

func (dl *DoublyLinkedList) withInRange(index int) bool {
	return index > -1 && index < dl.size
}

func (dl *DoublyLinkedList) string() string {
	str := "["
	for elem := dl.head; elem != nil; elem = elem.next {
		str += fmt.Sprintf("%v", elem.value)
		if elem.next != nil {
			str += " "
		}
	}
	str += "]"
	return str
}
