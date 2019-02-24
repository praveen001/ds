package arraylist

// Iterator .
type Iterator struct {
	list     *ArrayList
	curIndex int
}

// Iterator ..
func (al *ArrayList) Iterator() *Iterator {
	return &Iterator{
		list:     al,
		curIndex: -1,
	}
}

// HasPrevious ..
func (i *Iterator) HasPrevious() bool {
	if i.curIndex < 1 {
		return false
	}
	return true
}

// HasNext ..
func (i *Iterator) HasNext() bool {
	if i.curIndex == i.list.Len() {
		return false
	}
	return true
}

// Previous ..
func (i *Iterator) Previous() (interface{}, bool) {
	if i.curIndex != -1 {
		i.curIndex--
	}
	return i.Value(), i.curIndex != -1
}

// Next ..
func (i *Iterator) Next() (interface{}, bool) {
	if i.curIndex != i.list.Len() {
		i.curIndex++
	}
	return i.Value(), i.curIndex != i.list.Len()
}

// Index ..
func (i *Iterator) Index() int {
	return i.curIndex
}

// Value ..
func (i *Iterator) Value() interface{} {
	v, _ := i.list.Get(i.curIndex)
	return v
}
