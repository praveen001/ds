package main

import (
	"fmt"

	"github.com/praveen001/ds/utils"

	"github.com/praveen001/ds/heap/binaryheap"
	"github.com/praveen001/ds/list/arraylist"
	"github.com/praveen001/ds/list/linkedlist"
	"github.com/praveen001/ds/maps/hashmap"
	"github.com/praveen001/ds/queue"
	"github.com/praveen001/ds/stack"
	"github.com/praveen001/ds/tree/avltree"
	"github.com/praveen001/ds/tree/redblacktree"
)

func tryArrayList() {
	al := arraylist.New()

	al.Append(1, 2, 3, 4, 5, 6)
	fmt.Println(al)

	al.Swap(0, 1)
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

	al.Swap(0, 1)
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
	al.Remove(3)
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

	bt.Set(100, 100)
	bt.Set(50, 50)
	bt.Set(20, 20)
	bt.Set(5, 5)
	bt.Set(10, 10)
	bt.Set(60, 60)
	bt.Set(55, 55)
	bt.Set(80, 80)
	bt.Set(200, 200)

	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())
	bt.Remove(50)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Remove(10)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())
	bt.Remove(20)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Remove(100)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Remove(200)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Remove(55)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Remove(60)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Remove(80)
	fmt.Println(bt.InOrder(), bt.Height(), bt.Length())

	bt.Remove(5)
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

	// at.RAdd(10)
	// at.RAdd(3)
	// at.RAdd(18)
	// at.RAdd(2)
	// at.RAdd(4)
	// at.RAdd(13)
	// at.RAdd(40)

	// at.RAdd(39)

	// at.RAdd(12)
	// at.RAdd(38)
	// at.RAdd(14)
	// at.RAdd(11)

	// at.Remove(10)
	// at.Remove(12)

	// fmt.Println(at.InOrder(), at.Length())

	at.RSet(9, 9)
	at.RSet(5, 5)
	at.Set(10, 10)
	at.Set(0, 0)
	at.Set(6, 6)
	at.Set(11, 11)
	at.RSet(-1, 1)
	at.RSet(1, 1)
	at.RSet(2, 2)
	fmt.Println(at.PreOrder(), at.Length())

	at.Remove(10)
	fmt.Println(at.PreOrder(), at.Length())
}

func tryBinaryHeap() {
	bh := binaryheap.New(func(a, b interface{}) int {
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

	bh.Push(56, 34, 23, 78, 43, 61, 6, 3, 90, 30)

	fmt.Println(bh)

	fmt.Println(bh.Pop())

	fmt.Println(bh)
	// fmt.Println(bh.Pop())
	// fmt.Println(bh.Pop())
}

func tryHashMap() {
	hm := hashmap.New()

	hm.Set("hello", "world")

	fmt.Println(hm.Get("hello"))

	fmt.Println(hm.Keys())

	fmt.Println(hm.Values())
	fmt.Println(hm.Contains("hello"))

	fmt.Println(hm.Remove("hello"))
	hm.Set("hello2", "world2")
	fmt.Println(hm.Length())
	hm.Clear()
	fmt.Println(hm.Values())

}

func main() {
	// tryArrayList()
	// tryLinkedList()
	// tryStack()
	// tryQueue()
	// tryBinaryTree()
	// tryAVLTree()
	// tryBinaryHeap()
	// tryHashMap()

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

	// bt := binarytree.New(utils.IntComparator)
	// var wg sync.WaitGroup

	// for i := 0; i < 100; i++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		wg.Done()
	// 		bt.Add(10)
	// 	}()

	// 	go func() {
	// 		wg.Done()
	// 		bt.Remove(10)
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(bt.InOrder())

	rbt := redblacktree.New(utils.IntComparator)

	rbt.Add(11)
	rbt.Add(2)
	rbt.Add(14)
	rbt.Add(1)
	rbt.Add(7)
	rbt.Add(15)
	rbt.Add(5)
	rbt.Add(8)

	fmt.Println(rbt.InOrder())

	fmt.Println("------------------")
	rbt.Add(4)
	fmt.Println(rbt.InOrder())
}
