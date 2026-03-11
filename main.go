package main

import (
	"fmt"

	"cryogon/dsa-practice/ds"
)

func main() {
	list := ds.NewList()
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)
	fmt.Printf("Items: %d\n", list.GetAll())
	fmt.Printf("Len:%d\tMax Items:%d", list.Length(), list.MaxSize())
}
