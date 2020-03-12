package main

import "fmt"

func main() {
	a := make([]int, 2, 4) // len 은 2이고 cap은 4인 동적 배열 할당
	b := append(a, 3)      // a에 3의 값을 넣은 동적 배열에 주소값을 b에 넣는다.

	// a와 b는 같은 메모리 공간을 가진다.
	fmt.Printf("\na의 메모리 주소 : %p, b의 메모리 주소 : %p\n", &a, &b)
	fmt.Printf("a가 가리키는 동적 할당 메모리 공간 : %p, b가 가리키는 동적 할당 메모리 공간 : %p\n", a, b)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%d , ", a[i])
	}
	fmt.Println()

	for i := 0; i < len(b); i++ {
		fmt.Printf("%d , ", b[i])
	}
	fmt.Println()

	a[0] = 10000
	a[1] = 369369

	fmt.Println("\na와 b가 가리키고 있는 배열의 메모리 공간이 같다. 그래서 a의 배열값을 변경하면 b의 배열값도 변경된다.")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d , ", a[i])
	}
	fmt.Println()

	for i := 0; i < len(b); i++ {
		fmt.Printf("%d , ", b[i])
	}
	fmt.Println()
}
