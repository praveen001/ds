package main

import (
	"fmt"

	"github.com/praveen001/ds/tree/btree"
)

func main() {
	fmt.Println("Data Structures")

	bt := btree.New(func(a, b interface{}) int {
		astr := a.(string)
		bstr := b.(string)

		if []rune(astr)[0] > []rune(bstr)[0] {
			return 1
		} else if []rune(astr)[0] < []rune(bstr)[0] {
			return -1
		} else {
			return 0
		}
	}, 5)

	bt.Set("c", "c")
	bt.Set("n", "n")
	bt.Set("g", "g")
	bt.Set("a", "a")
	bt.Set("h", "h")
	bt.Set("e", "e")
	bt.Set("k", "k")
	bt.Set("q", "q")
	bt.Set("m", "m")
	bt.Set("f", "f")
	bt.Set("w", "w")
	bt.Set("l", "l")
	bt.Set("t", "t")
	bt.Set("z", "z")
	bt.Set("d", "d")
	bt.Set("p", "p")
	bt.Set("r", "r")
	bt.Set("x", "x")
	bt.Set("y", "y")
	bt.Set("s", "s")

	bt.Remove("h")
	fmt.Println("------------------")
	bt.Remove("t")

	fmt.Println(bt.InOrder())
}
