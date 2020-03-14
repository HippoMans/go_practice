package main

import "fmt"

// 동적배열은 cap과 len이 있으며, cap보다 len이 커질 때 새로운 메모리 공간에 값이 들어간다.
func main() {
	a := []int{1, 2}
	b := append(a, 3) // a에 3을 추가한 동적 배열의 메모리 값을 b에 할당

	fmt.Printf("a의 메모리 주소 : %p, b의 메모리 주소 : %p\n", &a, &b)
	fmt.Printf("a의 len : %d, a의 cap : %d\n", len(a), cap(a))
	fmt.Printf("b의 len : %d, b의 cap : %d\n", len(b), cap(b))

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d , ", a[i])
	}
	fmt.Println()

	for i := 0; i < len(b); i++ {
		fmt.Printf("%d , ", b[i])
	}
	fmt.Println()

	a[0] = 4
	a[1] = 5

	fmt.Printf("a의 메모리 주소 : %p\n", &a)
	fmt.Printf("a의 len : %d, a의 cap : %d\n", len(a), cap(a))

	b = append(b, 1000)
	fmt.Printf("b의 메모리 주소 : %p\n", &b)
	fmt.Printf("b의 len : %d, b의 cap : %d\n", len(b), cap(b))

	b = append(b, 10000)
	fmt.Printf("b의 메모리 주소 : %p\n", &b)
	fmt.Printf("b의 len : %d, b의 cap : %d\n", len(b), cap(b))

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d , ", a[i])
	}
	fmt.Println()

	for i := 0; i < len(b); i++ {
		fmt.Printf("%d , ", b[i])
	}
	fmt.Println()
}
