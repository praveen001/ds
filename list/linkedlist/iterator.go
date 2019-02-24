package linkedlist

import (
	"github.com/praveen001/ds/list/arraylist"
)

// Iterator .
type Iterator struct {
	*arraylist.Iterator
}

// Iterator ..
func (ll *LinkedList) Iterator() *Iterator {
	al := arraylist.New()
	for _, v := range ll.Values() {
		al.PushBack(v)
	}
	return &Iterator{al.Iterator()}
}

// Next ..
func (i *Iterator) Next() (interface{}, bool) {
	return i.Iterator.Next()
}

// Previous ..
func (i *Iterator) Previous() (interface{}, bool) {
	return i.Iterator.Previous()
}

// Index ..
func (i *Iterator) Index() int {
	return i.Iterator.Index()
}

// Value ..
func (i *Iterator) Value() interface{} {
	return i.Iterator.Value()
}
