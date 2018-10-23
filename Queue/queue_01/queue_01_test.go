package queue

import "testing"

func TestQueue(t *testing.T) {
	var q Queue

	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	if x := q.Size(); x != 6 {
		t.Errorf("error with enqueue, expected 6 elements, returned %d", x)
	}

	if x := q.Dequeue(); x != 0 {
		t.Errorf("error with dequeue, expected to dequeue 0, but received %d", x)
	}

	if x := q.Peek(); x != 1 {
		t.Errorf("error with peek, expected to 1, received %d", x)
	}
}
