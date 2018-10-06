package main

import "testing"

func TestStack(t *testing.T) {
	var s *Stack = new(Stack)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	for x := 5; x > 0; x-- {
		top1 := s.Peek()
		top2 := s.Pop()

		if top1 != top2 {
			t.Error("Peek and Pop are not returning the same value.")
		}

		if top1 != x {
			t.Error("Peek is returning an incorrect value.", top1)
		}

		if top2 != x {
			t.Error("Pop is returning an incorrect value.", top2)
		}
	}
}
