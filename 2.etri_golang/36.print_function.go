package main

import "fmt"

func main() {
	// fmt.Println() 함수
	// 정수
	fmt.Println("fmt.Println()함수 사용")
	var num1 int = 5
	// 부동소수점
	var num2 float32 = 5.2
	// 복소수
	var num3 complex64 = 5.5 + 8.1i
	// 문자열
	var s string = "Hello Go World"
	// 불린형
	var b bool = false
	// 슬라이스
	var a []int = []int{1, 2, 3}
	// 맵
	var m map[string]int = map[string]int{"Hippo": 20}
	// 포인터
	var p *int = new(int)
	// 구조체
	type Data struct{ a, b int }
	var data Data = Data{10, 20}
	// 빈 인터페이스(empty interface)
	var i interface{} = 5

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(m)
	fmt.Println(p)
	fmt.Println(data)
	fmt.Println(i)
	fmt.Println(num1, num2, num3, s, b)
	fmt.Println(p, a, m)
	fmt.Println(data, i)

	fmt.Println("\nfmt.Printf()함수 사용")
	//fmt.Printf() 함수
	fmt.Printf("%d\n", num1)
	fmt.Printf("%f\n", num2)
	fmt.Printf("%f\n", num3)
	fmt.Printf("%s\n", s)
	fmt.Printf("%t\n", b) // boolean의 format은 %t이다.
	// 아래에 사용하는 %v format은 모든 값 사용 가능한 기본 지정자이다.
	// slice, map, pointer, struct, interface 등에 사용한다.
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", m)
	fmt.Printf("%p\n", p)
	fmt.Printf("%v\n", data)
	fmt.Printf("%v\n", i)
	fmt.Printf("%d, %f %s\n", num1, num2, s)
}
