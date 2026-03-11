package main

import (
	"fmt"

	"cryogon/dsa-practice/ds"
)

func main() {
	list := ds.NewList(1, 2, 3, 4, 5)
	i, ok := list.Pop()
	if ok {
		fmt.Printf("Popped: %d\n", i)
	}

	i, ok = list.Shift()
	if ok {
		fmt.Printf("Shifted: %d\n", i)
	}

	fmt.Printf("Items: %d\n", list.GetAll())
	fmt.Printf("Len:%d\tMax Items:%d", list.Length(), list.MaxSize())
}
