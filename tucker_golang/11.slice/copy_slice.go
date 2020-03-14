package main

import "fmt"

func main() {
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2
	b := make([]int, len(a))

	fmt.Printf("a가 가리키는 메모리 주소 : %p, b가 가리키는 메모리 주소 : %p\n", a, b)
	fmt.Printf("a의 len : %d, a의 cap : %d\n", len(a), cap(a))
	fmt.Printf("b의 len : %d, b의 cap : %d\n", len(b), cap(b))

	//새로운 메모리 공간에 동적 배열 복사
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}

	b = append(b, 100)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("\na가 가리키는 메모리 주소 : %p, b가 가리키는 메모리 주소 : %p", a, b)
	b[0] = 10
	b[1] = 20
	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()
}
