package main

import (
	"fmt"
)

type Stack struct {
	slice []int
}

// adds int to the top of the stack
func (s *Stack) Push(val int) {
	s.slice = append(s.slice, val)
}

// return the top item from the stack
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

// removes element from top of stack
func (s *Stack) Pop() int {
	ret := s.Peek()
	s.slice = s.slice[0 : len(s.slice)-1] // inclusive:exclusive
	return ret
}

// print stack
func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

func main() {
	s := new(Stack)

	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s)

	top := s.Peek()

	fmt.Println(top)

	s.Pop()

	fmt.Println(s)
}
