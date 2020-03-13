package main

import "fmt"

func main() {
	//배열을 이용하여 스택을 만든다.
	stack := []int{}

	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	//전체 stack의 크기를 확인
	fmt.Println("스택(stack)")
	fmt.Println("입력값 : ", stack)

	//stack으로 값을 출력
	fmt.Print("출력값 : ")
	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Printf("%d, ", last)
	}

	fmt.Println("\n")
}
