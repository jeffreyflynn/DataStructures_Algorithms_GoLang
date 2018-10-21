package main

import "fmt"

// stack is Last In, First Out
type Stack struct {
	stack []interface{}
	count int
}

func (S *Stack) Print() {
	fmt.Println(S.stack)
}

func (S *Stack) Push(val interface{}) {
	x := append(S.stack, val)
	S.stack = x
	S.count += 1
}

func main() {
	S := Stack{}

	S.count = 0

	S.Push(1)
	S.Push(2)
	S.Push(3)

	S.Print()
}
