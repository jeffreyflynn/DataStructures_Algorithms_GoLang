package queue

type Queue struct {
	items []interface{}
}

// New initializes a new queue
func (Q *Queue) New() *Queue {
	var slice []interface{}
	Q.items = slice
	return Q
}

// Enqueue adds an item to the end of the queue
func (Q *Queue) Enqueue(val interface{}) {
	Q.items = append(Q.items, val)
}

// Dequeue removes an item from the beginning of the queue
func (Q *Queue) Dequeue() interface{} {
	item := Q.items[0]
	Q.items = Q.items[1:len(Q.items)]
	return item
}

// IsEmpty returns true if the queue is empty
func (Q *Queue) IsEmpty() bool {
	return len(Q.items) == 0
}

// Size returns the length of the queue
func (Q *Queue) Size() int {
	return len(Q.items)
}
