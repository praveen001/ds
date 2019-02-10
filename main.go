package main

import (
	"fmt"

	"github.com/praveen001/ds/tree/btree"
	"github.com/praveen001/ds/utils"
)

func main() {
	fmt.Println("Data Structures")

	bt := btree.New(utils.IntComparator)
	v, ok := bt.Get(3)
	fmt.Println(v, ok)
	bt.Set(1, "one")
	fmt.Println("________________INSERTED", 1)
	bt.Set(2, "two")
	fmt.Println("________________INSERTED", 2)
	bt.Set(3, "Three")
	fmt.Println("________________INSERTED", 3)
	bt.Set(4, "Four")
	fmt.Println("________________INSERTED", 4)
	bt.Set(5, "Five")
	fmt.Println("________________INSERTED", 5)
	bt.Set(6, "Six")
	fmt.Println("________________INSERTED", 6)
	bt.Set(7, "Seven")
	fmt.Println("________________INSERTED", 7)

	bt.Set(8, "Eight")
	bt.Set(9, "Nine")
	fmt.Println("________________INSERTED", 9)

	v, ok = bt.Get(1)
	fmt.Println(v, ok)
	v, ok = bt.Get(2)
	fmt.Println(v, ok)
	v, ok = bt.Get(3)
	fmt.Println(v, ok)
	v, ok = bt.Get(4)
	fmt.Println(v, ok)
	v, ok = bt.Get(5)
	fmt.Println(v, ok)
	v, ok = bt.Get(6)
	fmt.Println(v, ok)
	v, ok = bt.Get(7)
	fmt.Println(v, ok)
	v, ok = bt.Get(8)
	fmt.Println(v, ok)
	v, ok = bt.Get(9)
	fmt.Println(v, ok)
	v, ok = bt.Get(10)
	fmt.Println(v, ok)
	v, ok = bt.Get(11)
	fmt.Println(v, ok)
	v, ok = bt.Max()
	fmt.Println(v, ok)

	bt.Set(9, "Nine")

	fmt.Println(bt.Keys())
	fmt.Println(bt.PreOrder())
	fmt.Println(bt.PostOrder())
	fmt.Println(bt.Len())

	fmt.Println("-----------------------------")
	fmt.Println(bt.Remove(4))
	fmt.Println(bt.InOrder())
}
