package doublylinkedlist

import (
	"fmt"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/utils"
)

func (dl *DoublyLinkedList) append(values ...interface{}) {
	for _, value := range values {
		newElem := &element{value: value}

		if dl.length() == 0 {
			dl.head = newElem
		} else {
			dl.tail.next = newElem
			newElem.prev = dl.tail
		}

		dl.tail = newElem
		dl.size++
	}
}

func (dl *DoublyLinkedList) prepend(values ...interface{}) {
	n := dl.head
	for i := 0; i < len(values); i++ {
		h := &element{value: values[len(values)-1-i]}
		h.next = n
		if n != nil {
			n.prev = h
		}
		dl.head = h
		n = h
	}
	dl.size += len(values)
}

func (dl *DoublyLinkedList) get(index int) (interface{}, bool) {
	if dl.withInRange(index) {
		elem := dl.getElemByIdx(index)
		return elem.value, true
	}

	return nil, false
}

func (dl *DoublyLinkedList) set(index int, value interface{}) bool {
	if dl.withInRange(index) {
		elem := dl.getElemByIdx(index)
		elem.value = value
		return true
	}

	return false
}

func (dl *DoublyLinkedList) remove(index int) (interface{}, bool) {
	if dl.withInRange(index) {
		var value interface{}
		if index == 0 {
			elem := dl.head
			value = elem.value

			if elem.next != nil {
				elem.prev = nil
			}
			dl.head = elem.next
		} else {
			elem := dl.getElemByIdx(index)
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
	arr := make([]interface{}, dl.length())
	index := 0

	for elem := dl.head; elem != nil; elem = elem.next {
		arr[index] = elem.value
		index++
	}

	return arr
}

func (dl *DoublyLinkedList) length() int {
	return dl.size
}

func (dl *DoublyLinkedList) clear() {
	dl.head = nil
	dl.tail = nil
	dl.size = 0
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

// getElemByIdx returns the element at the given index
func (dl *DoublyLinkedList) getElemByIdx(index int) *element {
	elem := dl.head
	for i := 0; i < index; i++ {
		elem = elem.next
	}

	return elem
}

func (dl *DoublyLinkedList) filter(fn utils.FilterFunc) *DoublyLinkedList {
	nal := New()
	for elem := dl.head; elem != nil; elem = elem.next {
		if fn(elem.value) {
			nal.append(elem.value)
		}
	}

	return nal
}

func (dl *DoublyLinkedList) mp(fn utils.MapFunc) *DoublyLinkedList {
	nal := New()
	for elem := dl.head; elem != nil; elem = elem.next {
		nal.append(fn(elem.value))
	}

	return nal
}

func (dl *DoublyLinkedList) concat(lists ...list.List) *DoublyLinkedList {
	nal := dl.clone()
	for _, ls := range lists {
		nal.append(ls.Values()...)
	}

	return nal
}

func (dl *DoublyLinkedList) clone() *DoublyLinkedList {
	nal := New()
	nal.append(dl.values()...)

	return nal
}

func (dl *DoublyLinkedList) reverse() *DoublyLinkedList {
	nal := New()
	for elem := dl.tail; elem != nil; elem = elem.prev {
		nal.append(elem.value)
	}

	return nal
}

func (dl *DoublyLinkedList) swap(a, b int) bool {
	valA, validA := dl.get(a)
	valB, validB := dl.get(b)

	if !validA || !validB {
		return false
	}

	dl.set(a, valB)
	dl.set(b, valA)

	return true
}
