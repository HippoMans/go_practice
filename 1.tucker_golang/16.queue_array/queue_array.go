package main

import "fmt"

func main() {
	queue := []int{}
	var value int

	for i := 1; i < 5; i++ {
		queue = append(queue, i)
	}
	fmt.Println("큐(queue)")
	fmt.Println("입력값 : ", queue)

	fmt.Print("출력값 : ")
	for len(queue) > 0 {
		value, queue = queue[0], queue[1:len(queue)]
		fmt.Printf("%d, ", value)
	}
	fmt.Println()
}
