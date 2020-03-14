package main

import "fmt"

/*
slice(동적 배열)
동적배열은 메모리 공간을 유동적으로 사용하기 위한 것이다.
동적배열로 3개 쓰다가 6개 공간이 필요한 경우 새로운 메모리 공간에 6개의 메모리를 할당하고 동적배열로
할당한 3개의 메모리 값을 6개 메모리에 복사한다.
그 후 3개의 메모리 공간을 삭제해 주면 메모리 공간을 늘어나고 동적 배열의 메모리 주소가 다른 공간이 생긴다.
동적 배열은 배열의 주소를 가진 변수와 값을 가진 배열로 구성되어있다.

var a []int
이 예시를 분석하면 int로 구성된 배열의 주소를 담는 변수가 a이고, a가 가리키는 배열은 int변수로 되어있다.
동적배열의 사용함수 중에는 cap()과 len()이 있다.
* cap()함수는 capacity이고 총 배열의 크기
* len()함수는 배열에 들어있는 값의 배열 크기
* cap < len일 경우 새로운 주소의 배열이 생성되고 값이 복사된다. 이때 cap의 길이는 2배로 늘어난다.
*/

func main() {

	//기본 동적 배열 slice
	var a []int
	fmt.Println("기본 동적 배열(값 x)")
	fmt.Printf("동적 배열의 값 : %d\n", a)
	fmt.Printf("동적 배열의 len : %d\n", len(a))
	fmt.Printf("동적 배열의 cap : %d\n", cap(a))

	fmt.Println("\n동적 배열 할당")
	a = make([]int, 0, 8)
	a = append(a, 10)
	fmt.Printf("동적 배열의 값 : %d\n", a)
	fmt.Printf("동적 배열의 len : %d\n", len(a))
	fmt.Printf("동적 배열의 cap : %d\n", cap(a))
}
