package main

import (
	"fmt"
	"os"
)

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i

			if i == 5 {
				close(channel)
				os.Exit(1)
			}
		}
	}()

	for i := range channel {
		fmt.Println(i)
	}
}
