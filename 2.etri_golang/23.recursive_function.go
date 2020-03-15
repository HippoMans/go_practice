package main

import "fmt"

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("재귀함수")
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
}
