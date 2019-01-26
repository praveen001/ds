package arraylist

import (
	"fmt"

	"github.com/praveen001/ds/list"
	"github.com/praveen001/ds/utils"
)

func (al *ArrayList) append(values ...interface{}) {
	al.elements = append(al.elements, values...)
	al.size += len(values)
}

func (al *ArrayList) prepend(values ...interface{}) {
	length := len(values)

	newSubArr := make([]interface{}, length)
	for i := 0; i < len(values); i++ {
		newSubArr[i] = values[length-1-i]
	}

	al.elements = append(newSubArr, al.elements...)
	al.size += len(values)
}

func (al *ArrayList) get(index int) (interface{}, bool) {
	if al.withInRange(index) {
		return al.elements[index], true
	}

	return nil, false
}

func (al *ArrayList) set(index int, value interface{}) bool {
	if al.withInRange(index) {
		al.elements[index] = value
		return true
	}

	return false
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

func (al *ArrayList) length() int {
	return al.size
}

func (al *ArrayList) clear() {
	al.elements = make([]interface{}, 0)
	al.size = 0
}

func (al *ArrayList) withInRange(index int) bool {
	return index > -1 && index < al.size
}

func (al *ArrayList) string() string {
	return fmt.Sprintf("%v", al.elements)
}

func (al *ArrayList) filter(fn utils.FilterFunc) *ArrayList {
	nal := New()
	for _, value := range al.values() {
		if fn(value) {
			nal.Append(value)
		}
	}

	return nal
}

func (al *ArrayList) mp(fn utils.MapFunc) *ArrayList {
	nal := New()
	for _, value := range al.values() {
		nal.append(fn(value))
	}

	return nal
}

func (al *ArrayList) concat(lists ...list.List) *ArrayList {
	nal := al.clone()
	for _, ls := range lists {
		nal.append(ls.Values()...)
	}

	return nal
}

func (al *ArrayList) clone() *ArrayList {
	nal := New()
	nal.append(al.values()...)

	return nal
}

func (al *ArrayList) reverse() *ArrayList {
	nal := New()
	for _, value := range al.values() {
		nal.prepend(value)
	}

	return nal
}

func (al *ArrayList) swap(a, b int) bool {
	valA, validA := al.get(a)
	valB, validB := al.get(b)

	if !validA || !validB {
		return false
	}

	al.set(a, valB)
	al.set(b, valA)

	return true
}
