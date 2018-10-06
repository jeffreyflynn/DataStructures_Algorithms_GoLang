package main

import "fmt"

type Node struct {
	// pointers to the node's previous and next nodes
	next, prev *Node

	// the value stored within this node
	value interface{}
}

/*********************************************************************************/

type List struct {
	head *Node
	tail *Node
}

// add new node to tail of linked list
func (l *List) Insert_Tail(val interface{}) {
	n := new(Node)
	n.value = val

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
}

func (l *List) Insert_Head(val interface{}) {
	n := new(Node)
	n.value = val

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head.prev = n
		l.head = n
	}
}

func (l *List) Display() {
	cur := l.head

	for cur.next != nil {
		fmt.Printf("%+v -> ", cur.value)
		cur = cur.next
	}

	fmt.Printf("%+v", cur.value)
	fmt.Println()
}

func main() {
	link := List{}

	link.Insert_Head(0)
	link.Insert_Head(1)
	link.Insert_Head(2)
	link.Insert_Head(3)
	link.Insert_Head(4)
	link.Insert_Head(5)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.value)
	fmt.Printf("Tail: %v\n", link.tail.value)
	link.Display()
	fmt.Println("\n==============================\n")
	// fmt.Printf("head: %v\n", link.head.key)
	// fmt.Printf("tail: %v\n", link.tail.key)
	// link.Reverse()
	// fmt.Println("\n==============================\n")
}
