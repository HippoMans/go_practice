package main

/*
//소수점 사용
var num1 float32 = 0.1
var num2 float32 = .35
var num3 float32 = 132.73287

//지수 표기법 사용
var num4 float32 = 1e7
var num5 float64 = .12345E+2
var num6 float64 = 5.32521e-10
*/

import "fmt"

func main() {
	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}
	fmt.Println(a)
	if a == 9.0 {
		fmt.Println("결과값과 예상값이 같다.")
	} else {
		fmt.Println("결과값과 예상값이 다르다.")
	}
}
