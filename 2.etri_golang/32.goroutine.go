package main

import "fmt"

func Introduction() {
	ByeBye()
	fmt.Println("Hello, My name is Hippo\n")
}

func ByeBye() {
	fmt.Println("Good, Bye~~~\n")
}

func main() {
	// Introduction() 함수 고루틴 실행
	go Introduction()

	// ByeBye() 함수 고루틴 실행
	go ByeBye()
	//입력을 기다리고 있다. 이 작업은 go Introduction()이 실행된 후 에도 main() 함수가 끝나지 않도록 해준다.
	fmt.Scanln()
}
