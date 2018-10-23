package main

import (
	"fmt"
	"testing"
)

var s Stack

func TestPush(t *testing.T) {
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s.stack)

	if size := s.Length(); size != 5 {
		t.Errorf("wrong count, expected 5 and got %d", size)
	}
}
