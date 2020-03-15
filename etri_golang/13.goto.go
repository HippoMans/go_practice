package main

import "fmt"

func main() {
	var a int = 13

	if a == 1 {
		goto Error1
	} else if a == 2 {
		goto Error2
	} else if a == 3 {
		goto Error3
	} else {
		fmt.Println(a)
		fmt.Println("Success")
	}
	//레이블을 사용하기 위해서는 반드시 모든 return을 사용해야 한다.
	//조건문이 끝났는데도 레이블이 하나도 선택이 되지 않았을 경우 return이 없으면 첫번재 레이블이 실행된다. 주의해야한다.
	return

Error1:
	fmt.Println("Error1")
	return

Error2:
	fmt.Println("Error2")
	return

Error3:
	fmt.Println("Error3")
	return
}
