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
	S.count++
}

func (S *Stack) Pop() {
	S.stack = S.stack[:S.count-1]
	S.count--
}

func (S *Stack) Peek() interface{} {
	if S.count > 0 {
		return S.stack[0]
	} else {
		return nil
	}
}

func main() {
	S := Stack{}

	S.count = 0

	S.Push(1)
	S.Push(2)
	S.Push(3)

	S.Print()

	S.Pop()

	S.Print()

	S.Peek()
}
