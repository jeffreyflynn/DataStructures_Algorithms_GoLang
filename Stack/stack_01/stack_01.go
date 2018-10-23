package main

type Stack struct {
	stack []interface{}
	count int
}

// Push adds an element to the end of the stack
func (S *Stack) Push(val interface{}) {
	S.stack = append(S.stack, val)
	S.count++
}

// Pop removes and returns the last item in the stack
func (S *Stack) Pop() interface{} {
	p := S.stack[len(S.stack)-1]
	S.stack = S.stack[:len(S.stack)-1]
	S.count--
	return p
}

// Length returns the number of elements in the stack
func (S *Stack) Length() int {
	return S.count
}

func main() {
	
}
