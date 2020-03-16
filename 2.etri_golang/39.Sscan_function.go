package main

import "fmt"

func main() {
	var n1 int
	var n2 float32
	var s1 string
	what1 := "1\n1.1\nHello"

	fmt.Println("fmt.Sscan() 함수 사용")
	n, _ := fmt.Sscan(what1, &n1, &n2, &s1)
	fmt.Println("Input Number : ", n)
	fmt.Println(n1, n2, s1)

	fmt.Println("fmt.Sscanln() 함수 사용")
	what2 := "100 200.21 Good"
	n, _ = fmt.Sscanln(what2, &n1, &n2, &s1)
	fmt.Println("Input Number : ", n)
	fmt.Println(n1, n2, s1)

	fmt.Println("fmt.Sscanf() 함수 사용")
	what3 := "88, 3.2E+2, Victory"
	n, _ = fmt.Sscanf(what3, "%d, %f, %s", &n1, &n2, &s1)
	fmt.Println("Input Number : ", n)
	fmt.Println(n1, n2, s1)
}
