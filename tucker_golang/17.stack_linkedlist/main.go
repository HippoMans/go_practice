package main

import (
	"fmt"
	"stack_linkedlist/Architecture"
)

func main() {
	stacklinked := Architecture.NewStack()

	fmt.Println("stack")
	for i := 1; i <= 5; i++ {
		stacklinked.Push(i)
	}
	stacklinked.PrintStack()

	fmt.Println("All stack element pop")
	for !stacklinked.IsEmpty() {
		val := stacklinked.Pop()
		fmt.Printf("%d ", val)
	}
	fmt.Println("\nIs empty ? ")
	if stacklinked.IsEmpty() {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
