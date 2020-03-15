package main

import "fmt"

func main() {
	// 변수 선언 첫번째 방법
	var X string = "Hello World"
	fmt.Println(X)

	// 변수 선언 두번재 방법
	var Y string
	Y = "Hippo"
	fmt.Println(Y)

	// 변수 선언 및 변수 변경
	var value int = 10
	fmt.Println(value)
	value = 10000
	fmt.Println(value)

	//변수 값 비교
	var first string = "Mans"
	var second string = "grammer"
	fmt.Println(first == second)
	second = "Mans"
	fmt.Println(first == second)
}
