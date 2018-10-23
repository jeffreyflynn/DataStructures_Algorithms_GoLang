package queue

type Queue struct {
	queue []interface{}
	count int
}

// Enqueue add an element to the end of the queue
func (Q *Queue) Enqueue(val interface{}) {
	Q.queue = append(Q.queue, val)
	count++
}

// Dequeue removes an element from the front of the queue
func (Q *Queue) Dequeue() interface{} {
	x := Q.queue[0]
	Q.queue = Q.queue[1:]
	Q.count--
	return x
}
