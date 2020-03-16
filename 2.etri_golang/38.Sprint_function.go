package main

import "fmt"

func main() {
	//fmt.Sprint() 함수 사용
	fmt.Println("fmt.Sprint() 함수 사용")
	var s1 string
	s1 = fmt.Sprint(10, 4.2, "Hello World")
	fmt.Println(s1)

	//fmt.Sprintln() 함수 사용
	fmt.Println("fmt.Sprintln() 함수 사용")
	var s2 string
	s2 = fmt.Sprintln(100, 2.31e+4, "Success")
	fmt.Println(s2)

	//fmt.Sprintf() 함수 사용
	fmt.Println("fmt.Sprintf() 함수 사용")
	var s3 string
	s3 = fmt.Sprintf("%d %f %s\n", 99, 12.32, "Victory")
	fmt.Println(s3)
}
