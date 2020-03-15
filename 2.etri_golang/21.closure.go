package main

import "fmt"

func main() {

	// 첫번째  closure로 main함수에서 활용
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(3, 4))

	// 두번째 closure로 main함수에서 활용
	a := 0
	increment := func() int {
		a++
		return a
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
