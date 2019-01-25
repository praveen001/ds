package tree

import "github.com/praveen001/ds/list"

// Tree ..
type Tree interface {
	Insert(value interface{})
	Delete(value interface{}) bool
	Contains(value interface{}) bool
	Height() int
	Min() interface{}
	Max() interface{}
	Count() int
	Empty()
	InOrder() list.List
}
