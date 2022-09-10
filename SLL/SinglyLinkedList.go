package SLL

import "fmt"

// Node diinisialisasikan sebagai struct
// untuk temp storage ketika menginputkan data

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
	return L.Length() == 0
}

func (L *SinglyLinkedList) Length() int {
	return L.size
}

func (L *SinglyLinkedList) AddFirst(key interface{}) {
	input := &Node{key: key}

	if L.IsEmpty() {
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

	if L.IsEmpty() {
		L.head = input
		L.tail = input
	} else {
		L.tail.next = input
		L.tail = input

	}
	L.size++
}

func (L *SinglyLinkedList) RemoveLast() {
	pointer := L.head
	for i := 0; i < L.Length(); i++ {
		if pointer.next != L.tail {
			pointer = pointer.next
		}
	}
	if L.Length() == 1 {
		pointer = nil
		L.head = pointer
	} else {
		pointer.next = nil
	}
	L.tail = pointer
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
