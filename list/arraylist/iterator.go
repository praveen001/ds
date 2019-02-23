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
func (i *Iterator) Previous() bool {
	if i.curIndex != -1 {
		i.curIndex--
	}
	return i.HasPrevious()
}

// Next ..
func (i *Iterator) Next() bool {
	if i.HasNext() {
		i.curIndex++
	}
	return i.HasNext()
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

// Set ..
func (i *Iterator) Set(v interface{}) bool {
	return i.list.Set(i.curIndex, v)
}

// Add .
func (i *Iterator) Add(v interface{}) bool {
	return i.list.Insert(i.curIndex, v)
}

// Remove ..
func (i *Iterator) Remove() bool {
	_, ok := i.list.Remove(i.curIndex)
	return ok
}
