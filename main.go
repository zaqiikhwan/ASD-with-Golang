package main

import (
	"fmt"
	"asd-golang/SLL"
)

func main() {
	link := SLL.SinglyLinkedList{}
	link.AddFirst(5)
	// link.AddLast(9)
	link.AddLast(10)
	// link.RemoveLast()
	// link.RemoveLast()
	link.RemoveFirst()
	link.RemoveFirst()
	fmt.Println(link.Length())
	// link.InsertBefore(13)
	// link.InsertBefore(22)
	// link.InsertBefore(28)
	// link.InsertBefore(36)

	link.Display()
}