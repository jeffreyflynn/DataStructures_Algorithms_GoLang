package queue

import "testing"

var q Queue

func initQueue() *Queue {
	if q.items == nil {
		q = Queue{}
		q.New()
	}
	return &q
}

func TestEnqueue(t *testing.T) {
	q := initQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if size := q.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestDequeue(t *testing.T) {
	q.Dequeue()

	if size := len(q.items); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	q.Dequeue()
	q.Dequeue()

	if size := len(q.items); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	if !q.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}
