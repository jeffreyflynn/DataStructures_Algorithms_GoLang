package main

import (
	"fmt"
)

type Node struct {
	val        interface{} // the value stored within this node
	prev, next *Node
}

/*********************************************************************************/

type List struct {
	head *Node
	tail *Node
}

func (l *List) append(val interface{}) {
	// note: n.val = input val
	// note: &n.val = hex memory address
	n := &Node{val, nil, nil}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}

	l.tail = n
}

func (l *List) remove(v interface{}) interface{} {
	// if the head of the list contains the value we want to delete
	if l.head.val == v {
		ret := l.head.val

		if l.head.next == nil {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
			l.head.prev = nil
		}

		return ret
	}

	cur := l.head
	for cur.val != v && cur.next != nil {
		cur = cur.next
	}

	// if we made it through the entire list and still haven't found the value
	if cur.val != v {
		return nil
	}

	// if the tail of the list contains the value
	if cur.next == nil {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else { // the value is somewhere in the middle
		cur.prev.next = cur.next
		cur.next.prev = cur.prev
	}

	return cur.val
}

func (l *List) is_empty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func (l *List) print_list() {
	cur := l.head

	fmt.Println("\n\n-----------------------------------\n\n")

	for cur.next != nil {
		fmt.Print(cur.val)
		fmt.Print(" -> ")
		cur = cur.next
	}

	fmt.Print(cur.val)

	fmt.Println("\n\n-----------------------------------\n\n")
}

func main() {
	l := new(List)

	l.append(1)
	l.append(2)
	l.append(3)

	l.print_list()

	l.remove(1)

	l.print_list()
}
