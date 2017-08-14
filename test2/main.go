package main

import (
	"fmt"
	"container/list"
)

func main() {
	a := list.New()
	a.PushBack(23)
	fmt.Print(a.Front())
}