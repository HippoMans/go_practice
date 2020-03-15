package main

import "fmt"

func main() {
	var num1 uint8 = 3
	var num2 uint8 = 2

	fmt.Println("더하기 : ", num1+num2) // 5
	fmt.Println("빼기 : ", num1-num2)  // 1
	fmt.Println("곱하기 : ", num1*num2) // 6
	fmt.Println("나누기 : ", num1/num2) // 1
	fmt.Println("나머지 : ", num1%num2) // 1
}
