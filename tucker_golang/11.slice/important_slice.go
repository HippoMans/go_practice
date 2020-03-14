package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 10000
	s[1] = 20000
	s[2] = 30000

	fmt.Printf("주소 : %p, 배열의 길이 : %d, 배열의 용량 : %d\n", s, len(s), cap(s))
	fmt.Println(s)
	fmt.Println()

	//용량을 초과하여 더 큰 메모리 크기가 필요한 경우 새로운 메모리를 찾아서 복사
	u := append(s, 40, 50, 60, 70)
	fmt.Printf("주소 : %p, 배열의 길이 : %d, 배열의 용량 : %d\n", u, len(u), cap(u))
	fmt.Println(u)
	fmt.Println()

	//용량보다 작기 때문에 같은 메모리 공간에 값을 대입한다. 주소값이 같다.
	x := append(u, 8)
	fmt.Printf("주소 : %p, 배열의 길이 : %d, 배열의 용량 : %d\n", x, len(x), cap(x))
	fmt.Println(x)
	fmt.Println()

	//z의 메모리 공간과 x의 메모리 공간은 엄연히 다르다.
	z := make([]int, len(x))
	z = append(x, 9000000)
	fmt.Printf("주소 : %p, 배열의 길이 : %d, 배열의 용량 : %d\n", z, len(z), cap(z))
	fmt.Println(z)
}
