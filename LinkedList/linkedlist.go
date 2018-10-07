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

// add new node to head of linked list
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

// insert some value (val) before some other value (node_val)
func (l *List) Insert_Before(val, node_val interface{}) {
	n := new(Node)
	n.value = val

	cur := l.head

	for cur.next != nil && cur.value != node_val {
		cur = cur.next
	}

	if cur.value == node_val {
		n.prev = cur.prev
		n.next = cur
		cur.prev.next = n
		cur.prev = n
	}

	l.Display()
}

// returns the address of a node given some value
func (l *List) Find(val interface{}) **Node {
	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	if cur.value == val {
		return &cur
	}

	return nil
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

	link.Insert_Before(0)
	link.Insert_Before(1)
	link.Insert_Before(2)
	link.Insert_Before(3)
	link.Insert_Before(4)
	link.Insert_Before(5)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.value)
	fmt.Printf("Tail: %v\n", link.tail.value)
	link.Display()
	fmt.Println("\n==============================\n")

	x := link.Find(5)
	fmt.Println(x)
	fmt.Println("\n==============================\n")
	// fmt.Printf("head: %v\n", link.head.key)
	// fmt.Printf("tail: %v\n", link.tail.key)
	// link.Reverse()
	// fmt.Println("\n==============================\n")
}
