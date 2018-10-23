package queue

type Queue struct {
	queue []interface{}
	count int
}

// Enqueue add an element to the end of the queue
func (Q *Queue) Enqueue(val interface{}) {
	Q.queue = append(Q.queue, val)
	Q.count++
}

// Dequeue removes an element from the front of the queue
func (Q *Queue) Dequeue() interface{} {
	x := Q.queue[0]
	Q.queue = Q.queue[1:]
	Q.count--
	return x
}

// Size returns the number of elements in the queue
func (Q *Queue) Size() int {
	return Q.count
}

// Peek returns the top of the queue without removing any elements
func (Q *Queue) Peek() interface{} {
	return Q.queue[0]
}
