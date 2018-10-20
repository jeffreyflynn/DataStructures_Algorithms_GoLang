package main

import "fmt"

type Node struct {
	val  interface{}
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (L *LinkedList) Insert(val interface{}) {
	node := &Node{val, nil, nil}

	if L.head == nil {
		L.head = node
		L.tail = node
	} else {
		L.tail.next = node
		node.prev = L.tail
		L.tail = node
	}
}

func (L *LinkedList) Print() {
	cur := L.head

	for cur != nil {
		fmt.Printf("%+v -> ", cur.val)
		cur = cur.next
	}

	fmt.Println()
}

func (L *LinkedList) Remove(val interface{}) {
	cur := L.head

	for cur.val != val && cur.next != nil {
		cur = cur.next
	}

	if cur.val == val {
		// If the node to be removed is sandwiched between 2 other nodes ...
		if cur.prev != nil && cur.next != nil {
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			// else if there is only 1 node in the List
		} else if cur.prev == nil && cur.next == nil {
			L.head = nil
			// else if the node to be removed is the head of the List
		} else if cur.prev == nil {
			L.head = L.head.next
			L.head.prev = nil
			// else if the node to be removed is the tail of the List
		} else {
			L.tail = L.tail.prev
			L.tail.next = nil
		}
	}
}

func main() {
	L := LinkedList{nil, nil}

	L.Insert(1)
	L.Insert(2)
	L.Insert(3)
	L.Insert(4)
	L.Insert(5)

	L.Print()

	L.Remove(5)

	L.Print()
}
