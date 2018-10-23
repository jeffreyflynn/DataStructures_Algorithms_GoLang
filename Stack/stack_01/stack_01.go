package main

import "fmt"

type Stack struct {
	stack []interface{}
}

// Push adds an element to the end of the stack
func (S *Stack) Push(val interface{}) {
	S.stack = append(S.stack, val)
}

func main() {
	fmt.Println()
}
