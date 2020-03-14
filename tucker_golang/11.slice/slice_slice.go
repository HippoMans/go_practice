package main

import "fmt"

/*
동적 배열을 slice해서 동적배열을 만든다.
이때 만든 동적 배열의 시작 주소값은 기초 동적 배열 메모리 안에서 결정된다.
cap = 끝값 - 시작값을 의미한다.
슬라이스 한다고 새로운 배열의 메모리 공간이 생기는 것이 아니다.
*/

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println()
	fmt.Println("a : ", a)
	fmt.Println()
	//시작과 끝을 작성하여 slice
	b := a[5:8]
	fmt.Printf("a가 가리키는 메모리 주소 : %p, b가 가리키는 메모리 주소 : %p\n", a, b)
	fmt.Println("b : ", b)

	//시작만 작성하여 slice
	c := a[1:]
	fmt.Printf("\na가 가리키는 메모리 주소 : %p, c가 가리키는 메모리 주소 : %p\n", a, c)
	fmt.Println("c : ", c)

	//끝을 작성하여 slice
	d := a[:7]
	fmt.Printf("\na가 가리키는 메모리 주소 : %p, d가 가리키는 메모리 주소 : %p\n", a, d)
	fmt.Println("d : ", d)

	fmt.Println("\n만약 d의 값을 변경할 경우 a, b, c, d 모두 값이 변경하는 것을 확인할 수 있다.")
	d[5] = 1000
	d[6] = 9999
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
}
