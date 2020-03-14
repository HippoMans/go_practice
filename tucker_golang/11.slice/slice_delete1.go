package main

import "fmt"

func main() {
	//마지막 열을 제거한다.
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := len(a) - 1; i >= 0; i-- {
		a = a[:i]
		fmt.Println(a)
	}
}
