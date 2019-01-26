package main

import (
	"fmt"

	"github.com/praveen001/ds/list/arraylist"
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/queue"
	"github.com/praveen001/ds/stack"
	"github.com/praveen001/ds/tree/avltree"
)

func tryArrayList() {
	al := arraylist.New()

	al.Append(1, 2, 3, 4, 5, 6)
	fmt.Println(al)

	al.Prepend(0, -1)
	fmt.Println(al)

	fmt.Println(al.Filter(func(x interface{}) bool {
		if x.(int)%2 != 0 {
			return true
		}
		return false
	}))

	al.Remove(0)
	fmt.Println(al)

	al.Set(0, 10)
	fmt.Println(al)

	tal := al.Map(func(x interface{}) interface{} {
		return x.(int) * 2
	})

	fmt.Println(al.Concat(tal))

	fmt.Println(al.Contains(6), al.IndexOf(1))

	cl := al.Clone()
	cl.Clear()
	fmt.Println(cl, al)
}

func tryLinkedList() {
	al := linkedlist.New()

	al.Append(1, 2, 3, 4, 5, 6)
	fmt.Println(al)

	al.Prepend(0, -1)
	fmt.Println(al)

	fmt.Println(al.Filter(func(x interface{}) bool {
		if x.(int)%2 != 0 {
			return true
		}
		return false
	}))

	al.Remove(0)
	fmt.Println(al)

	al.Set(0, 10)
	fmt.Println(al)

	tal := al.Map(func(x interface{}) interface{} {
		return x.(int) * 2
	})

	fmt.Println(al.Concat(tal))

	fmt.Println(al.Contains(6), al.IndexOf(1))

	cl := al.Clone()
	cl.Clear()
	fmt.Println(cl, al)
}

func tryStack() {
	s := stack.New()
	fmt.Println(s)

	s.Push(10)
	s.Push(20)
	fmt.Println(s)

	fmt.Println(s.Peek())

	s.Pop()
	fmt.Println(s)

	fmt.Println(s.Length())

	s.Clear()
	fmt.Println(s)
}

func tryQueue() {
	q := queue.New()
	fmt.Println(q)

	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println(q)

	fmt.Println(q.Peek())

	q.Dequeue()
	fmt.Println(q)

	fmt.Println(q.Length())

	q.Clear()
	fmt.Println(q)
}

func tryBinaryTree() {
	bt := avltree.New(func(a, b interface{}) int {
		at := a.(int)
		bt := b.(int)

		if at > bt {
			return 1
		} else if at < bt {
			return -1
		} else {
			return 0
		}
	})

	bt.Insert(100)
	bt.Insert(50)
	bt.Insert(20)
	bt.Insert(5)
	bt.Insert(10)
	bt.Insert(60)
	bt.Insert(55)
	bt.Insert(80)
	bt.Insert(200)

	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())
	bt.Delete(50)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Delete(10)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())
	bt.Delete(20)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Delete(100)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Delete(200)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Delete(55)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Delete(60)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Delete(80)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Delete(5)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())
}

func tryAVLTree() {
	at := avltree.New(func(a, b interface{}) int {
		awt := a.(int)
		bwt := b.(int)

		if awt > bwt {
			return 1
		} else if awt < bwt {
			return -1
		} else {
			return 0
		}
	})

	// at.RInsert(10)
	// at.RInsert(3)
	// at.RInsert(18)
	// at.RInsert(2)
	// at.RInsert(4)
	// at.RInsert(13)
	// at.RInsert(40)

	// at.RInsert(39)

	// at.RInsert(12)
	// at.RInsert(38)
	// at.RInsert(14)
	// at.RInsert(11)

	// at.Delete(10)
	// at.Delete(12)

	// fmt.Println(at.InOrder(), at.Length())

	at.RInsert(9)
	at.RInsert(5)
	at.Insert(10)
	at.Insert(0)
	at.Insert(6)
	at.Insert(11)
	at.RInsert(-1)
	at.RInsert(1)
	at.RInsert(2)
	fmt.Println(at.PreOrder(), at.Length())

	at.Delete(10)
	fmt.Println(at.PreOrder(), at.Length())
}

func main() {
	// tryArrayList()
	// tryLinkedList()
	// tryStack()
	// tryQueue()
	tryBinaryTree()
	// tryAVLTree()

	// s := stack.New()
	// s.Push(10)
	// s.Push(20)
	// s.Push(30)
	// s.Push(40)
	// s.Push(50)
	// fmt.Println("Stack", s, s.Length())
	// for {
	// 	p, ok := s.Pop()
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(p)
	// }
}
