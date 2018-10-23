package main

import (
	"fmt"
	"testing"
)

var s Stack

func TestStack(t *testing.T) {
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s.stack)

	if x := s.Length(); x != 5 {
		t.Errorf("wrong count, expected 5 and got %d", x)
	}

	if x := s.Pop(); x != 4 {
		t.Errorf("error popping, expected 4 and got %d", x)
	}

	if x := s.Peek(); x != 3 {
		t.Errorf("error peeking, expected 3 and got %d", x)
	}
}
