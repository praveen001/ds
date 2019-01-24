package main

import (
	"fmt"

	"github.com/praveen001/ds/tree/binarytree"
)

func main() {
	// al := arraylist.New()
	// al.Append(10, 20, 30, 40, 50)
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

	bt := binarytree.New()

	bt.Insert(10)

	bt.Insert(5)

	bt.Insert(2)

	bt.Insert(20)

	fmt.Println(bt.InOrder())
}
