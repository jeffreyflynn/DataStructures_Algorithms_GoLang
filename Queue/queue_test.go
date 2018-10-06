package main

import "testing"

func TestQueue(t *testing.T) {
	q := new(Queue)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for x := 1; x <= 3; x++ {
		top1 := q.Peek()
		top2 := q.Dequeue()

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
