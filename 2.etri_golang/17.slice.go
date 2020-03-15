package main

import "fmt"

func main() {
	// 1.slice의 append() 함수 사용
	slice1 := []int{1, 2, 3}
	append_slice := append(slice1, 4, 5)
	fmt.Println("slice1 : ", slice1, "\nappend_slice : ", append_slice)

	// 2. slice의 copy() 함수 사용
	// slice1 내용을 copy_slice로 복사
	// copy_slice의 cap이 2로 설정하였기 때문에 slice1의 2가지 요소만 copy_slice에  들어갈 수 있다.
	copy_slice := make([]int, 2)
	copy(copy_slice, slice1)
	fmt.Println("slice1 : ", slice1, "\ncopy_slice : ", copy_slice)

	// 3. slice를 사용하는 여러가지 방법
	a := []int{1, 2, 3, 4, 5}

	fmt.Println("전체 슬라이스 값 출력 : ", a[:])
	fmt.Println("천체 슬라이스 값 출력 : ", a[0:])
	fmt.Println("전체 슬라이스 값 출력 : ", a[:5])
	fmt.Println("전체 슬라이스 값 출력 : ", a[0:len(a)])
	fmt.Println("4번째 요소부터 값 출력(인덱스는 0번부터 시작) : ", a[3:])
	fmt.Println("4번째 요소앞까지 값 출력(인덱스는 0번부터 시작) : ", a[:3])
	fmt.Println("1번재부터 4번째 요소 앞까지 값 출력(인덱스는 0번부터 시작) : ", a[1:3])
}
