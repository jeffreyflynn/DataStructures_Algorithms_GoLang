package main

import "fmt"

type Node struct {
	val interface{}
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

func (L *LinkedList)