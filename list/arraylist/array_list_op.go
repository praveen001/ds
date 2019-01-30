package arraylist

import (
	"fmt"

	"github.com/praveen001/ds/ds"
)

func (al *ArrayList) len() int {
	return al.size
}

func (al *ArrayList) front() (interface{}, bool) {
	return al.get(0)
}

func (al *ArrayList) back() (interface{}, bool) {
	return al.get(al.len() - 1)
}

func (al *ArrayList) pushFront(v interface{}) {
	nal := make([]interface{}, al.size+1)
	copy(nal, []interface{}{v})
	copy(nal[1:], al.elements)
	al.elements = nal
	al.size++
}

func (al *ArrayList) pushBack(v interface{}) {
	al.elements = append(al.elements, v)
	al.size++
}

func (al *ArrayList) set(i int, v interface{}) bool {
	if al.withInRange(i) {
		al.elements[i] = v
		return true
	}
	return false
}

func (al *ArrayList) get(i int) (interface{}, bool) {
	if al.withInRange(i) {
		return al.elements[i], true
	}
	return nil, false
}

func (al *ArrayList) remove(index int) (interface{}, bool) {
	if al.withInRange(index) {
		value := al.elements[index]

		newElems := make([]interface{}, al.size-1)

		copy(newElems, al.elements[:index])
		copy(newElems[index:], al.elements[index+1:])

		al.elements = newElems
		al.size--

		return value, true
	}

	return nil, false
}

func (al *ArrayList) clear() {
	al.elements = make([]interface{}, 0)
	al.size = 0
}

func (al *ArrayList) pushBackList(l ds.List) {
	al.elements = append(al.elements, l.Values()...)
	al.size += l.Len()
}

func (al *ArrayList) pushFrontList(l ds.List) {
	al.elements = append(l.Values(), al.elements...)
	al.size += l.Len()
}

func (al *ArrayList) contains(value interface{}) bool {
	if al.indexOf(value) == -1 {
		return false
	}

	return true
}

func (al *ArrayList) indexOf(value interface{}) int {
	for i, element := range al.elements {
		if element == value {
			return i
		}
	}

	return -1
}

func (al *ArrayList) values() []interface{} {
	return al.elements
}

func (al *ArrayList) clone() *ArrayList {
	nal := New()
	nal.pushFrontList(al)

	return nal
}

func (al *ArrayList) swap(a, b int) bool {
	if al.withInRange(a) && al.withInRange(b) {
		al.elements[a], al.elements[b] = al.elements[b], al.elements[a]
		return true
	}
	return false
}

func (al *ArrayList) withInRange(index int) bool {
	return index > -1 && index < al.size
}

func (al *ArrayList) string() string {
	return fmt.Sprintf("%v", al.elements)
}
