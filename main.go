package main

import "fmt"

func main() {
	a := 13
	b := &a
	c := *b

	fmt.Println("\n\n------------------------------\n\n")

	fmt.Println("The value of | a | is  ~  ", a)
	fmt.Println("The value of | b | is  ~  ", b)

	fmt.Println("\n\n------------------------------\n\n")

	fmt.Println("The value of | *b | is  ~  ", *b)

	fmt.Println("\n\n------------------------------\n\n")

	*b++

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("\n\n------------------------------\n\n")

}
