package main

import (
	"fmt"

	"github.com/praveen001/ds/tree"
	"github.com/praveen001/ds/tree/binarytree"
)

func main() {
	// al := linkedlist.New()
	// al.Append(10, 20, 30, 40, 50)
	// fmt.Println(al.Values())
	// fmt.Println("Append 10, 20, 30, 40, 50 =>", al)

	// al.Prepend(-10, 0, 5)
	// fmt.Println("Prepend -10, 0 =>", al)

	// // fmt.Println(al.Get(0), al.Get(-1), al.Get(6), al.Get(10))

	// al.Set(0, -20)
	// al.Set(6, 100)
	// fmt.Println(al)

	// fmt.Println("Removing")
	// fmt.Println(al.Remove(-1))
	// fmt.Println(al.Remove(0))
	// fmt.Println(al.Remove(6))
	// fmt.Println(al.Remove(3))
	// fmt.Println(al.Remove(10))
	// fmt.Println(al)

	// fmt.Println(al.Contains(40), al.Contains(100))
	// fmt.Println(al.IndexOf(40), al.IndexOf(100), al.Size())

	// var l list.List

	// s := stack.New()
	// s.Push(10)

	// fmt.Println(s.Pop())

	bt := binarytree.New(func(a, b interface{}) int {
		at := a.(string)
		bt := b.(string)

		if len(at) > len(bt) {
			return 1
		} else if at < bt {
			return -1
		} else {
			return 0
		}
	})

	var t tree.Tree
	t = bt

	t.Insert("aa")
	t.Insert("aaaa")
	t.Insert("aaa")
	t.Insert("a")
	t.Delete("a")
	fmt.Println(t.InOrder())

	// bt.Insert(100)
	// bt.Insert(50)
	// bt.Insert(20)
	// bt.Insert(5)
	// bt.Insert(10)
	// bt.Insert(60)
	// bt.Insert(55)
	// bt.Insert(80)
	// bt.Insert(200)

	// fmt.Println(bt.InOrder())
	// bt.Delete(50)
	// fmt.Println(bt.InOrder())

	// bt.Delete(10)
	// fmt.Println(bt.InOrder())
	// bt.Delete(20)
	// fmt.Println(bt.InOrder())

	// bt.Delete(100)
	// fmt.Println(bt.InOrder())

	// bt.Delete(200)
	// fmt.Println(bt.InOrder())

	// bt.Delete(55)
	// fmt.Println(bt.InOrder())

	// bt.Delete(60)
	// fmt.Println(bt.InOrder())

	// bt.Delete(80)
	// fmt.Println(bt.InOrder())

	// bt.Delete(5)
	// fmt.Println(bt.InOrder())
}
