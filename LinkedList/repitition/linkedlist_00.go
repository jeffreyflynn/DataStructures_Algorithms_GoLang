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
	
}