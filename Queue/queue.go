package main

import "fmt"

type Queue struct {
	slice []int
}

func (q *Queue) Peek() int {
	return q.slice[0]
}

func (q *Queue) Enqueue(val int) {
	q.slice = append(q.slice, val)
}

func (q *Queue) Dequeue() int {
	var ret int = q.Peek()
	q.slice = q.slice[1:]
	return ret
}

func (q *Queue) String() string {
	return fmt.Sprint(q.slice)
}

func main() {
	q := new(Queue)

	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q)

	top := q.Peek()

	fmt.Println("top of the queue: ", top)

	q.Dequeue()

	fmt.Println(q)
}
