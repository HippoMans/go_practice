package main

import "fmt"

var X string = "Hello World"

func main() {
	//main() 함수 밖에서 변수 선언한 것을 main() 함수 안에서 사용 가능
	fmt.Println(X)

	/* f() 함수에서 선언한 Y 변수를 main()함수에서 사용 불가
	fmt.Println(Y)
	*/
}

func f() {
	var Y string = "Hippo"
	fmt.Println(Y)
}
