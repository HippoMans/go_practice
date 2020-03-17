package main

import "fmt"

func sumfunction(first int, second int, channel chan int) {
	channel <- first + second
}

func main() {
	channel := make(chan int)
	go sumfunction(10, 3, channel)

	sum := <-channel
	fmt.Println(sum)
}
