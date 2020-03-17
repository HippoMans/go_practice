package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	s := "Hello, Go World"
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(s, n)
			time.Sleep(1)
		}(i)
	}
	fmt.Scanln()
}
