package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for aValue := len(a); aValue > 0; aValue-- {
		a = RemoveBack(a)
		fmt.Println(a)
	}
}

func RemoveBack(a []int) []int {
	sliceLen := len(a)
	b := make([]int, sliceLen)
	b = a[:sliceLen-1]
	return b
}
