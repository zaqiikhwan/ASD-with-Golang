package main

import (
	"fmt"
)

// Node diinisialisasikan sebagai
// struct untuk temp storage ketika menginputkan data
type Node struct {
	next *Node
	key interface{}
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (L *SinglyLinkedList) IsEmpty() bool {
	return L.length() == 0
}

func (L *SinglyLinkedList) length() int {
	return L.size
}

func (L *SinglyLinkedList) AddFirst(key interface{}) {
	input := &Node{key: key}

	if (L.IsEmpty()) {
		L.head = input
		L.tail = input
	} else {
		input.next = L.head
		L.head = input
	}
	L.size++
}

func (L *SinglyLinkedList) RemoveFirst() {
	if L.IsEmpty() {
		fmt.Println("SinglyLinkedList Kosong")
	} else {
		L.head = L.head.next
		L.size--
	}
}

func (L *SinglyLinkedList) AddLast(key interface{}) {
	input := &Node{key: key}

	if (L.IsEmpty()) {
		L.head = input
		L.tail = input
	} else {
		L.tail.next = input
		L.tail = input
		
	}
	L.size++
}

func (L *SinglyLinkedList) RemoveLast() {
	for i := 0; i < L.length(); i++ {
		if L.head.next != L.tail {
			L.head = L.head.next
		}
	}
	if L.length() > 1 {
		L.head.next = nil
	}
	L.tail = L.head
	L.size--
}

func (l *SinglyLinkedList) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v -> ", list.key)
		list = list.next
	}
	fmt.Print()
}

func main() {
	link := SinglyLinkedList{}
	// link.AddFirst(5)
	link.AddLast(9)
	link.AddLast(10)
	// link.RemoveFirst()
	link.RemoveLast()
	fmt.Println(link.length())
	// link.InsertBefore(13)
	// link.InsertBefore(22)
	// link.InsertBefore(28)
	// link.InsertBefore(36)

	link.Display()
}