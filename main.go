package main

import (
	"fmt"

	"github.com/praveen001/ds/list/arraylist"
)

func main() {
	al := arraylist.New()
	al.Add(10, 20, 30, 40, 50)

	e, _ := al.Get(4)
	al.Remove(5)
	fmt.Println(al, e, al.Contains(15), al.IndexOf(30), al.Pop(), al, al.Push(20))
}
