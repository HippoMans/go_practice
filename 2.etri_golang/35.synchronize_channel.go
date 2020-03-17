package main

import (
	"fmt"
	"time"
)

func main() {
	right := make(chan bool)
	count := 10

	go func() {
		for i := 0; i < count; i++ {
			if i%2 == 0 {
				right <- true
			} else {
				right <- false
			}

			fmt.Println("goroutine : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		if <-right {
			fmt.Println("main func : ", i)
		}
	}
}
