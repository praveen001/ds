package doublylinkedlist

// Iterator .
type Iterator struct {
	list     *DoublyLinkedList
	curIndex int
	cur      *element
}

// Iterator ..
func (al *DoublyLinkedList) Iterator() *Iterator {
	return &Iterator{
		list:     al.Clone().(*DoublyLinkedList),
		curIndex: -1,
	}
}

// Next ..
func (i *Iterator) Next() (interface{}, bool) {
	if i.curIndex < i.list.Len() {
		i.curIndex++
	}
	if !i.list.withInRange(i.curIndex) {
		return nil, false
	}
	if i.curIndex == 0 {
		i.cur = i.list.head
	} else {
		i.cur = i.cur.next
	}

	return i.Value(), true
}

// Previous ..
func (i *Iterator) Previous() (interface{}, bool) {
	if i.curIndex > -1 {
		i.curIndex--
	}
	if !i.list.withInRange(i.curIndex) {
		return nil, false
	}
	if i.curIndex == i.list.Len()-1 {
		i.cur = i.list.tail
	} else {
		i.cur = i.cur.prev
	}

	return i.Value(), true
}

// Index ..
func (i *Iterator) Index() int {
	return i.curIndex
}

// Value ..
func (i *Iterator) Value() interface{} {
	return i.cur.value
}
