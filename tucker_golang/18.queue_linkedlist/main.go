package main

import (
	"fmt"
	"queue_linkedlist/Architecture"
)

func main() {
	queue := Architecture.NewQueue()
	fmt.Println("queue")
	for i := 1; i <= 5; i++ {
		queue.Push(i)
	}
	queue.PrintQueue()

	fmt.Println("All queue element pop")
	for !queue.IsEmpty() {
		queue_value := queue.Pop()
		fmt.Printf("%d  ", queue_value)
	}
	fmt.Println()
}
